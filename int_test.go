package g

import (
	"testing"
)

func TestStringToInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		def      int
		err      bool
	}{
		{"123", 123, 0, false},
		{"abc", 10, 10, true},
		{"", 10, 10, true},
		{"", 0, 0, true},
	}

	for _, test := range tests {
		result, err := StringToInt(test.input, test.def)
		if (err != nil) != test.err {
			t.Errorf("StringToInt(%q) error = %v, wantErr %v", test.input, err, test.err)
			continue
		}

		if result != test.expected {
			t.Errorf("StringToInt(%q) = %d, want %d", test.input, result, test.expected)
		}
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		input    any // Using 'any' to handle multiple types
		expected string
	}{
		{123, "123"},
		{int64(123), "123"},
		{int32(123), "123"},
		{uint(123), "123"},
		{uint64(123), "123"},
		{uint32(123), "123"},
	}

	for _, test := range tests {
		var result string
		switch v := test.input.(type) {
		case int:
			result = IntToString(v)
		case int64:
			result = IntToString(v)
		case int32:
			result = IntToString(v)
		case uint:
			result = IntToString(v)
		case uint64:
			result = IntToString(v)
		case uint32:
			result = IntToString(v)
		default:
			t.Errorf("Unsupported type %T", v)
			continue
		}

		if result != test.expected {
			t.Errorf("IntToString(%v) = %q, want %q", test.input, result, test.expected)
		}
	}
}
