package g

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"
)

var (
	// The dataTimeFormats is predefined layouts for use
	// in DateToString, DateToStrings and StringToDate functions.
	// Each element shows by example the formatting of an element of
	// the reference time. Only these values are recognized.
	// Text in the layout string that is not recognized as part of the
	// reference time is echoed verbatim during Format and expected to appear
	// verbatim in the input to Parse.
	//
	//	Year: "2006" "06"
	//	Month: "Jan" "January" "01" "1"
	//	Day of the week: "Mon" "Monday"
	//	Day of the month: "2" "_2" "02"
	//	Day of the year: "__2" "002"
	//	Hour: "15" "3" "03" (PM or AM)
	//	Minute: "4" "04"
	//	Second: "5" "05"
	//	AM/PM mark: "PM"
	//
	// Numeric time zone offsets format as follows:
	//
	//	"-0700"     ±hhmm
	//	"-07:00"    ±hh:mm
	//	"-07"       ±hh
	//	"-070000"   ±hhmmss
	//	"-07:00:00" ±hh:mm:ss
	//
	// Replacing the sign in the format with a Z triggers
	// the ISO 8601 behavior of printing Z instead of an
	// offset for the UTC zone. Thus:
	//
	//	"Z0700"      Z or ±hhmm
	//	"Z07:00"     Z or ±hh:mm
	//	"Z07"        Z or ±hh
	//	"Z070000"    Z or ±hhmmss
	//	"Z07:00:00"  Z or ±hh:mm:ss
	dataTimeFormats = []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.Kitchen,
		time.Stamp,
		time.StampMilli,
		time.StampMicro,
		time.StampNano,
		time.DateTime,
		time.DateOnly,
		time.TimeOnly,

		"15:04", // short time: hours and minutes

		"2.1.06",
		"2.1.2006",
		"2.01.06",
		"2.01.2006",
		"02.01.2006",
		"02.1.06",
		"02.1.2006",

		"2.1.06 15:04",
		"2.1.2006 15:04",
		"2.01.06 15:04",
		"2.01.2006 15:04",
		"02.01.2006 15:04",
		"02.1.06 15:04",
		"02.1.2006 15:04",

		"2.1.06 15:04:05",
		"2.1.2006 15:04:05",
		"2.01.06 15:04:05",
		"2.01.2006 15:04:05",
		"02.01.2006 15:04:05",
		"02.1.06 15:04:05",
		"02.1.2006 15:04:05",

		"2/1/06",
		"2/1/2006",
		"2/01/06",
		"2/01/2006",
		"02/01/2006",
		"02/1/06",
		"02/1/2006",

		"2/1/06 15:04",
		"2/1/2006 15:04",
		"2/01/06 15:04",
		"2/01/2006 15:04",
		"02/01/2006 15:04",
		"02/1/06 15:04",
		"02/1/2006 15:04",

		"2/1/06 15:04:05",
		"2/1/2006 15:04:05",
		"2/01/06 15:04:05",
		"2/01/2006 15:04:05",
		"02/01/2006 15:04:05",
		"02/1/06 15:04:05",
		"02/1/2006 15:04:05",

		"2-1-06",
		"2-1-2006",
		"2-01-06",
		"2-01-2006",
		"02-01-2006",
		"02-1-06",
		"02-1-2006",

		"2-1-06 15:04",
		"2-1-2006 15:04",
		"2-01-06 15:04",
		"2-01-2006 15:04",
		"02-01-2006 15:04",
		"02-1-06 15:04",
		"02-1-2006 15:04",

		"2-1-06 15:04:05",
		"2-1-2006 15:04:05",
		"2-01-06 15:04:05",
		"2-01-2006 15:04:05",
		"02-01-2006 15:04:05",
		"02-1-06 15:04:05",
		"02-1-2006 15:04:05",

		"06/2/1",
		"2006/2/1",
		"06/2/01",
		"2006/2/01",
		"2006/02/01",
		"06/02/1",
		"2006/02/1",

		"06/2/1 15:04",
		"2006/2/1 15:04",
		"06/2/01 15:04",
		"2006/2/01 15:04",
		"2006/02/01 15:04",
		"06/02/1 15:04",
		"2006/02/1 15:04",

		"06/2/1 15:04:05",
		"2006/2/1 15:04:05",
		"06/2/01 15:04:05",
		"2006/2/01 15:04:05",
		"2006/02/01 15:04:05",
		"06/02/1 15:04:05",
		"2006/02/1 15:04:05",

		"06-2-1",
		"2006-2-1",
		"06-2-01",
		"2006-2-01",
		"2006-02-01",
		"06-02-1",
		"2006-02-1",

		"06-2-1 15:04",
		"2006-2-1 15:04",
		"06-2-01 15:04",
		"2006-2-01 15:04",
		"2006-02-01 15:04",
		"06-02-1 15:04",
		"2006-02-1 15:04",

		"06-2-1 15:04:05",
		"2006-2-1 15:04:05",
		"06-2-01 15:04:05",
		"2006-2-01 15:04:05",
		"2006-02-01 15:04:05",
		"06-02-1 15:04:05",
		"2006-02-1 15:04:05",

		"06/1/2",
		"2006/1/2",
		"06/01/2",
		"2006/01/2",
		"2006/01/02",
		"06/1/02",
		"2006/1/02",

		"06/1/2 15:04",
		"2006/1/2 15:04",
		"06/01/2 15:04",
		"2006/01/2 15:04",
		"2006/01/02 15:04",
		"06/1/02 15:04",
		"2006/1/02 15:04",

		"06/1/2 15:04:05",
		"2006/1/2 15:04:05",
		"06/01/2 15:04:05",
		"2006/01/2 15:04:05",
		"2006/01/02 15:04:05",
		"06/1/02 15:04:05",
		"2006/1/02 15:04:05",

		"06-1-2",
		"2006-1-2",
		"06-01-2",
		"2006-01-2",
		"06-1-02",
		"2006-1-02",

		"06-1-2 15:04",
		"2006-1-2 15:04",
		"06-01-2 15:04",
		"2006-01-2 15:04",
		"2006-01-02 15:04",
		"06-1-02 15:04",
		"2006-1-02 15:04",

		"06-1-2 15:04:05",
		"2006-1-2 15:04:05",
		"06-01-2 15:04:05",
		"2006-01-2 15:04:05",
		"06-1-02 15:04:05",
		"2006-1-02 15:04:05",
	}

	// The pythonToGolangFormats maps Python date format specifiers
	// to GoLang date format specifiers.
	//
	// Note that Python specifiers such as %U and %W have no
	// counterparts in GoLang.
	pythonToGolangFormats = map[string]string{
		"%d": "02",
		"%m": "01",
		"%Y": "2006",
		"%H": "15",
		"%M": "04",
		"%S": "05",
		"%a": "Mon",
		"%A": "Monday",
		"%w": "0",
		"%b": "Jan",
		"%B": "January",
		"%y": "06",
		"%I": "03",
		"%p": "PM",
		"%f": ".999999",
		"%z": "-0700",
		"%Z": "MST",
		"%j": "002",
		//"%U": "00", // ignored
		//"%W": "00", // ignored
		"%c": "Mon Jan 2 15:04:05 2006",
		"%x": "01/02/06",
		"%X": "15:04:05",
		"%%": "%",
	}
)

