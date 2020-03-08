// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/gen"
)

// DefaultRecomposer provides a shared Recomposer. Note that this should not
// be shared across go routines unless all types that will be used are
// registered first. That can be done explicitly or with a warm up run.
var DefaultRecomposer = Recomposer{
	composers: map[string]*composer{},
}

// RecomposeFunc should build an object from data in a map returning the
// recomposed object or an error.
type RecomposeFunc func(map[string]any) (any, error)

// RecomposeAnyFunc should build an object from data in an any
// returning the recomposed object or an error.
type RecomposeAnyFunc func(any) (any, error)

// Recomposer is used to recompose simple data into structs.
type Recomposer struct {

	// CreateKey identifies the creation key in decomposed objects.
	CreateKey string

	composers map[string]*composer
}

var jsonUnmarshalerType reflect.Type

func init() {
	jsonUnmarshalerType = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

}

// RegisterComposer regsiters a composer function for a value type. A nil
// function will still register the default composer which uses reflection.
func (r *Recomposer) RegisterComposer(val any, fun RecomposeFunc) error {
	_, err := r.registerComposer(reflect.TypeOf(val), fun)

	return err
}

// RegisterAnyComposer regsiters a composer function for a value type. A nil
// function will still register the default composer which uses reflection.
func (r *Recomposer) RegisterAnyComposer(val any, fun RecomposeAnyFunc) error {
	_, err := r.registerAnyComposer(reflect.TypeOf(val), fun)

	return err
}

// RegisterUnmarshalerComposer regsiters a composer function for a named
// value. This is only used to register cross package json.Unmarshaler
// composer which returns []byte.
func (r *Recomposer) RegisterUnmarshalerComposer(fun RecomposeAnyFunc) {
	name := "json.Unmarshaler"
	r.composers[name] = &composer{
		any:   fun,
		short: name,
		full:  name,
	}
}

func (r *Recomposer) registerComposer(rt reflect.Type, fun RecomposeFunc) (*composer, error) {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	full := rt.PkgPath() + "/" + rt.Name()
	// TBD could loosen this up and allow any type as long as a function is provided.
	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only structs can be recomposed. %s is not a struct type", rt)
	}
	c := r.composers[full]
	if c == nil {
		c = &composer{
			fun:   fun,
			short: rt.Name(),
			full:  full,
			rtype: rt,
		}
		c.indexes = indexType(c.rtype)
		r.composers[c.short] = c
		r.composers[c.full] = c
	} else {
		if fun != nil {
			c.fun = fun
		}
		// If already registered then there is no reason to walk the fields again.
		return c, nil
	}
	for i := rt.NumField() - 1; 0 <= i; i-- {
		f := rt.Field(i)
		// Private fields should be skipped.
		if len(f.Name) == 0 || ([]byte(f.Name)[0]&0x20) 