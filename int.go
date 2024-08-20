package g

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
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
// uint, uint64,uint32, etc.
func IntToString[T constraints.Integer](v T) string {
	return fmt.Sprintf("%v", v)
}
