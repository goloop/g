package do

import "reflect"

// If is a substitute for the ternary operator (?:)
// which is not available in Go.
//
// In languages like C/C++ and Python, you can use a ternary
// operator for a concise conditional expression:
//
//	C/C++:  int max = (a > b) ? a : b;
//	Python: max = a if a > b else b
//
// It takes three parameters: a boolean expression e, and two
// values of any type T (t and f). If the expression e is true,
// it returns t, otherwise it returns f.
//
// Example:
//
//	var a, b float64 // or int, or any other type
//	...
//	max := do.If(a > b, a, b)
//
// This function is generic and can work with any type T.
func If[T any](e bool, t, f T) T {
	if e {
		return t
	}

	return f
}

// All returns true if all values in the provided slice
// are not zero values for their types.
//
// If at least one value is a zero value, it immediately returns false.
// If the slice is empty, it returns false.
//
// This function is generic and can work with any type T.
func All[T any](v ...T) bool {
	if len(v) == 0 {
		return false
	}

	for _, val := range v {
		if IsEmpty(val) {
			return false
		}
	}

	return true
}

// Any returns true if at least one value in the provided slice
// is not a zero value for its type.
//
// As soon as it finds a value that is not a zero value, it returns true.
// If all values in the slice are zero values or the slice is empty,
// it returns false.
//
// This function is generic and can work with any type T.
func Any[T any](v ...T) bool {
	for _, val := range v {
		if !IsEmpty(val) {
			return true
		}
	}

	return false
}

// IsEmpty checks if the  v of any type T is "zero value" for that type.
//
// Zero values in Go are values that the variables of respective types hold
// upon their declaration, if they do not have any explicit initialization.
//
// For example, zero value of type int is 0, for type float64 is 0.0,
// for a pointer is nil, for a string is "", for a boolean is false, etc.
//
// This function is generic and can work with any type T.
func IsEmpty[T any](v T) bool {
	t := reflect.TypeOf(v)
	if t == nil {
		return true
	}
	zero := reflect.Zero(t).Interface()
	return reflect.DeepEqual(v, zero)
}

// IsPointer checks if a value is a pointer.
//
// The function takes an interface{} value `v` and checks if it is
// a pointer type. It returns true if `v` is a pointer, and false
// otherwise.
//
// Example:
//
//	str := "hello"
//	isPtr := IsPointer(&str)  // true
//	isPtr = IsPointer(str)    // false
//	isPtr = IsPointer(10)     // false
//	isPtr = IsPointer(nil)    // false
func IsPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

// IsNumber checks if a value is a numeric type.
//
// The function takes an interface{} value `v` and checks if it is
// of a numeric type, including integer and floating-point types.
// It returns true if `v` is a numeric type, and false otherwise.
//
// Example:
//
//	isNum := IsNumber(10)       // true
//	isNum = IsNumber(3.14)      // true
//	isNum = IsNumber("hello")   // false
//	isNum = IsNumber(true)      // false
//	isNum = IsNumber([]int{1})  // false
func IsNumber(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8,
		uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
