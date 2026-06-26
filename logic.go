package g

import (
	"reflect"

	"github.com/goloop/trit"
)

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
// Example usage:
//
//	// Condition is true.
//	max := g.If(3 > 2, 3, 2)  // Output: 3
//
//	// If condition is false.
//	max := g.If(2 > 3, 2, 3)  // Output: 3
//
//	// Using with strings.
//	greeting := g.If(user == "admin", "Hello, admin", "Hello, user")
//	fmt.Println(greeting) // Output: the appropriate greeting
//
// The function can work with Trit types, so it can get the status Unknown.
// In this case, if the function receives two possible solutions, it will
// return the second one (false):
//
//	max := g.If(trit.Unknown, 3, 2)  // Output: 2
//
// This function can accept an optional argument u, for cases of working
// with Trit types. In this case, if Trit is Unknow, the third value will
// be returned:
//
//	max := g.If(trit.Unknown, 3, 2, 1)  // Output: 1
//
// If more than 3 possible result values are passed, the others of
// the results will be ignored. The third and more likely results
// are ignored for boolean expression values.
//
//	max := g.If(trit.Unknown, 3, 2, 1, 5, 7)  // Output: 1 and 5, 7 are ignored
//	min := g.If(3 > 2, 3, 2, 1, 5, 7)  // Output: 3 and 1, 5, 7 are ignored
//
// Important: Go is not a lazy language, so both branches t and f are
// fully evaluated at the call site before If runs. This is unlike a real
// ternary operator. For example, g.If(p != nil, *p, 0) dereferences p
// even when p is nil, which panics. When a branch is expensive or unsafe
// to evaluate unconditionally, use IfFunc, which evaluates only the
// chosen branch.
func If[L bool | trit.Trit, T any](e L, t, f T, u ...T) T {
	switch val := interface{}(e).(type) {
	case trit.Tritter:
		if val.IsTrue() {
			return t
		} else if len(u) != 0 && val.IsUnknown() {
			return u[0]
		}
	case bool:
		if val {
			return t
		}
	}

	return f
}

