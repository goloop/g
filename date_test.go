package g

import (
	"reflect"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestFindDuplicatesDateTimeFormats tests dataTimeFormats.
func TestFindDuplicatesDateTimeFormats(t *testing.T) {
	seen := make(map[string]bool)
	for _, format := range dataTimeFormats {
		if seen[format] {
			t.Errorf("Duplicate format found: %s", format)
		}

		seen[format] = true
	}
}

func TestStringToDate_ImpossibleState(t *testing.T) {
	_, err := StringToDate("invalid date")

	if err.Error() == "failed to parse date" {
		t.Error("Got impossible error state: " +
			"'failed to parse date' without error list")
	}

	if !strings.Contains(err.Error(), "failed to parse date:") {
		t.Errorf("Expected error with parse errors list, got: %v", err)
	}
}

// TestDateFoundValue_SetValue tests SetValue method of dateFoundValue.
func TestDateFoundValue_SetValue(t *testing.T) {
	dfv := &dateFoundValue{}
	expectedTime := time.Date(2023, 12, 1, 15, 30, 45, 0, time.UTC)

	// Test setting true with time
	dfv.SetValue(true, expectedTime)
	if !dfv.found {
		t.Error("Expected found to be true")
	}
	if dfv.value != expectedTime {
		t.Errorf("Expected time to be %v, got %v", expectedTime, dfv.value)
	}

	// Test setting false with zero time
	zeroTime := time.Time{}
	dfv.SetValue(false, zeroTime)
	if dfv.found {
		t.Error("Expected found to be false")
	}
	if dfv.value != zeroTime {
		t.Errorf("Expected time to be %v, got %v", zeroTime, dfv.value)
	}
}

func TestDateFoundValue_GetValue(t *testing.T) {
	dfv := &dateFoundValue{}
	expectedTime := time.Date(2023, 12, 1, 15, 30, 45, 0, time.UTC)

	// Test getting value when not found
	_, err := dfv.GetValue()
	if err == nil {
		t.Error("Expected error when value not found")
	}
	if err.Error() != "date and time could not be recognized" {
		t.Errorf(
			"Expected error message 'date and time could not be "+
				"recognized',got '%s'", err.Error())
	}

	// Test getting value when found
	dfv.SetValue(true, expectedTime)
	value, err := dfv.GetValue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != expectedTime {
		t.Errorf("Expected time to be %v, got %v", expectedTime, value)
	}
}

func TestDateFoundValue_Concurrent(t *testing.T) {
	dfv := &dateFoundValue{}
	expectedTime := time.Date(2023, 12, 1, 15, 30, 45, 0, time.UTC)
	concurrentAccess := 100
	var wg sync.WaitGroup

	// Test concurrent access
	for i := 0; i < concurrentAccess; i++ {
		wg.Add(2)

		// Concurrent SetValue
		go func() {
			defer wg.Done()
			dfv.SetValue(true, expectedTime)
		}()

		// Concurrent GetValue
		go func() {
			defer wg.Done()
			dfv.GetValue()
		}()
	}

	wg.Wait()

	// Verify final state
	value, err := dfv.GetValue()
	if err != nil {
		t.Errorf("Unexpected error after concurrent access: %v", err)
	}
	if value != expectedTime {
		t.Errorf("Expected time to be %v after concurrent access, got %v",
			expectedTime, value)
	}
}

// TestStringToDate tests StringToDate function.
func TestStringToDate(t *testing.T) {
	minLoadPerGoroutine = 4
	ParallelTasks(2)
	tests := []struct {
		input    string
		patterns []string
		want     time.Time
		wantErr  bool
	}{
		{
			"7.07.2023",
			nil,
			time.Date(2023, time.July, 7, 0, 0, 0, 0, time.UTC),
			false,
		},
		{
			"7/07/2023",
			nil,
			time.Date(2023, time.July, 7, 0, 0, 0, 0, time.UTC),
			false,
		},
		{
			"2023/07/23",
			nil,
			time.Date(2023, time.July, 23, 0, 0, 0, 0, time.UTC),
			false,
		},
		{
			"7.07.2023 16:30",
			nil,
			time.Date(2023, time.July, 7, 16, 30, 0, 0, time.UTC),
			false,
		},
		{
			"16:30",
			nil,
			time.Date(0, time.January, 1, 16, 30, 0, 0, time.UTC),
			false,
		},

		// GoLang style.
		{
			"16/07/2023 14:00:00",
			[]string{"02/01/2006 15:04:05"},
			time.Date(2023, time.July, 16, 14, 0, 0, 0, time.UTC),
			false,
		},

		// Python style.
		{
			"16/07/2023 14:00:00",
			[]string{"%d/%m/%Y %H:%M:%S"},
			time.Date(2023, time.July, 16, 14, 0, 0, 0, time.UTC),
			false,
		},

		// Incorrect format.
		{
			"16/07/2023 14:00:00",
			[]string{"%Y/%m/%d %H:%M:%S"},
			time.Time{},
			true,
		},

		// Gorutines.
		{
			"16/07/2023 14:00:00",
			[]string{
				"02/01/2006 15:04:05",
				"7.07.2023",
				"7/07/2023",
				"2023/07/23",
				"7.07.2023 16:30",
				"16/07/2023 14:00:00",
			},
			time.Date(2023, time.July, 16, 14, 0, 0, 0, time.UTC),
			false,
		},
	}

	for _, test := range tests {
		got, err := StringToDate(test.input, test.patterns...)
		if (err != nil) != test.wantErr {
			t.Errorf("StrToDate(%q, %v) error = %v, wantErr %v",
				test.input, test.patterns, err, test.wantErr)
			continue
		}
		if !got.Equal(test.want) {
			t.Errorf("StrToDate(%q, %v) = %v, want %v",
				test.input, test.patterns, got, test.want)
		}
	}
}

// TestDateToString tests DateToString function.
func TestDateToString(t *testing.T) {
	t1 := time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC)

	s, err := DateToString(t1, "2006-01-02")
	if err != nil {
		t.Fatal(err)
	}
	if s != "2023-07-17" {
		t.Errorf("Expected 2023-07-17, got %s", s)
	}

	s, err = DateToString(t1, "02-01-2006", "01-02-2006")
	if err != nil {
		t.Fatal(err)
	}
	if s != "17-07-2023" {
		t.Errorf("Expected 17-07-2023, got %s", s)
	}

	s, _ = DateToString(t1, "abc")
	if s != "" {
		t.Errorf("Expected an empty string, got %s", s)
	}
}

