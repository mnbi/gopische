// gopische/scheme/tag_test.go

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
		{0x02, STRING, "string"},
		{0x03, SYMBOL, "symbol"},
		{0x04, CHARACTER, "character"},
		{0x70, NUMBER, "number"},
		{0x71, INT, "number(int)"},
		{0x72, FLOAT, "number(float)"},
		{0x73, COMPLEX, "number(complex)"},
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
