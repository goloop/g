package g

// Rank function returns the rank of a value when compared to a list of
// other values. Rank can rank values from largest to smallest (i.e. top sales)
// as well as smallest to largest (i.e. fastest time).
//
// It is not necessary to sort the values in the list before using Rank.
//
// The function has two modes of operation, controlled by the
// ascending argument. To rank values where the largest value is ranked #1,
// set ascending to false (or leave blank).
//
//	g.Rank(7, []float64{1, 5, 2, 3, 7, 8})       // descending, returns  1
//	g.Rank(7, []float64{1, 5, 2, 3, 7, 8}, true) // ascending, returns 4
//
// Set ascending to false when you want to rank something like top sales,
// where the largest sales number should rank #1, and to set ascending to
// true when you want to rank something like race results, where the
// shortest (fastest) time should rank #1.
//
// The Rank function will assign duplicate values to the same rank.
// For example, if a certain value has a rank of 3, and there are two
// instances of the value in the data, the Rank function will assign both
// instances a rank of 3. The next rank assigned will be 5, and no value
// will be assigned a rank of 4. If tied ranks are a problem, one workaround
// is to employ a tie-breaking strategy.
//
// The function works with both numbers and strings, all that satisfy
// the Verifiable interface.
//
// Example usage:
//
//	// Rank a number in a descending list.
//	result := g.Rank(7, []float64{1, 5, 2, 3, 7, 8})
//	fmt.Println(result) // Output: 1
//
//	// Rank a number in an ascending list.
//	result = g.Rank(7, []float64{1, 5, 2, 3, 7, 8}, true)
//	fmt.Println(result) // Output: 4
//
//	// Rank a number that doesn't exist in the list.
//	result = g.Rank(9, []float64{1, 5, 2, 3, 7, 8}, true)
//	fmt.Println(result) // Output: -1
func Rank[T Verifiable](number T, array []T, ascending ...bool) int {
	var removed int

	// Method remove sample element from array and returns
	// number of removed elements and new array.
	remove := func(array []T, sample T) (removed int, result []T) {
		for _, v := range array {
			if v == sample {
				removed++
			} else {
				result = append(result, v)
			}
		}

		return
	}

	rank := 0
	inverse := All(ascending...)
	for len(array) != 0 {
		var value T = If(inverse, MinList(array), MaxList(array))
		if value == number {
			return int(rank)
		}

		removed, array = remove(array, value)
		rank += removed
	}

	return -1
}

// HLookup looks up and retrieves data from a specific row in a table.
//
// The function takes a search value `v`, a slice of lookup values `lookup`,
// a slice of result values `result`, and an optional default value `def`.
// It searches for the first occurrence of `v` in the `lookup` slice and
// Output: the corresponding value from the `result` slice. If `v` is not
// found in the `lookup` slice, it returns the default value `def`.
//
// Example usage:
//
//	// Perform a horizontal lookup on a string slice and retrieve the
//	// corresponding value from an int slice.
//	lookup := []string{"A", "B", "C"}
//	result := []int{1, 2, 3}
//	val := g.HLookup("B", lookup, result, -1)
//	fmt.Println(val) // Output: 2
//
//	// Perform a horizontal lookup on a string slice with a value that
//	// doesn't exist, and return the default value.
//	val = g.HLookup("D", lookup, result, -1)
//	fmt.Println(val) // Output: -1
func HLookup[T comparable, U any](v T, lookup []T, result []U, def U) U {
	for i, item := range lookup {
		if item == v {
			return result[i]
		}
	}

	return def
}

// VLookup looks up and retrieves data from a specific column in a table.
//
// The function takes a search value `v`, a slice of lookup values `lookup`,
// a slice of result values `result`, and an optional default value `def`.
// It searches for the first occurrence of `v` in the `lookup` slice and
// Output: the corresponding value from the `result` slice. If `v` is not
// found in the `lookup` slice, it returns the default value `def`.
//
// Example usage:
//
//	// Perform a vertical lookup on a string slice and retrieve the
//	// corresponding value from an int slice.
//	lookup := []string{"A", "B", "C"}
//	result := []int{1, 2, 3}
//	val := g.VLookup("B", lookup, result, -1)
//	fmt.Println(val) // Output: 2
//
//	// Perform a vertical lookup on a string slice with a value that
//	// doesn't exist, and return the default value.
//	val = g.VLookup("D", lookup, result, -1)
//	fmt.Println(val) // Output: -1
func VLookup[T comparable, U any](v T, lookup []T, result []U, def U) U {
	for i, item := range lookup {
		if item == v {
			return result[i]
		}
	}

	return def
}
