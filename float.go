package g

import (
	"math"
	"strconv"
)

// StringToFloat converts a string to a float.
// On failure it returns the supplied default value (or 0 if none) together
// with a non-nil error: ErrEmptyString for an empty input, otherwise the
// strconv error (matchable with errors.Is against strconv.ErrSyntax or
// strconv.ErrRange).
//
// Example Usage:
//
//	f, err := StringToFloat("3.14") // 3.14, nil
//	f, err := StringToFloat("abc", 1.23) // 1.23, error
func StringToFloat(v string, def ...float64) (float64, error) {
	var d float64 = 0

	if len(def) > 0 {
		d = def[0]
	}

	if v == "" {
		return d, ErrEmptyString
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return d, err
	}

	return f, nil
}

// FloatToString converts a float to a string.
// It handles different floating-point types such as float32 and float64,
// using the shortest representation that round-trips for the value's width.
func FloatToString[T ~float32 | ~float64](v T) string {
	bitSize := 64
	// A ~float32 type cannot hold math.MaxFloat64, so the conversion overflows
	// to +Inf. This detects the type's width rather than the dynamic type, so
	// a named ~float32 (type Celsius float32) is also formatted at 32-bit width
	// - a type switch on any(v) would miss it and use 64-bit width. The bound
	// goes through a variable so the conversion is a runtime overflow (to +Inf)
	// rather than a compile-time constant-overflow error.
	maxF64 := math.MaxFloat64
	if math.IsInf(float64(T(maxF64)), 1) {
		bitSize = 32
	}
	return strconv.FormatFloat(float64(v), 'g', -1, bitSize)
}