// The dateFoundValue is a thread-safe value that indicates
// whether a date was found in a string.
type dateFoundValue struct {
	m     sync.Mutex
	found bool
	value time.Time
}

// SetValue sets a new value for the found object. It locks the Mutex
// before changing the value and unlocks it after the change is complete.
func (f *dateFoundValue) SetValue(found bool, value time.Time) {
	f.m.Lock()
	defer f.m.Unlock()

	f.value = value
	f.found = found
}

// GetValue retrieves the current value of the Found. It locks the Mutex
// before reading the value and unlocks it after the read is complete.
func (f *dateFoundValue) GetValue() (time.Time, error) {
	f.m.Lock()
	defer f.m.Unlock()

	if f.found {
		return f.value, nil
	}

	return time.Time{}, errors.New("date and time could not be recognized")
}

// The pythonToGolangFormat converts a Python date format specifier to a
// GoLang date format specifier.
func pythonToGolangFormat(format string) string {
	builder := &strings.Builder{}

	i := 0
	for i < len(format) {
		replaced := false
		for k, v := range pythonToGolangFormats {
			if strings.HasPrefix(format[i:], k) {
				builder.WriteString(v)
				i += len(k)
				replaced = true
				break
			}
		}

		if !replaced {
			builder.WriteByte(format[i])
			i++
		}
	}

	return builder.String()
}

