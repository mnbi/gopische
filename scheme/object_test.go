// gopische/scheme/object_test.go

package scheme

import (
	"testing"
)

func TestTagString(t *testing.T) {
	tests := []struct {
		id       int
		testcase Tag
		expected string
	}{
		{0x00, NIL, "nil"},
		{0x01, BOOLEAN, "boolean"},
		{0x02, NUMBER, "number"},
		{0x03, STRING, "string"},
		{0x04, SYMBOL, "symbol"},
		{0x81, LIST, "list"},
		{0xff, 0xff, "illegal"}, // id = 255
	}

	for _, tc := range tests {
		str := tc.testcase.String()
		if str != tc.expected {
			t.Fatalf("tests[%d] - wrong tag name, expected=%q, got=%q",
				tc.id, tc.expected, str)
		}
	}
}

func TestNewSchemeObjectNil(t *testing.T) {
	id := 300
	testcase := struct {
		tag   Tag
		value any
	}{NIL, nil}
	expected := Tag(NIL)

	sobj, err := NewSchemeObject(testcase.tag, testcase.value)
	if err != nil {
		t.Fatalf("tests[%d] - fail to create a scheme object, expected=%v",
			id, expected)
	}
	if sobj.Tag != expected {
		t.Fatalf("tests[%d] - wrong scheme object tag, expected=%v, got=%v",
			id, expected, sobj.Tag)
	}
}

func TestNewSchemeObject(t *testing.T) {
	type obj struct {
		tag   Tag
		value any
	}

	tests := []struct {
		id       int
		testcase obj
		expected Tag
	}{
		{301, obj{BOOLEAN, false}, Tag(BOOLEAN)},
		{302, obj{STRING, "x"}, Tag(STRING)},
		{303, obj{SYMBOL, "car"}, Tag(SYMBOL)},
		{310, obj{NUMBER, 0}, Tag(NUMBER)},
		{311, obj{NUMBER, int8(1)}, Tag(NUMBER)},
		{312, obj{NUMBER, int16(2)}, Tag(NUMBER)},
		{313, obj{NUMBER, int32(3)}, Tag(NUMBER)},
		{314, obj{NUMBER, int64(4)}, Tag(NUMBER)},
		{315, obj{NUMBER, uint(5)}, Tag(NUMBER)},
		{316, obj{NUMBER, uint8(6)}, Tag(NUMBER)},
		{317, obj{NUMBER, uint16(7)}, Tag(NUMBER)},
		{318, obj{NUMBER, uint32(8)}, Tag(NUMBER)},
		{319, obj{NUMBER, uint64(9)}, Tag(NUMBER)},
		{321, obj{NUMBER, float32(0.01)}, Tag(NUMBER)},
		{322, obj{NUMBER, float64(0.02)}, Tag(NUMBER)},
		{331, obj{NUMBER, complex64(0 + 1i)}, Tag(NUMBER)},
		{332, obj{NUMBER, complex128(0 + 2i)}, Tag(NUMBER)},
	}

	for _, tc := range tests {
		sobj, err := NewSchemeObject(tc.testcase.tag, tc.testcase.value)
		if err != nil {
			t.Fatalf("tests[%d] - fail to create a scheme object, expected=%v",
				tc.id, tc.expected)
		}
		if sobj.Tag != tc.expected {
			t.Fatalf("tests[%d] - wrong scheme object tag, expected=%v, got=%v",
				tc.id, tc.expected, sobj.Tag)
		}
	}
}
