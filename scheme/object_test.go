// gopische/scheme/object_test.go

package scheme

import (
	"testing"
)

func TestNewSchemeObjectNil(t *testing.T) {
	id := 100
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
	if sobj.Tag() != expected {
		t.Fatalf("tests[%d] - wrong scheme object tag, expected=%v, got=%v",
			id, expected, sobj.Tag())
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
		{201, obj{BOOLEAN, false}, Tag(BOOLEAN)},
		{202, obj{STRING, "x"}, Tag(STRING)},
		{203, obj{SYMBOL, "car"}, Tag(SYMBOL)},
		{210, obj{NUMBER, 0}, Tag(NUMBER)},
		{211, obj{NUMBER, int8(1)}, Tag(NUMBER)},
		{212, obj{NUMBER, int16(2)}, Tag(NUMBER)},
		{213, obj{NUMBER, int32(3)}, Tag(NUMBER)},
		{214, obj{NUMBER, int64(4)}, Tag(NUMBER)},
		{215, obj{NUMBER, uint(5)}, Tag(NUMBER)},
		{216, obj{NUMBER, uint8(6)}, Tag(NUMBER)},
		{217, obj{NUMBER, uint16(7)}, Tag(NUMBER)},
		{218, obj{NUMBER, uint32(8)}, Tag(NUMBER)},
		{219, obj{NUMBER, uint64(9)}, Tag(NUMBER)},
		{221, obj{NUMBER, float32(0.01)}, Tag(NUMBER)},
		{222, obj{NUMBER, float64(0.02)}, Tag(NUMBER)},
		{231, obj{NUMBER, complex64(0 + 1i)}, Tag(NUMBER)},
		{232, obj{NUMBER, complex128(0 + 2i)}, Tag(NUMBER)},
	}

	for _, tc := range tests {
		sobj, err := NewSchemeObject(tc.testcase.tag, tc.testcase.value)
		if err != nil {
			t.Fatalf("tests[%d] - fail to create a scheme object, expected=%v",
				tc.id, tc.expected)
		}
		if sobj.Tag() != tc.expected {
			t.Fatalf("tests[%d] - wrong scheme object tag, expected=%v, got=%v",
				tc.id, tc.expected, sobj.Tag())
		}
	}
}

func TestStringer(t *testing.T) {
	type obj struct {
		tag   Tag
		value any
	}

	tests := []struct {
		id       int
		testcase obj
		expected string
	}{
		{301, obj{NIL, 0}, "()"},
		{302, obj{BOOLEAN, false}, "#f"},
		{303, obj{BOOLEAN, true}, "#t"},
		{304, obj{STRING, "hoge"}, "\"hoge\""},
		{305, obj{SYMBOL, "foo"}, "foo"},
		{310, obj{NUMBER, 0}, "0"},
		{311, obj{NUMBER, -1}, "-1"},
		{312, obj{NUMBER, 3.14}, "3.14"},
		{313, obj{NUMBER, -1.41}, "-1.41"},
		{314, obj{NUMBER, +1.41}, "1.41"},
		{320, obj{NUMBER, 0 + 0i}, "(0+0i)"},
		{321, obj{NUMBER, 1 + 0i}, "(1+0i)"},
		{322, obj{NUMBER, 0.0 + 1i}, "(0+1i)"},
	}

	for _, tc := range tests {
		sobj, err := NewSchemeObject(tc.testcase.tag, tc.testcase.value)
		if err != nil {
			t.Fatalf("tests[%d] - fail to create a scheme object, expected=%v",
				tc.id, tc.expected)
		}
		str := sobj.String()
		if tc.expected != str {
			t.Fatalf("tests[%d] - wrong stringer for a scheme object, expected=%v, got=%v", tc.id, tc.expected, str)
		}
	}
}