// StringToDate converts a string to a time.Time object using the provided
// formats. If no format is given, it uses default date-time formats.
//
// This function leverages goroutines to parse the string concurrently
// using different formats. When the first successful parsing occurs,
// the function stops all other goroutines and returns the result.
// If no parsing is successful, the function returns an error.
//
// Example usage:
//
//	// Automatic pattern detection.
//	t, err := StringToDate("2006-01-02")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	t, err = StringToDate("2006/01/02 15:04:05")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	// Using custom formats as Python.
//	t, err = StringToDate("Jan 02, 2006", "%b %d, %Y")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	// Using Go's date formatting
//	t, err = StringToDate("Jan 02, 2006", "Jan 02, 2006")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
func StringToDate(s string, patterns ...string) (time.Time, error) {
	var (
		wg      sync.WaitGroup
		formats []string
	)

	for _, pattern := range patterns {
		if strings.Contains(pattern, "%") {
			formats = append(formats, pythonToGolangFormat(pattern))
		} else {
			formats = append(formats, pattern)
		}
	}

	if len(formats) == 0 {
		formats = dataTimeFormats
	}

	// Will use context to stop the rest of the goroutines
	// if the value has already been found.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p := parallelTasks
	found := &dateFoundValue{found: false, value: time.Time{}}

	// If the length of the slice is less than or equal to
	// the minLoadPerGoroutine, then we do not need
	// to use goroutines.
	if l := len(formats); l/p < minLoadPerGoroutine {
		for _, format := range formats {
			t, err := time.Parse(format, s)
			if err == nil {
				return t, nil
			}
		}

		return found.GetValue()
	}

	chunkSize := len(formats) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(formats)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, layout := range formats[start:end] {
				// Check if the context has been cancelled.
				select {
				case <-ctx.Done():
					return
				default:
				}

				t, err := time.Parse(layout, s)
				if err == nil {
					found.SetValue(true, t)
					cancel() // stop all other goroutines
					return
				}
			}
		}(start, end)
	}

	wg.Wait()
	return found.GetValue()
}

// DateToString converts a Date to a string based on the provided format.
//
// If multiple formats are provided, only the first one will be processed.
//
// Example usage:
//
//	date := time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC)
//	s, err := DateToString(date, "2006-01-02")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(s)  // Output: "2023-07-17"
func DateToString(t time.Time, patterns ...string) (string, error) {
	var formats []string
	if len(patterns) != 0 {
		formats = append(formats, patterns[0]) // first only value
	}

	s, err := DateToStrings(t, formats...)
	if err != nil {
		return "", err
	}

	return s[0], nil
}

// DateToStrings converts a Time object into string(s) using the
// provided formats. The function uses time.DateTime format as a default
// option if no format is provided. In case the pattern includes
// a percentage ("%") character, it converts Python strftime format
// to Go's time format before applying it.
//
// It returns an array of the results and an error if there was
// a problem with the formatting.
//
// Example usage:
//
//	t := time.Now()
//	patterns := []string{"%Y-%m-%d", "02 Jan 06", time.RFC3339}
//	results, err := DateToStrings(t, patterns...)
//
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	for i, str := range results {
//		fmt.Printf("Date in format %s: %s\n", patterns[i], str)
//	}
func DateToStrings(t time.Time, patterns ...string) ([]string, error) {
	var (
		formats []string
		results []string
	)

	for _, pattern := range patterns {
		if strings.Contains(pattern, "%") {
			formats = append(formats, pythonToGolangFormat(pattern))
		} else {
			formats = append(formats, pattern)
		}
	}

	if len(formats) == 0 {
		// Use the format from the system locale.
		formats = append(formats, time.DateTime)
	}

	for _, format := range formats {
		// Format does not return an error if the formatting data is written
		// incorrectly. It just leaves the unrecognized fragment as is.
		s := t.Format(format)

		// The only way to check the correct format - reconstruction of time.
		t2, err := time.Parse(format, s)
		if err != nil || t != t2 {
			return []string{}, errors.New("invalid format")
		}

		results = append(results, s)
	}

	return results, nil
}

// ChangeTimeZone returns a time where the hour and minute are the same
// as the input time, but the time zone is changed. This can be used to
// convert a local time to a different time zone while keeping the "clock
// time" the same.
//
// Example usage:
//
//	t, _ := time.Parse(time.RFC3339, "2023-06-17T08:15:45Z")
//	newTime, err := ChangeTimeZone(t, "America/New_York")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(newTime)
//	// Output: 2023-06-17 08:15:45 -0400 EDT
func ChangeTimeZone(t time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		loc,
	), nil
}

// SetTimeZone changes the time zone, and changes the local time
// according to the new time zone. This can be used to convert the time from
// one time zone to another. The time value will adjust to maintain the same
// moment in time, but the hour, minute, and second may change.
//
// Example usage:
//
//	t, _ := time.Parse(time.RFC3339, "2023-06-17T08:15:45Z")
//	newTime, err := SetTimeZone(t, "America/New_York")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(newTime)
//	// Output: 2023-06-17 04:15:45 -0400 EDT
func SetTimeZone(t time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	return t.In(loc), nil
}

// MoveTimeZone adds or subtracts multiple time zones from a given time.
// The function uses the time.Duration type to add/subtract hours from
// the current time. Positive 'tz' values move the time forward, and
// negative values move it backward.
//
// Example usage:
//
//	t := time.Now()
//	newTime := MoveTimeZone(t, -3)
//	fmt.Println(newTime)
//	// Output: <current time minus 3 hours>
func MoveTimeZone(t time.Time, tz int) time.Time {
	return t.Add(time.Duration(tz) * time.Hour)
}
