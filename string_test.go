package g

import (
	"testing"
)

// TestWeed tests the Weed function.
func TestWeed(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		patterns []string
		want     string
	}{
		{
			name:     "Default Weed",
			s:        "Hello\t World",
			patterns: []string{},
			want:     "Hello World",
		},
		{
			name:     "NonAlphabetics Weed",
			s:        "+380 (96) 123 4567",
			patterns: []string{NonAlphabetics},
			want:     "380961234567",
		},
		{
			name:     "Whitespaces Weed",
			s:        " i @goloop.one\n",
			patterns: []string{Whitespaces},
			want:     "i@goloop.one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Weed(tt.s, tt.patterns...); got != tt.want {
				t.Errorf("Weed() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestTrim tests the Trim function.
func TestTrim(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		patterns []string
		want     string
	}{
		{
			name:     "Default Trim",
			s:        " Hello\t World\r\n",
			patterns: []string{},
			want:     "Hello\t World",
		},
		{
			name:     "Simple Trim",
			s:        "    Go Loop\n   ",
			patterns: []string{},
			want:     "Go Loop",
		},
		{
			name:     "Whitespace Trim",
			s:        " \ti@goloop.one\n ",
			patterns: []string{Whitespaces},
			want:     "i@goloop.one",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.s, tt.patterns...); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
