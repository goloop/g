package g

import "testing"

// TestRank tests the Rank function.
func TestRank(t *testing.T) {
	rank := Rank(7, []int{1, 5, 2, 3, 7, 8})
	if rank != 1 {
		t.Errorf("Expected rank 1, but got %d", rank)
	}

	rank = Rank(7, []int{1, 5, 2, 3, 7, 8}, true)
	if rank != 4 {
		t.Errorf("Expected rank 4, but got %d", rank)
	}

	rank = Rank(9, []int{1, 5, 2, 3, 7, 8})
	if rank != -1 {
		t.Errorf("Expected rank -1, but got %d", rank)
	}

	rank = Rank(7, []int{1, 7, 7, 3, 7, 8})
	if rank != 1 {
		t.Errorf("Expected rank 1, but got %d", rank)
	}

	rank = Rank(7, []int{1, 7, 7, 3, 7, 8}, true)
	if rank != 2 {
		t.Errorf("Expected rank 2, but got %d", rank)
	}

	rank = Rank(4.5, []float64{1.2, 3.1, 4.5, 2.8, 4.5, 6.7})
	if rank != 1 {
		t.Errorf("Expected rank 1, but got %d", rank)
	}

	rank = Rank(4.5, []float64{1.2, 3.1, 4.5, 2.8, 4.5, 6.7}, true)
	if rank != 3 {
		t.Errorf("Expected rank 3, but got %d", rank)
	}

	rank = Rank("Banana", []string{
		"Apple", "Banana", "Cherry",
		"Banana", "Dragonfruit", "Elderberry",
	})
	if rank != 3 {
		t.Errorf("Expected rank 3, but got %d", rank)
	}

	rank = Rank("Banana", []string{
		"Apple", "Banana", "Cherry",
		"Banana", "Dragonfruit", "Elderberry",
	}, true)
	if rank != 1 {
		t.Errorf("Expected rank 1, but got %d", rank)
	}
}

// TestHLookup tests the HLookup function.
func TestHLookup(t *testing.T) {
	lookup := []string{"A", "B", "C"}
	result := []int{1, 2, 3}

	// Test case 1: Lookup value exists in the table.
	val := HLookup("B", lookup, result, -1)
	expected := 2
	if val != expected {
		t.Errorf("HLookup test case 1 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 2: Lookup value does not exist in the table.
	val = HLookup("D", lookup, result, -1)
	expected = -1
	if val != expected {
		t.Errorf("HLookup test case 2 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 3: Lookup value is the first element in the table.
	val = HLookup("A", lookup, result, -1)
	expected = 1
	if val != expected {
		t.Errorf("HLookup test case 3 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 4: Lookup value is the last element in the table.
	val = HLookup("C", lookup, result, -1)
	expected = 3
	if val != expected {
		t.Errorf("HLookup test case 4 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 5: Lookup value is not found,
	// with a non-default default value.
	val = HLookup("E", lookup, result, 0)
	expected = 0
	if val != expected {
		t.Errorf("HLookup test case 5 failed: expected %d, got %d",
			expected, val)
	}
}

// TestVLookup tests the VLookup function.
func TestVLookup(t *testing.T) {
	lookup := []string{"A", "B", "C"}
	result := []int{1, 2, 3}

	// Test case 1: Lookup value exists in the table.
	val := VLookup("B", lookup, result, -1)
	expected := 2
	if val != expected {
		t.Errorf("VLookup test case 1 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 2: Lookup value does not exist in the table.
	val = VLookup("D", lookup, result, -1)
	expected = -1
	if val != expected {
		t.Errorf("VLookup test case 2 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 3: Lookup value is the first element in the table.
	val = VLookup("A", lookup, result, -1)
	expected = 1
	if val != expected {
		t.Errorf("VLookup test case 3 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 4: Lookup value is the last element in the table.
	val = VLookup("C", lookup, result, -1)
	expected = 3
	if val != expected {
		t.Errorf("VLookup test case 4 failed: expected %d, got %d",
			expected, val)
	}

	// Test case 5: Lookup value is not found,
	// with a non-default default value.
	val = VLookup("E", lookup, result, 0)
	expected = 0
	if val != expected {
		t.Errorf("VLookup test case 5 failed: expected %d, got %d",
			expected, val)
	}
}
