// gopische/scheme/object.go

package scheme

import (
	"errors"
	"fmt"
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

type Object interface {
	// Returns a tag of a Scheme data object.
	Tag() Tag

	// Returns a subclass of a Scheme data object, if it has a subclass.
	// When no subclass, returns a zero value.
	SubClass() SubClass

	// Returns a value of a Scheme data object.
	Value() any
}

// Nil object
type Nil struct{}

func (sobj Nil) Tag() Tag {
	return Tag(NIL)
}

func (sobj Nil) SubClass() SubClass {
	return 0
}

func (sobj Nil) Value() any {
	return nil
}

// Boolean object
type Boolean struct {
	value bool
}

func (sobj Boolean) Tag() Tag {
	return Tag(BOOLEAN)
}

func (sobj Boolean) SubClass() SubClass {
	return 0
}

func (sobj Boolean) Value() any {
	return sobj.value
}

// String object
type String struct {
	value string
}

func (sobj String) Tag() Tag {
	return Tag(STRING)
}

func (sobj String) SubClass() SubClass {
	return 0
}

func (sobj String) Value() any {
	return sobj.value
}

// Symbol object
type Symbol struct {
	value string
}

func (sobj Symbol) Tag() Tag {
	return Tag(SYMBOL)
}

func (sobj Symbol) SubClass() SubClass {
	return 0
}

func (sobj Symbol) Value() any {
	return sobj.value
}

// Number object
type Number struct {
	tag   Tag
	value any
}

func (sobj Number) Tag() Tag {
	return Tag(NUMBER)
}

func (sobj Number) SubClass() SubClass {
	return sobj.tag.subClass()
}

func (sobj Number) Value() any {
	return sobj.value
}

// factory
func CreateSchemeObject(tag Tag, value any) (Object, error) {
	var sobj Object

	var ok bool
	var emsg string

	switch tag {
	case NIL:
		ok = true
		sobj = Nil{}
	case BOOLEAN:
		var bv bool
		if bv, ok = value.(bool); ok {
			sobj = Boolean{value: bv}
		} else {
			emsg = fmt.Sprintf("illegal boolean value, %v", value)
		}
	case STRING:
		var str string
		if str, ok = value.(string); ok {
			sobj = String{value: str}
		} else {
			emsg = fmt.Sprintf("illegal string value, %v", value)
		}
	case SYMBOL:
		var sym string
		if sym, ok = value.(string); ok {
			sobj = Symbol{value: sym}
		} else {
			emsg = fmt.Sprintf("illegal symbol value, %v", value)
		}
	case NUMBER:
		sobj, ok = newNumber(value)
		if !ok {
			emsg = fmt.Sprintf("illegal number value, %v", value)
		}
	default:
		ok = false
		emsg = fmt.Sprintf("illegal tag as Sobj, %s", tag)
	}

	if !ok {
		return nil, errors.New(emsg)
	}

	return sobj, nil
}

func newNumber(v any) (sobj Object, ok bool) {
	switch v.(type) {
	case int:
		iv := v.(int)
		sobj, ok = Number{tag: INT, value: int64(iv)}, true
	case int8:
		i8v := v.(int8)
		sobj, ok = Number{tag: INT, value: int64(i8v)}, true
	case int16:
		i16v := v.(int16)
		sobj, ok = Number{tag: INT, value: int64(i16v)}, true
	case int32:
		i32v := v.(int32)
		sobj, ok = Number{tag: INT, value: int64(i32v)}, true
	case int64:
		sobj, ok = Number{tag: INT, value: v.(int64)}, true
	case uint:
		uv := v.(uint)
		sobj, ok = Number{tag: INT, value: int64(uv)}, true
	case uint8:
		u8v := v.(uint8)
		sobj, ok = Number{tag: INT, value: int64(u8v)}, true
	case uint16:
		u16v := v.(uint16)
		sobj, ok = Number{tag: INT, value: int64(u16v)}, true
	case uint32:
		u32v := v.(uint32)
		sobj, ok = Number{tag: INT, value: int64(u32v)}, true
	case uint64:
		u64v := v.(uint64)
		sobj, ok = Number{tag: INT, value: int64(u64v)}, true
	case float32:
		f32v := v.(float32)
		sobj, ok = Number{tag: FLOAT, value: float64(f32v)}, true
	case float64:
		sobj, ok = Number{tag: FLOAT, value: v.(float64)}, true
	case complex64:
		c64v := v.(complex64)
		sobj, ok = Number{tag: COMPLEX, value: complex128(c64v)}, true
	case complex128:
		sobj, ok = Number{tag: COMPLEX, value: v.(complex128)}, true
	default:
		sobj, ok = Nil{}, false
	}
	return
}
