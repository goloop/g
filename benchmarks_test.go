package g

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Helper functions/
func generateIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1000000)
	}
	return slice
}

func generateFloat64Slice(size int) []float64 {
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Float64() * 1000000
	}
	return slice
}

func generateStringSlice(size int) []string {
	slice := make([]string, size)
	for i := 0; i < size; i++ {
		slice[i] = fmt.Sprintf("test-string-%d", rand.Intn(1000000))
	}
	return slice
}

// Sort function benchmarks.
func BenchmarkSort(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Sort/Ints/%d", size), func(b *testing.B) {
			data := generateIntSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				testData := make([]int, len(data))
				copy(testData, data)
				Sort(testData)
			}
		})

		b.Run(fmt.Sprintf("Sort/Float64s/%d", size), func(b *testing.B) {
			data := generateFloat64Slice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				testData := make([]float64, len(data))
				copy(testData, data)
				Sort(testData)
			}
		})

		b.Run(fmt.Sprintf("Sort/Strings/%d", size), func(b *testing.B) {
			data := generateStringSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				testData := make([]string, len(data))
				copy(testData, data)
				Sort(testData)
			}
		})
	}
}

// StringToDate function benchmarks.
func BenchmarkStringToDate(b *testing.B) {
	dates := []string{
		"2023-12-01",
		"01/12/2023",
		"2023.12.01",
		"01-12-2023 15:04:05",
		"2023/12/01 15:04",
		"Dec 1, 2023",
	}

	for _, date := range dates {
		b.Run(fmt.Sprintf("StringToDate/%s", date), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = StringToDate(date)
			}
		})
	}
}

// In function benchmarks.
func BenchmarkIn(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("In/Ints/%d/Found", size), func(b *testing.B) {
			data := generateIntSlice(size)
			searchValue := data[size/2] // Значення посередині
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = In(searchValue, data...)
			}
		})

		b.Run(fmt.Sprintf("In/Ints/%d/NotFound", size), func(b *testing.B) {
			data := generateIntSlice(size)
			searchValue := -1 // Гарантовано відсутнє значення
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = In(searchValue, data...)
			}
		})
	}
}

// Range function benchmarks.
func BenchmarkRange(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Range/Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Range(size)
			}
		})

		b.Run(fmt.Sprintf("Range/WithStep/Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Range(0, size, 2)
			}
		})
	}
}

// Set operation benchmarks.
func BenchmarkSetOperations(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Union/Size_%d", size), func(b *testing.B) {
			set1 := generateIntSlice(size)
			set2 := generateIntSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Union(set1, set2)
			}
		})

		b.Run(fmt.Sprintf("Intersection/Size_%d", size), func(b *testing.B) {
			set1 := generateIntSlice(size)
			set2 := generateIntSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Intersection(set1, set2)
			}
		})

		b.Run(fmt.Sprintf("Difference/Size_%d", size), func(b *testing.B) {
			set1 := generateIntSlice(size)
			set2 := generateIntSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Difference(set1, set2)
			}
		})
	}
}

// String operation benchmarks.
func BenchmarkStringOperations(b *testing.B) {
	testString := "Hello, World! This is a test string with some " +
		"numbers: 12345 and symbols: @#$%"

	b.Run("Weed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Weed(testString, "!@#$%^&*()")
		}
	})

	b.Run("Preserve", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Preserve(testString, Letters, Numbers)
		}
	})

	b.Run("Trim", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Trim(testString, Whitespaces)
		}
	})
}

// Math operation benchmarks.
func BenchmarkMathOperations(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Sum/Size_%d", size), func(b *testing.B) {
			data := generateIntSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Sum(data...)
			}
		})

		b.Run(fmt.Sprintf("Average/Size_%d", size), func(b *testing.B) {
			data := generateFloat64Slice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Average(data...)
			}
		})

		b.Run(fmt.Sprintf("Median/Size_%d", size), func(b *testing.B) {
			data := generateFloat64Slice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Median(data...)
			}
		})
	}
}

// Random operation benchmarks.
func BenchmarkRandomOperations(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		slice := generateIntSlice(size)

		b.Run(fmt.Sprintf("RandomList/Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RandomList(slice)
			}
		})

		b.Run(fmt.Sprintf("RandomListPlural/Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = RandomListPlural(10, slice)
			}
		})
	}
}

// Conversion operation benchmarks.
func BenchmarkTypeConversion(b *testing.B) {
	// StringToInt
	b.Run("StringToInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = StringToInt("12345")
		}
	})

	// StringToFloat
	b.Run("StringToFloat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = StringToFloat("123.45")
		}
	})

	// StringToBool
	b.Run("StringToBool", func(b *testing.B) {
		values := []string{"true", "false", "yes", "no", "on", "off"}
		for _, v := range values {
			b.Run(fmt.Sprintf("Value_%s", v), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					_, _ = StringToBool(v)
				}
			})
		}
	})
}

// Zip operation benchmarks.
func BenchmarkZip(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Zip/Size_%d", size), func(b *testing.B) {
			slice1 := generateIntSlice(size)
			slice2 := generateStringSlice(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = Zip(slice1, slice2)
			}
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
