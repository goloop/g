[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/g)](https://goreportcard.com/report/github.com/goloop/g) [![License](https://img.shields.io/badge/godoc-A+-brightgreen)](https://godoc.org/github.com/goloop/g) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/g/blob/master/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine&color=FFD500&labelColor=007ACC&style=flat)](https://u24.gov.ua/)

# G


The **g** package is a comprehensive utility library for the Go programming language, augmenting Go's functionality with a collection of efficient helper functions aimed at simplifying and streamlining software development tasks. It robustly employs Go's generics, introduced in Go 1.18+, to provide versatile and type-safe functions capable of operating across a plethora of data types.


## Features

Key features of the **g** package include:

- **Extensive use of Generics**: The package exploits Go's generic features to their fullest extent, allowing for flexible and type-safe operations across diverse data types.

- **Simplified Conditionals with the If Method**: The `If` function acts as a shorthand similar to the ternary operator found in languages like C/C++ and Python. It simplifies conditional statements, enhancing code readability and maintenance.

    Example:
    ```go
    fn := g.If(g.In(ip, blacklist...), prohibitedFn, successFn)
    ```

- **Zero-Value Checks**: Includes `IsEmpty`, `All`, `Any`, and `Value` functions that perform checks for zero-values in various contexts, thus saving valuable coding time and making code easier to comprehend.

- **Numeric Operations**: The package furnishes flexible and efficient functions like `Min`, `Max`, `Sum`, `Average`, `Median` that work on all numeric types. These functions greatly simplify the manipulation of numerical data, providing one-line solutions for common numerical operations.

- **Functional Programming Aids**: The package offers functional programming helpers like `Sort`, `Map`, `Filter`, `Reduce` that apply functions to slices and return new slices. These functions help in crafting cleaner and more manageable code by reducing boilerplate and enhancing readability.

- **Element Operations**: The `Contains` and `Index` functions check if a slice contains a certain element and fetch the index of an element, thereby simplifying tasks related to data manipulation and retrieval.

- **Pairing and More Complex Functionality**: A `Zip` function is provided that pairs elements from two slices together into a slice of Pair structures. It also offers complex functionalities like ranking, looking up data in arrays and maps, and manipulating lists in diverse ways including sorting, shuffling, and reduction.

Please note that usage of the do package requires **Go 1.20** or later, due to its extensive use of generics. This package is an invaluable tool for developers seeking to write clean, concise, efficient, and maintainable Go code, harnessing the full power of Go generics.


## Functions

It is the detail list of functions provided in Go package **g**:

- **Abs**[T Numerable](v T) T

  The Abs function takes a numeric input value of any type that satisfies the Numerable interface and returns its absolute value. It works by using the negation operator for numeric types that support it or returning the original value for unsigned integer types.

- **All**[T any](v ...T) bool

  The All function checks if all values in the given slice are non-zero values for their types. If at least one value is a zero value, it immediately returns false. If the slice is empty, it also returns false.

- **Any**[T any](v ...T) bool

  The Any function checks if at least one value in the provided slice is a non-zero value for its type. As soon as it finds a value that is not a zero value, it returns true. If all values in the slice are zero values or the slice is empty, it returns false.

- **Average**[T Numerable](v ...T) float64

  The Average function calculates the average of a variable number of values that are of a type satisfying the Numerable interface. It computes the sum of all the values and divides it by the number of values to get the average. If no values are provided, it returns 0. This function returns the average as a float64, regardless of the input type.

- **CartesianProduct**[T any](a []T, b []T) [][2]T

  The CartesianProduct function returns all possible pairs from two slices. It generates a slice of pairs, where each pair consists of an element from the first input slice and an element from the second input slice. The length of the returned slice is equal to the product of the lengths of the input slices. This function is generic and can work with any type T. Note that this function does not preserve the order of elements.

- **Complement**[T comparable](a []T, b []T) []T

  The Complement function takes a universal set (b) and a subset of it (a) and returns a new slice containing items present in the universal set but not in the subset. It returns the complement of the subset. The function is generic and can work with any type T that is comparable. Note that this function does not preserve the order of elements.

- **Contains**[T comparable](vs []T, v T) bool

  The Contains function checks if a given element is present in the given slice. It iterates over the slice and returns true if it encounters the element, else returns false.

- **DateToStr**(t time.Time, patterns ...string) (string, error)

  The DateToStr converts a Date to a string based on the provided format. Supports both Golang and Python style formatting. If multiple formats are provided, only the first one will be processed.

- **DateToStrPlural**(t time.Time, patterns ...string) ([]string, error)

  The DateToStrPlural converts a Date to a string based on the provided format. Supports both Golang and Python style formatting.

- **Difference**[T comparable](a []T, b []T) []T

  The Difference function takes two slices and returns a new slice that contains the items present in the first slice but not in the second slice. It returns the difference of the slices. The function is generic and can work with any type T that is comparable. Note that this function does not preserve the order of elements.

- **Diff**[T comparable](a []T, b []T) []T

  Diff is an alias for the **Difference** function.

- **Distinct**[T comparable](v []T) []T

  Distinct returns a new slice with only unique elements from the input slice, removing any duplicates.

- **Filter**[T any](vs []T, f func(T) bool) []T

  Filter applies a predicate function to each element in a slice and returns a new slice containing only the elements for which the predicate function returned true.

- **HLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  HLookup function looks up and retrieves data from a specific row in a table. It takes a search value, a slice of lookup values, a slice of result values, and an optional default value. If the search value is found in the lookup slice, it returns the corresponding value from the result slice, otherwise it returns the default value.

- **In**[T Verifiable](v T, list ...T) bool

  The In function checks if a given value 'v' of type 'T' is present in a list of 'T'. It employs the 'Verifiable' type constraint, permitting it to function with numeric types and strings. It utilizes goroutines for concurrent processing, significantly enhancing performance for larger data sets. The function returns true if the value is found in the list, otherwise it returns false.

- **Index**[T comparable](vs []T, v T) int

  The Index function returns the index of the first occurrence of the provided element in the given slice, or -1 if the element is not present.

- **Intersection**[T comparable](a []T, b []T) []T

  The Intersection function takes two slices and returns a new slice that contains the common items present in both slices. It returns the intersection of the slices. The function is generic and can work with any type T that is comparable. Note that this function does not preserve the order of elements.

- **If**[T any](e bool, t, f T) T

  This function is a substitute for the ternary operator `(?:)` not available in Go. It accepts a boolean expression, and two values of any type. If the expression is true, it returns the first value, otherwise it returns the second value.

  ```
  Python | max = a if a > b else b
  C/C++ | int max = (a > b) ? a : b;
  Go | max := g.If(a > b, a, b)
  ```

- **IsEmpty**[T any](v T) bool

  The IsEmpty function checks if a given value of any type is a "zero value" for that type. Zero values in Go are values that the variables of respective types hold upon their declaration, if they do not have any explicit initialization.

- **IsEven**[T Numerable](v T, f ...bool) bool

  The IsEven function checks if a value is an even number. It accepts a value of any type T that satisfies the Numerable interface and considers only the integer part of the value if the 'f' argument is set to true.

- **IsNumber**(v interface{}) bool

  The IsNumber function checks if a given value is of a numeric type, including integer and floating-point types. It returns true if the value is a numeric type, and false otherwise.

- **IsOdd**[T Numerable](v T, f ...bool) bool

  The IsOdd function checks if a value is an odd number. It accepts a value of any type T that satisfies the Numerable interface and considers only the integer part of the value if the 'f' argument is set to true.

- **IsPointer**(v interface{}) bool

  The IsPointer function checks if a given value is a pointer type. It returns true if the value is a pointer, and false otherwise.

- **IsWhole**[T Numerable](v T) bool

  The IsWhole function checks if a value is a whole number. It accepts a value of any type T that satisfies the Numerable interface and returns true if the value does not have a fractional part.

- **Map**[T any, U any](vs []T, f func(T) U) []U

  Map applies a function to every element of the input slice and returns a new slice with the results.

- **Max**[T Verifiable](v T, more ...T) T

  The Max function returns the largest value among all input values. It takes at least one parameter of a type that satisfies the Numerable interface, and additional values can be passed using variadic arguments. It works by iterating through all the passed values and returning the largest one.

- **MaxList**[T Verifiable](v []T, defaults ...T) T

  The MaxList function returns the largest value from a list of values of a type satisfying the Numerable interface. If the input list is empty, it uses provided default values or returns the minimal value for the Numerable type.

- **Median**[T Numerable](v ...T) float64

  The Median function calculates the median value of a variable number of values that are of a type satisfying the Numerable interface. It sorts the input values and then returns the middle value or the average of the two middle values if the number of values is even.

- **Merge**[T Verifiable](a []T, b []T) []T

  Merge takes two sorted slices as input and returns a single sorted slice that contains all the elements from both input slices.

- **Min**[T Verifiable](v T, more ...T) T

  The Min function returns the smallest value among all input values. It takes at least one parameter of a type that satisfies the Numerable interface, and additional values can be passed using variadic arguments. It works by iterating through all the passed values and returning the smallest one.

- **MinList**[T Verifiable](v []T, defaults ...T) T

  The MinList function returns the smallest value from a list of values of a type satisfying the Numerable interface. If the input list is empty, it uses provided default values or returns the minimal value for the Numerable type.

- **Preserve**(s string, patterns ...string) string

  The Preserve function keeps only characters specified by the patterns in the string.

- **Product**[T Numerable](v ...T) T

  The Product function calculates the product of all numeric values in the input slice.

- **Random**[T Numerable](v ...T) T

  The Random function generates a random value of type T based on the provided arguments. It supports different scenarios:

    - When called without any arguments, it returns the zero value of type T.
    - When called with one argument, it returns a random value from 0 to n-1.
    - When called with two arguments, it returns a random value from a to b-1.
    - When called with more than two arguments, it returns a randomly selected value from the provided arguments.

  The function uses the time in nanoseconds as a seed for the random number generator.

- **RandomList**[T any](v []T) T

  The RandomList function returns a random element from the given list. If the list is empty, it returns the zero value of type T.

- **RandomListPlural**[T any](n int, v []T) []T

  The RandomListPlural function returns a slice of n random elements from the given list v. If n is less than or equal to zero or if the list is empty, it returns an empty slice.

- **RandomMap**[K comparable, T any](m map[K]T) T

  The RandomMap function returns a random value from the given map. If the map is empty, it returns the zero value of type T.

- **RandomMapPlural**[K comparable, T any](n int, m map[K]T) []T

  The RandomMapPlural function returns a slice of n random values from the given map m. If n is less than or equal to zero or if the map is empty, it returns an empty slice.

- **Range**(a int, opt ...int) []int

  Range generates a slice of integers based on the provided parameters.

- **Rangef**[T any](fn func(int) T, a int, opt ...int) []T

  Rangef generates a slice of values based on the provided parameters and a given function as func(int) T.

- **Rank**[T Verifiable](number T, array []T, ascending ...bool) int

  This function returns the rank of a given value when compared to a list of other values. It can rank both from largest to smallest and smallest to largest based on the optional boolean argument. It assigns the same rank for duplicate values. The function works with any types satisfying the Verifiable interface.

- **Reduce**[T any, U any](vs []T, f func(U, T) U, init U) U

  Reduce combines all elements of a slice into a single value using a provided function.

- **Reverse**[T any](v []T) []T

  Reverse changes slice with elements in reverse order.

- **Shuffle**[T any](v []T)

  The Shuffle function rearranges the elements of the slice in a random order.

- **Sort**[T Numerable](v []T, inverse ...bool)

  The Sort function sorts a slice in ascending or descending order, depending on the value of the 'inverse' parameter.

- **StrToDate**(s string, patterns ...string) (time.Time, error)

  StrToDate converts a string to a time.Time object using the provided formats. If no format is given, it uses default date-time formats. Supports both Golang and Python style formatting.

- **Sum**[T Numerable](v ...T) T

  The Sum function calculates the sum of all input values of any type that satisfies the Numerable interface. This function does not handle overflow - if the sum of the input values exceeds the maximum value for type T, the result will wrap around.

- **SymmetricDifference**[T comparable](a []T, b []T) []T

  The SymmetricDifference function takes two slices and returns a new slice that contains the items present in one of the slices but not in both. It returns the symmetric difference of the slices. The function is generic and can work with any type T that is comparable. Note that this function does not preserve the order of elements.

- **Sdiff**[T comparable](a []T, b []T) []T

  Sdiff is an alias for the SymmetricDifference function.

- **Trim**(s string, patterns ...string) string

  Trim function is a utility that removes all leading and trailing occurrences of specified characters from the string. If no characters are provided, it trims leading and trailing whitespace. This function is especially useful to tidy up user input or to normalize strings for consistent processing. The function utilizes the TrimFunc from the standard strings package, making it efficient for processing large strings.

- **Union**[T comparable](a []T, b []T) []T

  The Union function takes two slices and returns a new slice that contains all unique items from both slices. It removes duplicates and returns the union of the slices. The function is generic and can work with any type T that is comparable. Note that this function does not preserve the order of elements.

- **Value**[T any](v T, more ...T) T

  The Value function returns the first non-zero value from the parameters or zero value of type T if all parameters are zero values.

- **VLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  Similar to HLookup, the VLookup function looks up and retrieves data from a specific column in a table. It works in the same way as HLookup but used for vertical lookups.

- **Weed**(s string, patterns ...string) string

  Weed function is a utility that helps 'clean up' your strings by removing any unwanted characters, termed as 'weeds'. If no patterns are specified, it removes the most common breakers characters (e.g., newline, tab). However, users can specify their own patterns according to their needs. An efficient mapping approach is employed to perform the cleaning process, making this function particularly effective for large strings.

- **Zip**[T, U any](a []T, b []U) []Pair[T, U]

  The Zip function takes two slices and returns a slice of pairs, where each pair contains one element from each of the input slices. The function stops when the shorter input slice is exhausted.


## Why exactly

When developing software in C/C++, I like to use the ternal operator `max = a > b ? a : b`, in Python as `max = a if a > b ese b`. I find this a convenient solution for small operations, and I'm very unhappy about the lack of this possibility in Go, an the classic solution is very bulky:

```go
max := a
if a < b {
     max = b
}
```

That's why I included an `If` method in this package:

```go
max := g.If(a>b, a, b)
```

A fairly common task is to check for the absence of something in a certain slice, this is well implemented in Python:

```python
if a in some_slice:
    # do something...
```

On Go, you need to implement the `in` function first, but a simple implementation that iterates through the slice elements in a loop is not efficient, so I added an `In` function checks the slice in multiple goroutines, so it works much faster for large lists:


```go
if g.In(a, someSlice...) {
    // do something...
}
```

Functions which are missing in Go but are, for example, in Python like `Any`, `And` etc.

And such functions as `Min`, `Max`, `Sum`, `Union`, `Distinct`, `Map`, `Filter` are generally found in almost all projects.

So I made a list of the most popular features (in my opinion) and grouped them together for quick access.

What about functions like SymmetricDifference, CartesianProduct or IsWhole which are not so popular? If the Union, Intersection function has already been added, it would be wrong not to add SymmetricDifference, CartesianProduct. Therefore, some functions have been added because they fit this package in terms of meaning.