// TestDateToStrings tests DateToStrings function.
func TestDateToStrings(t *testing.T) {
	t.Run("Test with multiple formats", func(t *testing.T) {
		testTime := time.Date(2020, 7, 17, 0, 0, 0, 0, time.UTC)
		patterns := []string{"2006-01-02", "02 Jan 06", time.RFC3339}
		expected := []string{"2020-07-17", "17 Jul 20", "2020-07-17T00:00:00Z"}
		results, err := DateToStrings(testTime, patterns...)

		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected %v, got %v", expected, results)
		}
	})

	t.Run("Test with no formats", func(t *testing.T) {
		testTime := time.Date(2023, 6, 16, 0, 0, 0, 0, time.UTC)
		expected := []string{"2023-06-16 00:00:00"} // default DateTime format
		results, err := DateToStrings(testTime)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected %v, got %v", expected, results)
		}
	})

	// Assuming pythonToGolangFormat can convert %Y-%m-%d to 2006-01-02
	t.Run("Test with python strftime format", func(t *testing.T) {
		testTime := time.Date(2023, 6, 16, 0, 0, 0, 0, time.UTC)
		patterns := []string{"%Y-%m-%d"}
		expected := []string{"2023-06-16"}
		results, err := DateToStrings(testTime, patterns...)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if !reflect.DeepEqual(results, expected) {
			t.Errorf("Expected %v, got %v", expected, results)
		}
	})

	// Incorrect format.
	t.Run("Test with incorrect format", func(t *testing.T) {
		testTime := time.Date(2023, 6, 16, 0, 0, 0, 0, time.UTC)
		patterns := []string{"abc"}
		r, err := DateToStrings(testTime, patterns...)
		if err == nil {
			t.Errorf("An error is expected for an invalid format, but %v", r)
		}
	})
}

// TestChangeTimeZone tests ChangeTimeZone function.
func TestChangeTimeZone(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-06-17T08:15:45Z")
	t2, _ := ChangeTimeZone(t1, "America/New_York")

	if t2.Location().String() != "America/New_York" {
		t.Errorf("Expected 'America/New_York', got %s", t2.Location())
	}

	if t1.Hour() != t2.Hour() || t1.Minute() != t2.Minute() {
		t.Errorf("Expected hour and minute to remain the same, " +
			"but they changed")
	}

	_, err := ChangeTimeZone(t1, "Invalid/TimeZone")
	if err == nil {
		t.Errorf("Expected an error, but didn't get one")
	}
}

// TestSetTimeZone tests SetTimeZone function.
func TestSetTimeZone(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2023-06-17T08:15:45Z")
	t2, _ := SetTimeZone(t1, "America/New_York")

	if t2.Location().String() != "America/New_York" {
		t.Errorf("Expected 'America/New_York', got %s", t2.Location())
	}

	if t1.Hour() == t2.Hour() && t1.Minute() == t2.Minute() {
		t.Errorf("Expected hour and minute to change, " +
			"but they remained the same")
	}

	_, err := SetTimeZone(t1, "Invalid/TimeZone")
	if err == nil {
		t.Errorf("Expected an error, but didn't get one")
	}
}

// TestMoveTimeZone tests MoveTimeZone function.
func TestMoveTimeZone(t *testing.T) {
	t1 := time.Date(2023, 6, 17, 12, 0, 0, 0, time.UTC)

	t2 := MoveTimeZone(t1, 3)
	expected := time.Date(2023, 6, 17, 15, 0, 0, 0, time.UTC)
	if !t2.Equal(expected) {
		t.Errorf("Expected %s, got %s", expected, t2)
	}

	t3 := MoveTimeZone(t1, -2)
	expected = time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC)
	if !t3.Equal(expected) {
		t.Errorf("Expected %s, got %s", expected, t3)
	}
}
