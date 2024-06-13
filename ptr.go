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

// PtrIf conditionally creates a pointer to a given value or returns nil
// based on a boolean expression. This function simplifies pointer
// management in conditional logic, avoiding the need for temporary
// variables or manual pointer handling.
//
// Parameters:
//
//	exp bool - A boolean expression that determines whether a pointer
//	           is returned or nil.
//	v ...T   - An optional variable of type T from which to create a
//	           pointer. If the expression is false, or no value is
//	           provided, the function returns nil.
//
// Returns:
//
//	*T - A pointer to the value of T if the expression is true and a value
//	     is provided, otherwise nil.
//
// Example usage:
//
//	// where returns a WHERE clause for a SQL query based on the provided
//	// boolean values. It accepts three boolean pointers: isActive, isStaff,
//	// and isSuperuser. If any of these pointers are nil, the corresponding
//	// condition is not included in the WHERE clause.
//	func where(isActive, isStaff, isSuperuser *bool) string {
//		check := [...]struct {
//			name  string
//			value *bool
//		}{
//			{name: "is_active", value: isActive},
//			{name: "is_staff", value: isStaff},
//			{name: "is_superuser", value: isSuperuser},
//		}
//
//	    and := make([]string, 0, len(check))
//	    for _, m := range check {
//	    	if m.value != nil {
//	    		and = append(and, fmt.Sprintf("%s=%t", m.name, *m.value))
//	    	}
//	    }
//
//	    if len(and) != 0 {
//	    	return "WHERE " + strings.Join(and, ", ")
//	    }
//
//		return ""
//	}
//
//	// Some data is stored in map.
//	argsMap := make(map[string]bool)
//	argsMap["isActive"] = true
//	argsMap["isStaff"] = false
//
//	// Using PtrIf to conditionally pass pointers.
//	isActive, isActiveOk := argsMap["isActive"]
//	isStaff, isStaffOk := argsMap["isStaff"]
//	isSuperuser, isSuperuserOk := argsMap["isSuperuser"]
//	query := where(
//	    g.PtrIf(isActiveOk, isActive),
//	    g.PtrIf(isStaffOk, isStaff),
//	    g.PtrIf(isSuperuserOk, isSuperuser), // nil pointer
//	) // Result: "WHERE is_active=true, is_staff=false"
func PtrIf[T any](exp bool, v ...T) *T {
	if !exp {
		return nil
	}

	return Ptr[T](v...)
}