// IfFunc is a lazy variant of If: it evaluates only the branch it returns.
//
// Unlike If, which (because Go is not lazy) evaluates both branch values
// at the call site, IfFunc takes the branches as thunks and calls exactly
// one of them. This makes it safe and cheap when a branch is expensive or
// must not run unless selected.
//
// Example usage:
//
//	// Only parse(s) runs when s is non-empty; fallback() never runs.
//	v := g.IfFunc(s != "",
//	    func() int { return parse(s) },
//	    func() int { return fallback() },
//	)
//
//	// Safe nil-guard: *p is dereferenced only when p != nil.
//	val := g.IfFunc(p != nil,
//	    func() int { return *p },
//	    func() int { return 0 },
//	)
func IfFunc[T any](cond bool, t, f func() T) T {
	if cond {
		return t()
	}

	return f()
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
//	allNonZero := g.All(1, 2, 3)
//	fmt.Println(allNonZero) // Output: true
//
//	someZero := g.All(1, 0, 3)
//	fmt.Println(someZero) // Output: false
//
//	allNonZeroMixed := g.All(1, "a", true)
//	fmt.Println(allNonZeroMixed) // Output: true
//
//	empty := g.All()
//	fmt.Println(empty) // Output: false
//
// Warning: the function checks the list as the whole object, that is:
//
//	l := []bool{false, false, false}
//	g.All(l)    // Returns: true, because list is not an empty
//	            // and not an empty list is true
//	g.All(l...) // Returns: false, because not all elements
//	            // of the list are true
func All[T any](v ...T) bool {
	if len(v) == 0 {
		return false
	}

	for _, b := range v {
		if IsFalse(b) {
			return false
		}
	}

	return true
}

// AllList is a synonym for the All function that accepts
// a set of elements as a slice.
//
// This function prevents accidentally passing a slice
// as a value (whole object).
func AllList[T any](v []T) bool {
	return All(v...)
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
//	resultI := g.Any(ints...)
//	fmt.Println(resultI) // Output: true
//
//	// Check if any element in a slice of strings is non-empty.
//	strings := []string{"", "hello", "", ""}
//	resultS := g.Any(strings...)
//	fmt.Println(resultS) // Output: true
//
//	// Check if any element in a slice of booleans is true.
//	bools := []bool{false, false, true, false}
//	resultB := g.Any(bools...)
//	fmt.Println(resultB) // Output: true
//
// Warning: the function checks the list as the whole object, that is:
//
//	l := []bool{false, false, false}
//	g.Any(l)    // Returns: true, because list is not an empty
//	            // and not an empty list is true
//	g.Any(l...) // Returns: false, because not all elements
//	            // of the list are true
func Any[T any](v ...T) bool {
	for _, b := range v {
		if !IsFalse(b) {
			return true
		}
	}

	return false
}

// AnyList is a synonym for the Any function that accepts
// a set of elements as a slice.
//
// This function prevents accidentally passing a slice
// as a value (whole object).
func AnyList[T any](v []T) bool {
	return Any(v...)
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
//	result := g.IsEmpty(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is zero.
//	var f float64
//	result := g.IsEmpty(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a pointer variable is nil.
//	var ptr *int
//	result := g.IsEmpty(ptr)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is empty.
//	var str string
//	result := g.IsEmpty(str)
//	fmt.Println(result) // Output: true
//
//	// Check if a boolean variable is false.
//	var flag bool
//	result := g.IsEmpty(flag)
//	fmt.Println(result) // Output: true
func IsEmpty[T any](v T) bool {
	// For Trit the empty state is Unknown only.
	if t, ok := any(v).(trit.Tritter); ok {
		return t.IsUnknown()
	}

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return true // untyped nil interface
	}

	// For slices and arrays we check if the length is zero.
	if k := rv.Kind(); k == reflect.Slice || k == reflect.Array {
		return rv.Len() == 0
	}

	return rv.IsZero()
}

// IsFalse checks if the v of any type T is a false (zero) value for that type.
//
// P.s. trit.False and trit.Unknown are false values.
func IsFalse[T any](v T) bool {
	// For Trit not true states are: False and Unknown.
	// Anything that is not true is false: when converting a three-valued
	// object to a boolean object the Unknown is False.
	if t, ok := any(v).(trit.Tritter); ok {
		return !t.IsTrue()
	}

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return true // untyped nil interface
	}

	// For slices and arrays we check if the length is zero.
	if k := rv.Kind(); k == reflect.Slice || k == reflect.Array {
		return rv.Len() == 0
	}

	return rv.IsZero()
}

// IsTrue checks if the v of any type T is a true (non-zero) value for that type.
//
// P.s. trit.False and trit.Unknown are false values.
func IsTrue[T any](v T) bool {
	// Anything that is not a true is a false, when converting
	// three-valued object to a boolean object the Unknown is False.
	return !IsFalse(v)
}

// IsPointer checks if a value is a pointer.
//
// The function takes a value v of any type and checks if it is a pointer
// type. It returns true if v is a pointer, and false otherwise.
//
// Example usage:
//
//	// Check if a variable holding a pointer to a string is a pointer.
//	str := "hello"
//	result := g.IsPointer(&str)
//	fmt.Println(result) // Output: true
//
//	// Check if a variable holding a string is a pointer.
//	result = g.IsPointer(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding an integer is a pointer.
//	result = g.IsPointer(10)
//	fmt.Println(result) // Output: false
//
//	// Check if a variable holding a typed nil pointer is a pointer.
//	var ptr *int
//	result = g.IsPointer(ptr)
//	fmt.Println(result) // Output: true
func IsPointer(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.IsValid() && rv.Kind() == reflect.Ptr
}

// IsNumber checks if a value is a numeric type.
//
// The function takes a value v of any type and checks if it is of a numeric
// type, including integer and floating-point types. It returns true if v
// is a numeric type, and false otherwise.
//
// Example usage:
//
//	// Check if an integer variable is a numeric type.
//	num := 10
//	result := g.IsNumber(num)
//	fmt.Println(result) // Output: true
//
//	// Check if a float variable is a numeric type.
//	f := 3.14
//	result = g.IsNumber(f)
//	fmt.Println(result) // Output: true
//
//	// Check if a string variable is a numeric type.
//	str := "hello"
//	result = g.IsNumber(str)
//	fmt.Println(result) // Output: false
//
//	// Check if a boolean variable is a numeric type.
//	flag := true
//	result = g.IsNumber(flag)
//	fmt.Println(result) // Output: false
//
//	// Check if a slice of integers is a numeric type.
//	nums := []int{1, 2, 3}
//	result = g.IsNumber(nums)
//	fmt.Println(result) // Output: false
func IsNumber(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8,
		uint16, uint32, uint64, uintptr, float32, float64:
		return true
	default:
		return false
	}
}
