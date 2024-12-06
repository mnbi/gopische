// gopische/scheme/object.go

package scheme

import (
	"errors"
	"fmt"
	"log"
)

// scheme.Object represents of Scheme data objects.
//
// Simple data objects:
// - boolean
// - number
// - character (not implemented yet in this version)
// - string
// - symbol
//
// Compound data objects:
// - list
// - vector (not implemented yet in this version)

type Tag uint8

const (
	NIL = 0x00
	// simple data object
	BOOLEAN = 0x01
	NUMBER  = 0x02
	STRING  = 0x03
	SYMBOL  = 0x04
	// compound data object
	LIST = 0x81
)

type Object struct {
	Tag   Tag
	value any
}

func (st Tag) String() string {
	var str string

	switch st {
	case NIL:
		str = "nil"
	case BOOLEAN:
		str = "boolean"
	case NUMBER:
		str = "number"
	case STRING:
		str = "string"
	case SYMBOL:
		str = "symbol"
	case LIST:
		str = "list"
	default:
		str = "illegal"
	}

	return str
}

func NewSchemeObject(tag Tag, value any) (*Object, error) {
	var ok bool
	var gobj any
	var emsg string

	switch tag {
	case NIL:
		ok = true
		gobj = nil
	case BOOLEAN:
		ok = true
		gobj = &value
	case NUMBER:
		gobj, ok = newNumber(value)
		log.Printf("NUMBER: %v, %t\n", gobj, ok)
		if !ok {
			emsg = fmt.Sprintf("illegal number value, %v", value)
		}
	case STRING:
		ok = true
		gobj = &value
	case SYMBOL:
		ok = true
		gobj = &value
	default:
		ok = false
		emsg = fmt.Sprintf("illegal tag as Sobj, %s", tag)
	}

	if !ok {
		return nil, errors.New(emsg)
	}

	return &Object{Tag: tag, value: gobj}, nil
}

func newNumber(v any) (value any, ok bool) {
	switch v.(type) {
	case int:
		iv := v.(int)
		value, ok = int64(iv), true
	case int8:
		i8v := v.(int8)
		value, ok = int64(i8v), true
	case int16:
		i16v := v.(int16)
		value, ok = int64(i16v), true
	case int32:
		i32v := v.(int32)
		value, ok = int64(i32v), true
	case int64:
		value, ok = v.(int64), true
	case uint:
		uv := v.(uint)
		value, ok = int64(uv), true
	case uint8:
		u8v := v.(uint8)
		value, ok = int64(u8v), true
	case uint16:
		u16v := v.(uint16)
		value, ok = int64(u16v), true
	case uint32:
		u32v := v.(uint32)
		value, ok = int64(u32v), true
	case uint64:
		u64v := v.(uint64)
		value, ok = int64(u64v), true
	case float32:
		f32v := v.(float32)
		value, ok = float64(f32v), true
	case float64:
		value, ok = v.(float64), true
	case complex64:
		c64v := v.(complex64)
		value, ok = complex128(c64v), true
	case complex128:
		value, ok = v.(complex128), true
	default:
		value, ok = -1, false
	}
	return
}
