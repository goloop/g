package g

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"
)

var (
	dataTimeFormats = []string{
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
		"2006-01-02",
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
		"2006-01-02 15:04:05",
		"06-1-02 15:04:05",
		"2006-1-02 15:04:05",
	}

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

type foundDate struct {
	m     sync.Mutex
	value bool
	t     time.Time
}

// SetValue sets a new value for the Found. It locks the Mutex before
// changing the value and unlocks it after the change is complete.
func (f *foundDate) SetValue(value bool, t time.Time) {
	f.m.Lock()
	defer f.m.Unlock()
	f.t = t
	f.value = value
}

// GetValue retrieves the current value of the Found. It locks the Mutex
// before reading the value and unlocks it after the read is complete.
func (f *foundDate) GetValue() (time.Time, error) {
	f.m.Lock()
	defer f.m.Unlock()
	if f.value {
		return f.t, nil
	}

	return time.Time{}, errors.New("unable format")
}

func pythonToGolangFormat(format string) string {
	for k, v := range pythonToGolangFormats {
		format = strings.ReplaceAll(format, k, v)
	}
	return format
}

// StrToDate converts a string to a time.Time object using the provided
// formats. If no format is given, it uses default date-time formats.
// This function leverages goroutines to parse the string concurrently
// using different formats. When the first successful parsing occurs,
// the function stops all other goroutines and returns the result.
// If no parsing is successful, the function returns an error.
//
// Example usage:
//
//	t, err := StrToDate("2006-01-02")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	t, err = StrToDate("2006/01/02 15:04:05")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	// Using custom formats as Python.
//	t, err = StrToDate("Jan 02, 2006", "%b %d, %Y")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
//
//	// Using Go's date formatting
//	t, err = StrToDate("Jan 02, 2006", "Jan 02, 2006")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	fmt.Println(t)
func StrToDate(s string, patterns ...string) (time.Time, error) {
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

	v := formats
	p := parallelTasks
	found := &foundDate{value: false, t: time.Time{}}

	chunkSize := len(v) / p
	for i := 0; i < p; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == p-1 {
			end = len(v)
		}

		go func(start, end int) {
			defer wg.Done()

			for _, layout := range v[start:end] {
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

// DateToStr converts a Date to a string based on the provided format.
//
// If multiple formats are provided, only the first one will be processed.
//
// Example usage:
//
//	date := time.Date(2023, 7, 17, 0, 0, 0, 0, time.UTC)
//	s, err := DateToStr(date, "2006-01-02")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	fmt.Println(s)  // Output: "2023-07-17"
func DateToStr(t time.Time, patterns ...string) (string, error) {
	var formats []string
	if len(patterns) != 0 {
		formats = append(formats, patterns[0]) // first only value
	}

	s, err := DateToStrPlural(t, formats...)
	if err != nil {
		return "", err
	}

	return s[0], nil
}

// DateToStrPlural converts a Time object (t) into string(s) using the
// provided formats. The function uses RFC3339 format as a default
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
//	results, err := DateToStrPlural(t, patterns...)
//
//	if err != nil {
//		log.Fatal(err)
//		return
//	}
//
//	for i, str := range results {
//		fmt.Printf("Date in format %s: %s\n", patterns[i], str)
//	}
func DateToStrPlural(t time.Time, patterns ...string) ([]string, error) {
	var formats []string

	for _, pattern := range patterns {
		if strings.Contains(pattern, "%") {
			formats = append(formats, pythonToGolangFormat(pattern))
		} else {
			formats = append(formats, pattern)
		}
	}

	if len(formats) == 0 {
		// Use the format from the system locale.
		formats = append(formats, time.RFC3339)
	}

	var results []string
	for _, format := range formats {
		s := t.Format(format)
		_, err := time.Parse(format, s)
		if err != nil {
			return []string{}, errors.New("unable format")
		}
		results = append(results, s)
	}

	return results, nil
}
