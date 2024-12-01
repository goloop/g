package g

// Signed is a constraint that permits any signed integer type.
// If future releases of Go add new predeclared signed integer types,
// this constraint will be modified to include them.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint that permits any unsigned integer type.
// If future releases of Go add new predeclared unsigned integer types,
// this constraint will be modified to include them.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint that permits any integer type.
// If future releases of Go add new predeclared integer types,
// this constraint will be modified to include them.
type Integer interface {
	Signed | Unsigned
}

// Float is a constraint that permits any floating-point type.
// If future releases of Go add new predeclared floating-point types,
// this constraint will be modified to include them.
type Float interface {
	~float32 | ~float64
}

// Complex is a constraint that permits any complex numeric type.
// If future releases of Go add new predeclared complex numeric types,
// this constraint will be modified to include them.
type Complex interface {
	~complex64 | ~complex128
}

// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
// If future releases of Go add new ordered types,
// this constraint will be modified to include them.
type Ordered interface {
	Integer | Float | ~string
}

// Numerable is an interface type that is satisfied by all numeric types
// in Go, both integer and floating point. This includes int, int8, int16,
// int32, int64, uint, uint8, uint16, uint32, uint64, float32, and float64.
// It allows functions to operate on any of these types where numerical
// operations such as addition, subtraction, multiplication, and division
// are needed. It enables generic programming techniques for numeric types.
type Numerable interface {
	Integer | Float
}

// Verifiable is an interface type that is satisfied by classical types
// like numeric types and strings in Go.
//
// The purpose of the Verifiable interface is to enable generic programming
// techniques for numeric types and strings. Functions can use this interface
// as a constraint to operate on any of these types where numerical operations
// or string operations are needed.
type Verifiable interface {
	Integer | Float | string | rune
}
