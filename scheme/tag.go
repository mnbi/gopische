// gopische/scheme/tag.go

package scheme

// Tag holds a class of a Scheme object.
//
//	0b xxxx xxxx xxxx xxxx
//	   <-------> <--> <-->
//	    |         |    |
//	    |         |    +-- sub class (e.g. integer, float, ... in number class)
//	    |         +-- data class (boolean, number, ...)
//	    +-- reserved for the future use
type Tag uint16

// Class and SubClass are used to store bit pattern which are
// represent a class of data, and subclass of a class.
type Class uint16
type SubClass uint16

const (
	// nil
	NIL = 0x0000
	// simple data object
	BOOLEAN   = 0x0010 // 0b 0000 0000 0001 0000
	STRING    = 0x0020 // 0b 0000 0000 0010 0000
	SYMBOL    = 0x0030 // 0b 0000 0000 0011 0000
	CHARACTER = 0x0040 // 0b 0000 0000 0100 0000
	// gap(0x50 - 0x6f) - reserved for the future
	NUMBER = 0x0070 // 0b 0000 0000 0111 0000
	// 0b 1000 0000 - not used
	LIST = 0x0090 // 0b 0000 0000 1001 0000
	// number class (NumClass)
	// - 0b 0000 0000 0111 0000 - (not used)
	// - 0b 0000 0000 0111 0xxx - represents with go primitive types
	// - 0b 0000 0000 0111 1000 - (not used)
	// - 0b 0000 0000 0111 1xxx - arbitrary-precision integer
	INT     = 0x0071
	FLOAT   = 0x0072
	COMPLEX = 0x0073
	BIGINT  = 0x0079 // resereved for future
)

func bitsNil() Class {
	return Class(NIL >> 4)
}

func bitsBoolean() Class {
	return Class(BOOLEAN >> 4)
}

func bitsString() Class {
	return Class(STRING >> 4)
}

func bitsSymbol() Class {
	return Class(SYMBOL >> 4)
}

func bitsNumber() Class {
	return Class(NUMBER >> 4)
}

func bitsList() Class {
	return Class(LIST >> 4)
}

func bitsInt() SubClass {
	return SubClass(INT & subClassMask)
}

func bitsFloat() SubClass {
	return SubClass(FLOAT & subClassMask)
}

func bitsComplex() SubClass {
	return SubClass(COMPLEX & subClassMask)
}

func bitsBigint() SubClass {
	return SubClass(BIGINT & subClassMask)
}

const (
	classMask    = 0x00f0
	subClassMask = 0x000f
	compoundMask = 0x0080
)

func (tag Tag) Class() Class {
	var bits Class = Class((classMask & tag) >> 4)
	return bits
}

func (tag Tag) subClass() SubClass {
	var bits SubClass = SubClass(subClassMask & tag)
	return bits
}

func (tag Tag) String() string {
	var name string

	switch tag {
	case NIL:
		name = "nil"
	case BOOLEAN:
		name = "boolean"
	case STRING:
		name = "string"
	case SYMBOL:
		name = "symbol"
	case CHARACTER:
		name = "character"
	case LIST:
		name = "list"
	case NUMBER:
		name = "number"
	case INT:
		name = "number(int)"
	case FLOAT:
		name = "number(float)"
	case COMPLEX:
		name = "number(complex)"
	case BIGINT:
		name = "number(bigint)"
	default:
		name = "illegal"
	}

	return name
}
