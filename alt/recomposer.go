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
		if len(f.Name) == 0 || ([]byte(f.Name)[0]&0x20) != 0 {
			continue
		}
		ft := f.Type
		switch ft.Kind() {
		case reflect.Array, reflect.Slice, reflect.Map, reflect.Ptr:
			ft = ft.Elem()
		}
		if _, has := r.composers[ft.Name()]; has {
			continue
		}
		_, _ = r.registerComposer(ft, nil)
	}
	return c, nil
}

func (r *Recomposer) registerAnyComposer(rt reflect.Type, fun RecomposeAnyFunc) (*composer, error) {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	full := rt.PkgPath() + "/" + rt.Name()
	if rt.Kind() != reflect.Struct {
		return nil, fmt.Errorf("only structs can be recomposed. %s is not a struct type", rt)
	}
	c := r.composers[full]
	if c == nil {
		c = &composer{
			any:   fun,
			short: rt.Name(),
			full:  full,
			rtype: rt,
		}
		c.indexes = indexType(c.rtype)
		r.composers[c.short] = c
		r.composers[c.full] = c
	} else {
		c.any = fun
	}
	return c, nil
}

// Recompose simple data into more complex go types.
func (r *Recomposer) Recompose(v any, tv ...any) (out any, err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = ojg.NewError(rec)
			out = nil
		}
	}()
	out = r.MustRecompose(v, tv...)
	return
}

// MustRecompose simple data into more complex go types.
func (r *Recomposer) MustRecompose(v any, tv ...any) (out any) {
	if 0 < len(tv) {
		if um, ok := tv[0].(json.Unmarshaler); ok {
			if comp := r.composers["json.Unmarshaler"]; comp != nil {
				b, _ := comp.any(v) // Special case. Must return []byte.
				if err := um.UnmarshalJSON(b.([]byte)); err != nil {
					panic(err)
				}
				return um
			}
		}
		out = tv[0]
		rv := reflect.ValueOf(tv[0])
		switch rv.Kind() {
		case reflect.Array, reflect.Slice:
			rv = reflect.New(rv.Type())
			r.recomp(v, rv)
			out = rv.Elem().Interface()
		case reflect.Map:
			r.recomp(v, rv)
		case reflect.Ptr:
			r.recomp(v, rv)
			switch rv.Elem().Kind() {
			case reflect.Slice, reflect.Array, reflect.Map, reflect.Interface:
				out = rv.Elem().Interface()
			}
		default:
			panic(fmt.Errorf("only a slice, map, or pointer is allowed as an optional argument"))
		}
	} else {
		out = r.recompAny(v)
	}
	return
}

func (r *Recomposer) recompAny(v any) any {
	switch tv := v.(type) {
	case nil, bool, int64, float64, string, time.Time:
	case int:
		v = int64(tv)
	case int8:
		v = int64(tv)
	case int16:
		v = int64(tv)
	case int32:
		v = int64(tv)
	case uint:
		v = int64(tv)
	case uint8:
		v = int64(tv)
	case uint16:
		v = int64(tv)
	case uint32:
		v = int64(tv)
	case uint64:
		v = int64(tv)
	case float32:
		// This small rounding makes the conversion from 32 bit to 64 bit
		// display nicer.
		f, i := math.Frexp(float64(tv))
		f = float64(int64(f*fracMax)) / fracMax
		v = math.Ldexp(f, i)
	case []any:
		a := make([]any, len(tv))
		for i, m := range tv {
			a[i] = r.recompAny(m)
		}
		v = a
	case map[string]any:
		if cv := tv[r.CreateKey]; cv != nil {
			tn, _ := cv.(string)
			if c := r.composers[tn]; c != nil {
				if c.fun != nil {
					val, err := c.fun(tv)
					if err != nil {
						panic(err)
					}
					return val
				}
				rv := reflect.New(c.rtype)
				r.recomp(v, rv)
				return rv.Interface()
			}
		}
		o := map[string]any{}
		for k, m := range tv {
			o[k] = r.recompAny(m)
		}
		v = o

	case gen.Bool:
		v = bool(tv)
	case gen.Int:
		v = int64(tv)
	case gen.Float:
		v = float64(tv)
	case gen.String:
		v = string(tv)
	case gen.Time:
		v = time.Time(tv)
	case gen.Big:
		v = string(tv)
	case gen.Array:
		a := make([]any, len(tv))
		for i, m := range tv {
			a[i] = r.recompAny(m)
		}
		v = a
	case gen.Object:
		if cv := tv[r.CreateKey]; cv != nil {
			gn, _ := cv.(gen.String)
			tn := string(gn)
			if c := r.composers[tn]; c != nil {
				simple, _ := tv.Simplify().(map[string]any)
				if c.fun != nil {
					val, err := c.fun(simple)
					if err != nil {
						panic(err)
					}
					return val
				}
				rv := reflect.New(c.rtype)
				r.recomp(simple, rv)
				return rv.Interface()
			}
		}
		o := map[string]any{}
		for k, m := range tv {
			o[k] = r.recompAny(m)
		}
		v = o

	default:
		panic(fmt.Errorf("can not recompose a %T", v))
	}
	return v
}

func (r *Recomposer) recomp(v any, rv reflect.Value) {
	as, _ := rv.Interface().(AttrSetter)
	if rv.Kind() =