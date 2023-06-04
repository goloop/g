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
//	// Condition is true.
//	max := do.If(3 > 2, 3, 2)  // Output: 3
//
//	// If condition is false.
//	max := do.If(2 > 3, 2, 3)  // Output: 3
//
//	// Using with strings.
//	greeting := do.If(user == "admin", "Hello, admin", "Hello, user")
//	fmt.Println(greeting) // Output: the appropriate greeting
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
//
// Example usage:
//
//	allNonZero := do.All(1, 2, 3)
//	fmt.Println(allNonZero) // Output: true
//
//	someZero := do.All(1, 0, 3)
//	fmt.Println(someZero) // Output: false
//
//	allNonZeroMixed := do.All(1, "a", true)
//	fmt.Println(allNonZeroMixed) // Output: true
//
//	empty := do.All()
//	fmt.Println(empty) // Output: false
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
// Example usage:
//
//	// Check if any element in a slice of integers is non-zero.
//	ints := []int{0, 0, 0, 1, 0}
//	resultI := do.Any(ints...)
//	fmt.Println(resultI) // Output: true
//
//	// Check if any element in a slice of strings is non-empty.
//	strings := []string{"", "hello", "", ""}
//	resultS := do.Any(strings...)
//	fmt.Println(resultS) // Output: true
//
//	// Check if any element in a slice of booleans is true.
//	bools := []bool{false, false, true, false}
//	resultB := do.Any(bools...)
//	fmt.Println(resultB) // Output: true
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
// Example usage:
//
//	// Check if an integer variable is zero.
//	var num int
//	result := do.IsEmpty(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is zero.
//	var f float64
//	result := do.IsEmpty(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a pointer variable is nil.
//	var ptr *int
//	result := do.IsEmpty(ptr)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is empty.
//	var str string
//	result := do.IsEmpty(str)
//	fmt.Println(result) // Output: true
//
//	// Check if a boolean variable is false.
//	var flag bool
//	result := do.IsEmpty(flag)
//	fmt.Println(result) // Output: true
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
// Example usage:
//
//	// Check if a variable holding a pointer to a string is a pointer.
//	str := "hello"
//	result := do.IsPointer(&str)
//	fmt.Println(result) // Output: true
//
//	// Check if a variable holding a string is a pointer.
//	result = do.IsPointer(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding an integer is a pointer.
//	result = do.IsPointer(10)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding nil is a pointer.
//	var ptr *int
//	result = do.IsPointer(ptr)
//	fmt.Println(result) // Output: false
func IsPointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Ptr
}

// IsNumber checks if a value is a numeric type.
//
// The function takes an interface{} value `v` and checks if it is
// of a numeric type, including integer and floating-point types.
// It returns true if `v` is a numeric type, and false otherwise.
//
// Example usage:
//
//	// Check if an integer variable is a numeric type.
//	num := 10
//	result := do.IsNumber(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is a numeric type.
//	f := 3.14
//	result = do.IsNumber(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is a numeric type.
//	str := "hello"
//	result = do.IsNumber(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a boolean variable is a numeric type.
//	flag := true
//	result = do.IsNumber(flag)
//	fmt.Println(result) // Output: false
//
//	// Check if a slice of integers is a numeric type.
//	nums := []int{1, 2, 3}
//	result = do.IsNumber(nums)
//	fmt.Println(result) // Output: false
func IsNumber(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8,
		uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
