package g

import (
	"strconv"
)

// StringToInt converts a string to an integer.
// On failure it returns the supplied default value (or 0 if none) together
// with a non-nil error: ErrEmptyString for an empty input, otherwise the
// strconv error (matchable with errors.Is against strconv.ErrSyntax or
// strconv.ErrRange).
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
		return d, ErrEmptyString
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
