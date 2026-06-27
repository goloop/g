package g

import (
	"errors"
	"strconv"
)

// StringToInt converts a string to an integer.
// If the conversion fails and a default value is provided, it
// returns the default value. Otherwise, it returns an error.
//
// Example Usage:
//
//	i, err := StringToInt("41") // 41, nil
//	i, err := StringToInt("abc", 7) // 7, error
func StringToInt(v string, def ...int) (int, error) {
	var d int = 0

	if len(def) > 0 {
		d = def[0]
	}

	if v == "" {
		return d, errors.New("empty string and no default value")
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		return d, err
	}

	return i, nil
}

// IntToString converts an integer to a string.
// It handles different integer types such as int, int64, int32,
// uint, uint64, uint32, etc.
func IntToString[T Integer](v T) string {
	// Detect unsigned types generically: for an unsigned T, zero-1 wraps to
	// the type's maximum (> 0); for a signed T it is -1 (< 0). This picks the
	// correct formatter without reflection and works for named types too.
	var zero T
	if zero-1 > zero {
		return strconv.FormatUint(uint64(v), 10)
	}
	return strconv.FormatInt(int64(v), 10)
}
