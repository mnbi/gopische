// gopische/scheme/object.go

// Package scheme implements representations of Scheme data objects.
// It also provides an generalized interface to handle a scheme data
// object.
package scheme

import (
	"errors"
	"fmt"
)

// Object provides a generalized interface to handle a Scheme data
// object.
//
// Simple data objects:
// - boolean
// - string
// - symbol
// - character (not implemented yet in this version)
// - number
//
// Compound data objects:
// - list
// - procedure
// - vector (not implemented yet in this version)
// - port (not implemented yet in this version)
type Object interface {
	// Returns a tag of a Scheme data object.
	Tag() Tag

	// Returns a subclass of a Scheme data object, if it has a subclass.
	// When no subclass, returns a zero value.
	SubClass() SubClass

	// Returns a value of a Scheme data object.
	Value() any

	// Following method is used to implement Scheme predicates for
	// adata object.
	IsClass(bits Class) bool

	String() string
}

// Nil object
type Nil struct{}

var nilObject = &Nil{}

// Empty list (aka nil) should be a singleton.
var EmptyList = nilObject

func (sobj *Nil) Tag() Tag {
	return Tag(NIL)
}

func (sobj *Nil) SubClass() SubClass {
	return 0
}

func (sobj *Nil) Value() any {
	return nil
}

func (sobj *Nil) IsClass(bits Class) bool {
	return bits == bitsNil()
}

func (sobj *Nil) String() string {
	return fmt.Sprint("()")
}

// Boolean object
type Boolean struct {
	value bool
}

func (sobj *Boolean) Tag() Tag {
	return Tag(BOOLEAN)
}

func (sobj *Boolean) SubClass() SubClass {
	return 0
}

func (sobj *Boolean) Value() any {
	return sobj.value
}

func (sobj *Boolean) IsClass(bits Class) bool {
	return bits == bitsBoolean()
}

func (sobj *Boolean) String() string {
	if sobj.value {
		return fmt.Sprint("#t")
	} else {
		return fmt.Sprint("#f")
	}
}

// String object
type String struct {
	value string
}

func (sobj *String) Tag() Tag {
	return Tag(STRING)
}

func (sobj *String) SubClass() SubClass {
	return 0
}

func (sobj *String) Value() any {
	return sobj.value
}

func (sobj *String) IsClass(bits Class) bool {
	return bits == bitsString()
}

func (sobj *String) String() string {
	return fmt.Sprintf("\"%s\"", sobj.value)
}

// Symbol object
type Symbol struct {
	value string
}

func (sobj *Symbol) Tag() Tag {
	return Tag(SYMBOL)
}

func (sobj *Symbol) SubClass() SubClass {
	return 0
}

func (sobj *Symbol) Value() any {
	return sobj.value
}

func (sobj *Symbol) IsClass(bits Class) bool {
	return bits == bitsSymbol()
}

func (sobj *Symbol) String() string {
	return sobj.value
}

// Number object
type Number struct {
	tag   Tag
	value any
}

func (sobj *Number) Tag() Tag {
	return Tag(NUMBER)
}

func (sobj *Number) SubClass() SubClass {
	return sobj.tag.subClass()
}

func (sobj *Number) Value() any {
	return sobj.value
}

func (sobj *Number) IsClass(bits Class) bool {
	return bits == bitsNumber()
}

func (sobj *Number) String() (str string) {
	switch sobj.value.(type) {
	case int64:
		iv := sobj.value.(int64)
		str = fmt.Sprintf("%d", iv)
	case float64:
		fv := sobj.value.(float64)
		str = fmt.Sprintf("%g", fv)
	case complex128:
		cv := sobj.value.(complex128)
		str = fmt.Sprintf("%g", cv)
	}
	return
}

// factory
func NewSchemeObject(tag Tag, value any) (Object, error) {
	var sobj Object

	var ok bool
	var emsg string

	switch tag {
	case NIL:
		sobj, ok = newNil(value) // newNil will never fail.
	case BOOLEAN:
		if sobj, ok = newBoolean(value); !ok {
			emsg = fmt.Sprintf("illegal boolean value, %v", value)
		}
	case STRING:
		if sobj, ok = newString(value); !ok {
			emsg = fmt.Sprintf("illegal string value, %v", value)
		}
	case SYMBOL:
		if sobj, ok = newSymbol(value); !ok {
			emsg = fmt.Sprintf("illegal symbol value, %v", value)
		}
	case NUMBER:
		if sobj, ok = newNumber(value); !ok {
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

func newNil(v any) (Object, bool) {
	return EmptyList, true
}

func newBoolean(v any) (sobj Object, ok bool) {
	var bv bool
	if bv, ok = v.(bool); ok {
		sobj = &Boolean{value: bv}
	}
	return
}

func newString(v any) (sobj Object, ok bool) {
	var raw, cooked string
	if raw, ok = v.(string); ok {
		if cooked, ok = unescapeGoStr(raw); ok {
			sobj = &String{value: cooked}
		}
	}
	return
}

func newSymbol(v any) (sobj Object, ok bool) {
	var sym string
	if sym, ok = v.(string); ok {
		sobj = &Symbol{value: sym}
	}
	return
}

func newNumber(v any) (sobj Object, ok bool) {
	switch v.(type) {
	case int:
		iv := v.(int)
		sobj, ok = &Number{tag: INT, value: int64(iv)}, true
	case int8:
		i8v := v.(int8)
		sobj, ok = &Number{tag: INT, value: int64(i8v)}, true
	case int16:
		i16v := v.(int16)
		sobj, ok = &Number{tag: INT, value: int64(i16v)}, true
	case int32:
		i32v := v.(int32)
		sobj, ok = &Number{tag: INT, value: int64(i32v)}, true
	case int64:
		sobj, ok = &Number{tag: INT, value: v.(int64)}, true
	case uint:
		uv := v.(uint)
		sobj, ok = &Number{tag: INT, value: int64(uv)}, true
	case uint8:
		u8v := v.(uint8)
		sobj, ok = &Number{tag: INT, value: int64(u8v)}, true
	case uint16:
		u16v := v.(uint16)
		sobj, ok = &Number{tag: INT, value: int64(u16v)}, true
	case uint32:
		u32v := v.(uint32)
		sobj, ok = &Number{tag: INT, value: int64(u32v)}, true
	case uint64:
		u64v := v.(uint64)
		sobj, ok = &Number{tag: INT, value: int64(u64v)}, true
	case float32:
		f32v := v.(float32)
		sobj, ok = &Number{tag: FLOAT, value: float64(f32v)}, true
	case float64:
		sobj, ok = &Number{tag: FLOAT, value: v.(float64)}, true
	case complex64:
		c64v := v.(complex64)
		sobj, ok = &Number{tag: COMPLEX, value: complex128(c64v)}, true
	case complex128:
		sobj, ok = &Number{tag: COMPLEX, value: v.(complex128)}, true
	default:
		sobj, ok = &Nil{}, false
	}
	return
}

func unescapeGoStr(raw string) (string, bool) {
	length := len(raw)

	if length < 1 {
		return "", false
	}

	result := make([]byte, 0, length)

	var ch byte
	for i := 0; i < length; i++ {
		ch = raw[i]
		if ch == 0x5c { // '\'
			if i < length-1 {
				switch raw[i+1] {
				case 0x22: // '"'
					i++
				default:
					// nothing to do
				}
				ch = raw[i]
			} else {
				return "", false
			}
		}
		result = append(result, ch)
	}
	return string(result), true
}
