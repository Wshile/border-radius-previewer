// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen

import (
	"sort"
	"strconv"
	"time"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
)

func (wr *Writer) colorSEN(data any, depth int) {
	switch td := data.(type) {
	case nil:
		wr.buf = append(wr.buf, wr.NullColor...)
		wr.buf = append(wr.buf, []byte("null")...)

	case bool:
		wr.buf = append(wr.buf, wr.BoolColor...)
		if td {
			wr.buf = append(wr.buf, []byte("true")...)
		} else {
			wr.buf = append(wr.buf, []byte("false")...)
		}

	case int:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int8:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int16:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int32:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendInt(wr.buf, int64(td), 10)
	case int64:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendInt(wr.buf, td, 10)
	case uint:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint8:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = strconv.AppendUint(wr.buf, uint64(td), 10)
	case uint16:
		wr.buf = append(wr.buf, wr.NumberColor...)
		wr.buf = st