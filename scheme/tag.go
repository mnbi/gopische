// gopische/scheme/tag.go

package scheme

// Tag holds a type of a Scheme object. It uses upper 4 bits in a
// unit8 value.
type Tag uint8

// SubClass represents a class of a Scheme object, such as integer,
// floating-point, or complex in a number.  It uses lower 4 bits in a
// uint8 value of a Tag.
type SubClass uint8

const (
	// nil (Tag)
	NIL = 0x00
	// simple data object
	BOOLEAN   = 0x10 // 0b 0001 0000
	STRING    = 0x20 // 0b 0010 0000
	SYMBOL    = 0x30 // 0b 0011 0000
	CHARACTER = 0x40 // 0b 0100 0000
	// gap(0x50 - 0x6f) - reserved for the future
	NUMBER = 0x70 // 0b 0111 0000
	// 0b 1000 0000 - not used
	LIST = 0x90 // 0b 1001 0000
	// number class (NumClass)
	// - 0b 0111 0000 - (not used)
	// - 0b 0111 0xxx - represents with go primitive types
	// - 0b 0111 1000 - (not used)
	// - 0b 0111 1xxx - arbitrary-precision integer, and so on
	INT     = 0x71
	FLOAT   = 0x72
	COMPLEX = 0x73
	BIGNUM  = 0x79 // resereved for future
)

const (
	compoundMask = 0x80
	subClassMask = 0x0f
)

func (tag Tag) isCompound() bool {
	return compoundMask&tag != 0
}

func (tag Tag) isNumber() bool {
	return tag>>4 == LIST>>4
}

func (tag Tag) subClass() SubClass {
	return SubClass(subClassMask & tag)
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
	case BIGNUM:
		name = "number(bignum)"
	default:
		name = "illegal"
	}

	return name
}
