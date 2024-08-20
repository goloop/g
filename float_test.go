package g

import (
	"testing"
)

func TestStringToFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		def      float64
		err      bool
	}{
		{"3.14", 3.14, 0, false},
		{"abc", 1.23, 1.23, true},
		{"", 1.23, 1.23, true},
		{"", 0, 0, true},
	}

	for _, test := range tests {
		result, err := StringToFloat(test.input, test.def)
		if (err != nil) != test.err {
			t.Errorf("StringToFloat(%q) error = %v, wantErr %v", test.input, err, test.err)
			continue
		}

		if result != test.expected {
			t.Errorf("StringToFloat(%q) = %f, want %f", test.input, result, test.expected)
		}
	}
}

func TestFloatToString(t *testing.T) {
	tests := []struct {
		input    any // Using 'any' to handle multiple types
		expected string
	}{
		{3.14, "3.14"},
		{float32(3.14), "3.14"},
		{float64(3.14), "3.14"},
	}

	for _, test := range tests {
		var result string
		switch v := test.input.(type) {
		case float32:
			result = FloatToString(v)
		case float64:
			result = FloatToString(v)
		default:
			t.Errorf("Unsupported type %T", v)
			continue
		}

		if result != test.expected {
			t.Errorf("FloatToString(%v) = %q, want %q", test.input, result, test.expected)
		}
	}
}
