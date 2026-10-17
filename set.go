package g

// Union takes two slices and returns a new slice that contains
// all unique items from both slices. The type T must be comparable.
//
// It removes duplicates and returns the union of the slices.
//
// The function is generic and can work with any type T that
// is comparable.
//
// The result preserves first-seen order: elements appear in the order they
// are first encountered while scanning a and then b.
//
// Example usage:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := g.Union(a, b)
//	fmt.Println(result) // Output: [1 2 3 4 5]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := g.Union(a, b)
//	fmt.Println(result) // Output: [a b c d e]
func Union[T comparable](a []T, b []T) []T {
	seen := make(map[T]struct{}, len(a)+len(b))
	result := make([]T, 0, len(a)+len(b))

	for _, val := range a {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			result = append(result, val)
		}
	}
	for _, val := range b {
		if _, ok := seen[val]; !ok {
			seen[val] = struct{}{}
			result = append(result, val)
		}
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
//	result := g.Intersection(a, b)
//	fmt.Println(result) // Output: [3]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := g.Intersection(a, b)
//	fmt.Println(result) // Output: [c]
func Intersection[T comparable](a []T, b []T) []T {
	setB := make(map[T]struct{}, len(b))
	for _, item := range b {
		setB[item] = struct{}{}
	}

	result := make([]T, 0)
	for _, item := range a {
		if _, ok := setB[item]; !ok {
			continue
		}
		delete(setB, item)
		result = append(result, item)
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
//	result := g.Difference(a, b)
//	fmt.Println(result) // Output: [1, 2]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := g.Difference(a, b)
//	fmt.Println(result) // Output: ["a", "b"]
func Difference[T comparable](a []T, b []T) []T {
	setB := make(map[T]struct{}, len(b))
	for _, item := range b {
		setB[item] = struct{}{}
	}

	seen := make(map[T]struct{}, len(a))
	result := make([]T, 0)
	for _, item := range a {
		if _, inB := setB[item]; inB {
			continue
		}
		if _, dup := seen[item]; dup {
			continue
		}
		seen[item] = struct{}{}
		result = append(result, item)
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
//	result := g.SymmetricDifference(a, b)
//	fmt.Println(result) // Output: [1, 2, 4, 5]
//
//	a := []string{"a", "b", "c"}
//	b := []string{"c", "d", "e"}
//	result := g.SymmetricDifference(a, b)
//	fmt.Println(result) // Output: ["a", "b", "d", "e"]
func SymmetricDifference[T comparable](a []T, b []T) []T {
	setA := make(map[T]struct{}, len(a))
	for _, item := range a {
		setA[item] = struct{}{}
	}
	setB := make(map[T]struct{}, len(b))
	for _, item := range b {
		setB[item] = struct{}{}
	}

	result := make([]T, 0, len(a)+len(b))
	for item := range setA {
		if _, ok := setB[item]; !ok {
			result = append(result, item)
		}
	}
	for item := range setB {
		if _, ok := setA[item]; !ok {
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
//	result := g.Complement(a, u)
//	fmt.Println(result) // Output: [4, 5]
//
//	u := []string{"a", "b", "c", "d", "e"}
//	a := []string{"a", "b", "c"}
//	result := g.Complement(a, u)
//	fmt.Println(result) // Output: ["d", "e"]
func Complement[T comparable](a []T, b []T) []T {
	setA := make(map[T]struct{}, len(a))
	for _, item := range a {
		setA[item] = struct{}{}
	}

	result := make([]T, 0)
	for _, item := range b {
		if _, ok := setA[item]; !ok {
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
//	result := g.CartesianProduct(a, b)
//	// Output: [[1, 3], [1, 4], [2, 3], [2, 4]]
//
//	x := []string{"A", "B"}
//	y := []string{"C", "D"}
//	result := g.CartesianProduct(x, y)
//	// Output: [["A", "C"], ["A", "D"], ["B", "C"], ["B", "D"]]
func CartesianProduct[T any](a []T, b []T) [][2]T {
	result := make([][2]T, 0, len(a)*len(b))

	for _, itemA := range a {
		for _, itemB := range b {
			result = append(result, [2]T{itemA, itemB})
		}
	}

	return result
}
