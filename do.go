// The 'do' package is a robust utility library for Go 1.20+ that brings
// advanced generic and type-safe helper functions to streamline various
// development tasks. This package employs Go's generics to their fullest
// potential, catering to an array of needs for modern software development.
//
// Key features of the 'do' package encompass easy-to-use conditionals,
// zero-value checks, numeric operations, functional programming aids, and
// element operations. Moreover, it extends its usefulness to more complex
// functionalities like pairing, ranking, looking up data in arrays and maps,
// and manipulating lists in diverse ways including sorting, shuffling, and
// reduction. Random number and element selection operations are also
// provided, amplifying the package's versatility.
//
// The 'do' package champions a design philosophy that encourages code
// clarity and promotes maintainability and reliability in Go applications.
// It is an ideal tool for developers seeking to write clean, concise, and
// efficient code while harnessing the power of Go generics.
package do

import (
	"math/rand"
	"time"
)

// Pair is a generic struct type with two fields: First and Second.
// Used as the result element of the Zip function and can contain
// a pair of values of any T and U types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Numerable is an interface type that is satisfied by all numeric types
// in Go, both integer and floating point. This includes int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, uint64, float32, and float64.
// It allows functions to operate on any of these types where numerical
// operations such as addition, subtraction, multiplication, and division
// are needed. It enables generic programming techniques for numeric types.
type Numerable interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

// Verifiable is an interface type that is satisfied by classical types
// like numeric types and strings in Go.
//
// The purpose of the Verifiable interface is to enable generic programming
// techniques for numeric types and strings. Functions can use this interface
// as a constraint to operate on any of these types where numerical operations
// or string operations are needed.
type Verifiable interface {
	Numerable | string | rune
}

// The randomGenerator is a global variable that is used by the Random function
// to generate random numbers. It is initialized in the init() function.
var randomGenerator *rand.Rand

// The init initializes the randomGenerator variable.
func init() {
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
}
