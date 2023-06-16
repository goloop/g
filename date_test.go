package g

import (
	"reflect"
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

// TestStrToDate tests StrToDate function.
func TestStrToDate(t *testing.T) {
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

// TestDateToStrPlural tests DateToStrPlural function.
func TestDateToStrPlural(t *testing.T) {
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
}
