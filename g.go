package g

// Pair is a generic struct type with two fields: First and Second.
// Used as the result element of the Zip function and can contain
// a pair of values of any T and U types.
type Pair[T, U any] struct {
	First  T
	Second U
}

// MaxRangeSize is the maximum size for range generation
// by the Range function.
const MaxRangeSize = 100_000_000
