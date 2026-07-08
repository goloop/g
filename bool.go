package g

import (
	"errors"
	"strings"
)

// ErrEmptyString is returned by StringToBool, StringToInt and StringToFloat
// when the input string is empty. Any default value is still returned
// alongside this error, so callers can detect the empty-input case with
// errors.Is regardless of whether a default was supplied.
var ErrEmptyString = errors.New("g: empty input string")

// ErrInvalidBool is returned by StringToBool for a non-empty string that is
// not a recognised boolean literal.
var ErrInvalidBool = errors.New("g: invalid boolean string")

// StringToBool converts a string to a boolean.
// It handles various string representations of boolean values such as
// "true", "false", "yes", "on", "off". On failure it returns the supplied
// default value (or false if none) together with a non-nil error; the error
// is ErrEmptyString for an empty input and ErrInvalidBool for an unrecognised
// string, both matchable with errors.Is.
//
// Example Usage:
//
//	b, err := StringToBool("true") // true, nil
//	b, err := StringToBool("yes")  // true, nil
//	b, err := StringToBool("abc", false) // false, ErrInvalidBool
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
		return d, ErrEmptyString
	default:
		return d, ErrInvalidBool
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
