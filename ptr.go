package g

// Ptr creates a pointer from a literal or defaults to a pointer to a zero
// value of T if no arguments are given. This function is particularly
// useful when a pointer to a literal or default zero value is needed
// directly in expressions or function calls, simplifying syntax and
// avoiding the need for temporary variables.
//
// Parameters:
//
//	v ...T - An optional variable of type T from which to create a pointer.
//	         If not provided, the function returns a pointer to the zero
//	         value of T.
//
// Returns:
//
//	*T - A pointer to the value of T provided, or to the zero value of T
//	     if none is provided.
//
// Example usage:
//
//	// Function that returns an int value.
//	func Sum(a, b int) int { return a + b }
//
//	// Function that requires a pointer to an int.
//	func IsMoreThanTen(n *int) bool { return *n > 10 }
//
//	// Classical usage with variable.
//	v := Sum(3, 7)
//	r1 := IsMoreThanTen(&v)
//
//	// Using Ptr to simplify passing a pointer to a function.
//	r2 := IsMoreThanTen(Ptr(Sum(3, 7)))
//
//	// Create a pointer from a literal.
//	r3 := IsMoreThanTen(Ptr(21))
//
//	// Get a pointer to the zero value of the specified type.
//	zeroPtr := Ptr[int]()
//
//	// Specify the type of the literal.
//	var int64Ptr *int64 = Ptr[int64](21)
func Ptr[T any](v ...T) *T {
	if len(v) == 0 {
		var r T
		return &r
	}

	r := any(v[0]).(T)
	return &r
}
