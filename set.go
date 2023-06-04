package do

// Union takes two slices and returns a new slice that contains
// all unique items from both slices. The type T must be comparable.
//
// It removes duplicates and returns the union of the slices.
//
// The function is generic and can work with any type T that
// is comparable.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := do.Union(a, b)
//	fmt.Println(result) // Output: [1 2 3 4 5]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := do.Union(a, b)
//	fmt.Println(result) // Output: [a b c d e]
func Union[T comparable](a []T, b []T) []T {
	distinctValues := map[T]bool{}

	for _, val := range a {
		distinctValues[val] = true
	}

	for _, val := range b {
		distinctValues[val] = true
	}

	result := make([]T, 0, len(distinctValues))
	for val := range distinctValues {
		result = append(result, val)
	}

	return result
}

// Intersection takes two slices and returns a new slice that
// contains the common items present in both slices. The type T
// must be comparable.
//
// The function returns the intersection of the slices.
//
// The function is generic and can work with any type T that
// is comparable.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := do.Intersection(a, b)
//	fmt.Println(result) // Output: [3]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := do.Intersection(a, b)
//	fmt.Println(result) // Output: [c]
func Intersection[T comparable](a []T, b []T) []T {
	m1 := make(map[T]bool)
	for _, item := range a {
		m1[item] = true
	}
	m2 := make(map[T]bool)
	for _, item := range b {
		m2[item] = true
	}

	result := make([]T, 0)
	for item := range m1 {
		if m2[item] {
			result = append(result, item)
		}
	}

	return result
}

// Difference takes two slices and returns a new slice that
// contains the items present in the first slice but not in the
// second slice. The type T must be comparable.
//
// The function returns the difference of the slices.
//
// The function is generic and can work with any type T that
// is comparable.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := do.Difference(a, b)
//	fmt.Println(result) // Output: [1, 2]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := do.Difference(a, b)
//	fmt.Println(result) // Output: ["a", "b"]
func Difference[T comparable](a []T, b []T) []T {
	m1 := make(map[T]bool)
	for _, item := range a {
		m1[item] = true
	}
	m2 := make(map[T]bool)
	for _, item := range b {
		m2[item] = true
	}

	result := make([]T, 0)
	for item := range m1 {
		if !m2[item] {
			result = append(result, item)
		}
	}

	return result
}

// Diff is an alias for Difference function.
func Diff[T comparable](a []T, b []T) []T {
	return Difference(a, b)
}

// SymmetricDifference takes two slices and returns a new slice
// that contains the items present in one of the slices but not in both.
// The type T must be comparable.
//
// The function returns the symmetric difference of the slices.
//
// The function is generic and can work with any type T that
// is comparable.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := do.SymmetricDifference(a, b)
//	fmt.Println(result) // Output: [1, 2, 4, 5]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := do.SymmetricDifference(a, b)
//	fmt.Println(result) // Output: ["a", "b", "d", "e"]
func SymmetricDifference[T comparable](a []T, b []T) []T {
	m1 := make(map[T]bool)
	for _, item := range a {
		m1[item] = true
	}
	m2 := make(map[T]bool)
	for _, item := range b {
		m2[item] = true
	}

	result := make([]T, 0)
	for item := range m1 {
		if !m2[item] {
			result = append(result, item)
		}
	}
	for item := range m2 {
		if !m1[item] {
			result = append(result, item)
		}
	}

	return result
}

// Sdiff is an alias for SymmetricDifference function.
func Sdiff[T comparable](a []T, b []T) []T {
	return SymmetricDifference(a, b)
}

// Complement takes a universal set (b) and a subset of it (a) and
// Output: a new slice containing items present in the universal set
// but not in the subset. The type T must be comparable.
//
// The function returns the complement of the subset.
//
// The function is generic and can work with any type T that
// is comparable.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	u := []int{1, 2, 3, 4, 5}
//	a := []int{1, 2, 3}
//	result := do.Complement(a, u)
//	fmt.Println(result) // Output: [4, 5]
//
//	u := []string{"a", "b", "c", "d", "e"}
//	a := []string{"a", "b", "c"}
//	result := do.Complement(a, u)
//	fmt.Println(result) // Output: ["d", "e"]
func Complement[T comparable](a []T, b []T) []T {
	m1 := make(map[T]bool)
	for _, item := range a {
		m1[item] = true
	}

	result := make([]T, 0)
	for _, item := range b {
		if !m1[item] {
			result = append(result, item)
		}
	}

	return result
}

// CartesianProduct returns all possible pairs from two slices.
//
// The function generates a slice of pairs, where each pair
// consists of an element from the first input slice and an
// element from the second input slice. The length of the
// returned slice is equal to the product of the lengths of
// the input slices.
//
// This function is generic and can work with any type T.
//
// Note: This function does not preserve the order of elements.
// The order of elements in the returned slice can be different
// from the order of elements in the input slices.
//
// Example usage:
//
//	a := []int{1, 2}
//	b := []int{3, 4}
//	result := do.CartesianProduct(a, b)
//	// Output: [[1, 3], [1, 4], [2, 3], [2, 4]]
//
//	x := []string{"A", "B"}
//	y := []string{"C", "D"}
//	result := do.CartesianProduct(x, y)
//	// Output: [["A", "C"], ["A", "D"], ["B", "C"], ["B", "D"]]
func CartesianProduct[T any](a []T, b []T) [][2]T {
	var result [][2]T

	for _, itemA := range a {
		for _, itemB := range b {
			result = append(result, [2]T{itemA, itemB})
		}
	}

	return result
}
