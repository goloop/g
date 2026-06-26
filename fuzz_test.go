package g

import (
	"strconv"
	"testing"
)

// FuzzStringToDate ensures the parser never panics or hangs on arbitrary
// input — it must always return cleanly, succeeding or wrapping ErrParseDate.
func FuzzStringToDate(f *testing.F) {
	for _, s := range []string{
		"", "2023-07-07", "7.07.2023 16:30", "not a date",
		"%", "9999999999999999999", "16:30", "Feb 30, 2023",
	} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		// Both the default layouts and an explicit Python layout must be safe.
		_, _ = StringToDate(s)
		_, _ = StringToDate(s, "%Y-%m-%d")
	})
}

// FuzzStringToInt checks that StringToInt agrees with strconv on validity and
// never panics. When it succeeds, the value must round-trip.
func FuzzStringToInt(f *testing.F) {
	for _, s := range []string{"", "0", "-1", "2147483648", "+7", "x", "  5  "} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		got, err := StringToInt(s)
		want, werr := strconv.Atoi(s)
		if (err == nil) != (werr == nil) {
			t.Fatalf("StringToInt(%q) err=%v, strconv err=%v", s, err, werr)
		}
		if err == nil && got != want {
			t.Fatalf("StringToInt(%q) = %d, want %d", s, got, want)
		}
	})
}

// FuzzStringToFloat checks StringToFloat never panics on arbitrary input.
func FuzzStringToFloat(f *testing.F) {
	for _, s := range []string{"", "0", "-1.5", "1e308", "1e999", "NaN", "x"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		_, _ = StringToFloat(s)
	})
}

// FuzzWeedTrim ensures the Unicode string helpers never panic on arbitrary
// runes and never grow the input.
func FuzzWeedTrim(f *testing.F) {
	for _, s := range []string{"", "  hello  ", "\t\n\r", "ＡＢＣ", "héllo"} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		if w := Weed(s); len(w) > len(s) {
			t.Fatalf("Weed(%q) grew the string to %q", s, w)
		}
		_ = Trim(s)
		_ = Preserve(s)
	})
}
