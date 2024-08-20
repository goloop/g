package g

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

// StringToFloat converts a string to a float.
// If the conversion fails and a default value is provided, it
// returns the default value. Otherwise, it returns an error.
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
		return d, errors.New("empty string and no default value")
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return d, err
	}

	return f, nil
}

// FloatToString converts a float to a string.
// It handles different floating-point types such as float32 and float64.
func FloatToString[T constraints.Float](v T) string {
	return fmt.Sprintf("%v", v)
}
