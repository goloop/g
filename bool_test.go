package g

import (
	"testing"
)

func TestStringToBool(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
		def      bool
		err      bool
	}{
		{"true", true, false, false},
		{"false", false, false, false},
		{"yes", true, false, false},
		{"no", false, false, false},
		{"on", true, false, false},
		{"off", false, false, false},
		{"", true, true, true},
		{"", false, false, true},
		{"abc", false, false, true},
	}

	for _, test := range tests {
		result, err := StringToBool(test.input, test.def)
		if (err != nil) != test.err {
			t.Errorf("StringToBool(%q) error = %v, wantErr %v",
				test.input, err, test.err)
			continue
		}

		if result != test.expected {
			t.Errorf("StringToBool(%q) = %v, want %v",
				test.input, result, test.expected)
		}
	}
}

func TestBoolToString(t *testing.T) {
	tests := []struct {
		input    bool
		expected string
	}{
		{true, "true"},
		{false, "false"},
	}

	for _, test := range tests {
		result := BoolToString(test.input)
		if result != test.expected {
			t.Errorf("BoolToString(%v) = %q, want %q",
				test.input, result, test.expected)
		}
	}
}
