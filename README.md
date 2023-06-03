[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/do)](https://goreportcard.com/report/github.com/goloop/do) [![License](https://img.shields.io/badge/godoc-A+-brightgreen)](https://godoc.org/github.com/goloop/do) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/do/blob/master/LICENSE)

# Do

The **do** package is a comprehensive utility library for the Go programming language, augmenting Go's functionality with a collection of efficient helper functions aimed at simplifying and streamlining software development tasks. It robustly employs Go's generics, introduced in Go 1.18+, to provide versatile and type-safe functions capable of operating across a plethora of data types.

Key features of the **do** package include:

- **Extensive use of Generics**: The package exploits Go's generic features to their fullest extent, allowing for flexible and type-safe operations across diverse data types.

- **Simplified Conditionals with the If Method**: The `If` function acts as a shorthand similar to the ternary operator found in languages like C/C++ and Python. It simplifies conditional statements, enhancing code readability and maintenance.

    Example:
    ```go
    resp := do.If(do.In(ip, blacklist...), Response{...}, Response{...})
    ```

- **Zero-Value Checks**: Includes `IsEmpty`, `All`, `Any`, and `Value` functions that perform checks for zero-values in various contexts, thus saving valuable coding time and making code easier to comprehend.

- **Numeric Operations**: The package furnishes flexible and efficient functions like `Min`, `Max`, `Sum`, `Average`, `Median` that work on all numeric types. These functions greatly simplify the manipulation of numerical data, providing one-line solutions for common numerical operations.

- **Functional Programming Aids**: The package offers functional programming helpers like `Sort`, `Map`, `Filter`, `Reduce` that apply functions to slices and return new slices. These functions help in crafting cleaner and more manageable code by reducing boilerplate and enhancing readability.

- **Element Operations**: The `Contains` and `Index` functions check if a slice contains a certain element and fetch the index of an element, thereby simplifying tasks related to data manipulation and retrieval.

- **Pairing and More Complex Functionality**: A `Zip` function is provided that pairs elements from two slices together into a slice of Pair structures. It also offers complex functionalities like ranking, looking up data in arrays and maps, and manipulating lists in diverse ways including sorting, shuffling, and reduction.

Please note that usage of the do package requires **Go 1.20** or later, due to its extensive use of generics. This package is an invaluable tool for developers seeking to write clean, concise, efficient, and maintainable Go code, harnessing the full power of Go generics.


## Functions

It is the detail list of functions provided in Go package **do**:


- **If**[T any](e bool, t, f T) T

  This function is a substitute for the ternary operator `(?:)` not available in Go. It accepts a boolean expression, and two values of any type. If the expression is true, it returns the first value, otherwise it returns the second value.

  ```
  Python | max = a if a > b else b
  C/C++ | int max = (a > b) ? a : b;
  Go | max := do.If(a > b, a, b)
  ```


- **Abs**[T Numerable](v T) T

  The Abs function takes a numeric input value of any type that satisfies the Numerable interface and returns its absolute value. It works by using the negation operator for numeric types that support it or returning the original value for unsigned integer types.

- **All**[T any](v ...T) bool

  The All function checks if all values in the given slice are non-zero values for their types. If at least one value is a zero value, it immediately returns false. If the slice is empty, it also returns false.

- **Any**[T any](v ...T) bool

  The Any function checks if at least one value in the provided slice is a non-zero value for its type. As soon as it finds a value that is not a zero value, it returns true. If all values in the slice are zero values or the slice is empty, it returns false.

- **Average**[T Numerable](v ...T) float64

  The Average function calculates the average of a variable number of values that are of a type satisfying the Numerable interface. It computes the sum of all the values and divides it by the number of values to get the average. If no values are provided, it returns 0. This function returns the average as a float64, regardless of the input type.

- **Contains**[T comparable](vs []T, v T) bool

  The Contains function checks if a given element is present in the given slice. It iterates over the slice and returns true if it encounters the element, else returns false.

- **Distinct**[T comparable](v []T) []T

  Distinct returns a new slice with only unique elements from the input slice, removing any duplicates.

- **Filter**[T any](vs []T, f func(T) bool) []T

  Filter applies a predicate function to each element in a slice and returns a new slice containing only the elements for which the predicate function returned true.

- **HLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  HLookup function looks up and retrieves data from a specific row in a table. It takes a search value, a slice of lookup values, a slice of result values, and an optional default value. If the search value is found in the lookup slice, it returns the corresponding value from the result slice, otherwise it returns the default value.

- **In**[T Verifiable](v T, list ...T) bool

  The In function checks if a given value 'v' of type 'T' is present in a list of 'T'. It employs the 'Verifiable' type constraint, permitting it to function with numeric types and strings. It utilizes goroutines for concurrent processing, significantly enhancing performance for larger data sets. The function returns true if the value is found in the list, otherwise it returns false.

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

- **Index**[T comparable](vs []T, v T) int

  The Index function returns the index of the first occurrence of the provided element in the given slice, or -1 if the element is not present.

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

- **Rank**[T Verifiable](number T, array []T, ascending ...bool) int

  This function returns the rank of a given value when compared to a list of other values. It can rank both from largest to smallest and smallest to largest based on the optional boolean argument. It assigns the same rank for duplicate values. The function works with any types satisfying the Verifiable interface.

- **Reduce**[T any, U any](vs []T, f func(U, T) U, init U) U

  Reduce combines all elements of a slice into a single value using a provided function.

- **Shuffle**[T any](v []T)

  The Shuffle function rearranges the elements of the slice in a random order.

- **Sort**[T Numerable](v []T, inverse ...bool)

  The Sort function sorts a slice in ascending or descending order, depending on the value of the 'inverse' parameter.

- **Sum**[T Numerable](v ...T) T

  The Sum function calculates the sum of all input values of any type that satisfies the Numerable interface. This function does not handle overflow - if the sum of the input values exceeds the maximum value for type T, the result will wrap around.

- **Value**[T any](v T, more ...T) T

  The Value function returns the first non-zero value from the parameters or zero value of type T if all parameters are zero values.

- **VLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  Similar to HLookup, the VLookup function looks up and retrieves data from a specific column in a table. It works in the same way as HLookup but used for vertical lookups.

- **Zip**[T, U any](a []T, b []U) []Pair[T, U]

  The Zip function takes two slices and returns a slice of pairs, where each pair contains one element from each of the input slices. The function stops when the shorter input slice is exhausted.
