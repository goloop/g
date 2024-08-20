package g

import (
	"errors"
	"strings"
)

// StringToBool converts a string to a boolean.
// It handles various string representations of boolean values such as
// "true", "false", "yes", "no", "on", "off". If the conversion fails
// and a default value is provided, it returns the default value. Otherwise,
// it returns an error.
//
// Example Usage:
//
//	b, err := StringToBool("true") // true, nil
//	b, err := StringToBool("yes")  // true, nil
//	b, err := StringToBool("abc", false) // false, error
func StringToBool(v string, def ...bool) (bool, error) {
	var d bool = false

	if len(def) > 0 {
		d = def[0]
	}

	v = strings.ToLower(v)

	switch v {
	case "true", "yes", "on":
		return true, nil
	case "false", "no", "off":
		return false, nil
	case "":
		return d, errors.New("empty string and no default value")
	default:
		return d, errors.New("invalid boolean string")
	}
}

// BoolToString converts a boolean to a string.
// It returns "true" for true and "false" for false.
//
// Example Usage:
//
//	s := BoolToString(true) // "true"
//	s := BoolToString(false) // "false"
func BoolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
