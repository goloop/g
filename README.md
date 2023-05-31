[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/do)](https://goreportcard.com/report/github.com/goloop/do) [![License](https://img.shields.io/badge/godoc-A+-brightgreen)](https://godoc.org/github.com/goloop/do) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/do/blob/master/LICENSE)

# Do

The **do** package is a powerful utility library for Go programming language, extending Go's functionality with a suite of helper functions that simplify software development. As of Go 1.20, we leverage generics to provide flexible and type-safe functions that can operate across many data types.

Some of the key features of **do** package are:

 - **Generics**: The package uses Go's new generics features extensively, allowing for flexible and type-safe operations across a variety of data types.

 - **Simplified Conditionals**: Includes an If function, providing a shorthand similar to the ternary operator found in languages like C/C++ and Python.


 - **Zero-Value Checks**: Includes `IsEmpty`, `All`, `Any`, and `Value` functions that perform checks for zero-values in various contexts.

 - **Numeric Operations**: The package provides flexible and efficient functions like `Min`, `Max`, `Sum`, `Average`, `Median` that work on all numeric types.

 - **Functional Programming**: It also offers functional programming helpers like `Sort`, `Map`, `Filter`, `Reduce` that apply functions to slices and return new slices.

 - **Element Operations**: The `Contains` and `Index` functions to check if a slice contains a certain element and to get the index of an element.

 - **Pairing**: A `Zip` function is provided that pairs elements from two slices together into a slice of Pair structures.

Please note that usage of the do package requires **Go 1.20** or later, due to its use of generics.


## Functions

It is the detail list of functions provided in Go package do:

- **If**[T any](e bool, t, f T) T

  This function is a substitute for the ternary operator `(?:)` not available in Go. It accepts a boolean expression, and two values of any type. If the expression is true, it returns the first value, otherwise it returns the second value.

   ```
   Python |  max = a if a > b else b
   C/C++  |  int max = (a > b) ? a : b;
   Go     |  max := do.If(a > b, a, b)
   ```

- **All**[T any](v ...T) bool

  The All function checks if all values in the given slice are non-zero values for their types. If at least one value is a zero value, it immediately returns false. If the slice is empty, it also returns false.

- **Any**[T any](v ...T) bool

  The Any function checks if at least one value in the provided slice is a non-zero value for its type. As soon as it finds a value that is not a zero value, it returns true. If all values in the slice are zero values or the slice is empty, it returns false.

- **IsEmpty**[T any](v T) bool

  The IsEmpty function checks if a given value of any type is a "zero value" for that type. Zero values in Go are values that the variables of respective types hold upon their declaration, if they do not have any explicit initialization.

- **IsPointer**(v interface{}) bool

  The IsPointer function checks if a given value is a pointer type. It returns true if the value is a pointer, and false otherwise.

- **IsNumber**(v interface{}) bool

  The IsNumber function checks if a given value is of a numeric type, including integer and floating-point types. It returns true if the value is a numeric type, and false otherwise.

- **Rank**[T Verifiable](number T, array []T, ascending ...bool) int

  This function returns the rank of a given value when compared to a list of other values. It can rank both from largest to smallest and smallest to largest based on the optional boolean argument. It assigns the same rank for duplicate values. The function works with any types satisfying the Verifiable interface.

- **HLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  HLookup function looks up and retrieves data from a specific row in a table. It takes a search value, a slice of lookup values, a slice of result values, and an optional default value. If the search value is found in the lookup slice, it returns the corresponding value from the result slice, otherwise it returns the default value.

- **VLookup**[T comparable, U any](v T, lookup []T, result []U, def U) U

  Similar to HLookup, the VLookup function looks up and retrieves data from a specific column in a table. It works in the same way as HLookup but used for vertical lookups.

- **Abs**[T Numerable](v T) T

  The Abs function takes a numeric input value of any type that satisfies the Numerable interface and returns its absolute value. It works by using the negation operator for numeric types that support it or returning the original value for unsigned integer types.

- **Average**[T Numerable](v ...T) float64

  The Average function calculates the average of a variable number of values that are of a type satisfying the Numerable interface. It computes the sum of all the values and divides it by the number of values to get the average. If no values are provided, it returns 0. This function returns the average as a float64, regardless of the input type.

- **Median**[T Numerable](v ...T) float64

  The Median function calculates the median value of a variable number of values that are of a type satisfying the Numerable interface. It sorts the input values and then returns the middle value or the average of the two middle values if the number of values is even.

- **Max**[T Verifiable](v T, more ...T) T

  The Max function returns the largest value among all input values. It takes at least one parameter of a type that satisfies the Numerable interface, and additional values can be passed using variadic arguments. It works by iterating through all the passed values and returning the largest one.

- **MaxList**[T Verifiable](v []T, defaults ...T) T

  The MaxList function returns the largest value from a list of values of a type satisfying the Numerable interface. If the input list is empty, it uses provided default values or returns the minimal value for the Numerable type.

- **Min**[T Verifiable](v T, more ...T) T

  The Min function returns the smallest value among all input values. It takes at least one parameter of a type that satisfies the Numerable interface, and additional values can be passed using variadic arguments. It works by iterating through all the passed values and returning the smallest one.

- **MinList**[T Verifiable](v []T, defaults ...T) T

  The MinList function returns the smallest value from a list of values of a type satisfying the Numerable interface. If the input list is empty, it uses provided default values or returns the minimal value for the Numerable type.

- **Sum**[T Numerable](v ...T) T

  The Sum function calculates the sum of all input values of any type that satisfies the Numerable interface. This function does not handle overflow - if the sum of the input values exceeds the maximum value for type T, the result will wrap around.

- **IsEven**[T Numerable](v T, f ...bool) bool

  The IsEven function checks if a value is an even number. It accepts a value of any type T that satisfies the Numerable interface and considers only the integer part of the value if the 'f' argument is set to true.

- **IsOdd**[T Numerable](v T, f ...bool) bool

  The IsOdd function checks if a value is an odd number. It accepts a value of any type T that satisfies the Numerable interface and considers only the integer part of the value if the 'f' argument is set to true.

- **IsWhole**[T Numerable](v T) bool

  The IsWhole function checks if a value is a whole number. It accepts a value of any type T that satisfies the Numerable interface and returns true if the value does not have a fractional part.

- **Contains**[T comparable](vs []T, v T) bool

  The Contains function checks if a given element is present in the given slice. It iterates over the slice and returns true if it encounters the element, else returns false.

- **Filter**[T any](vs []T, f func(T) bool) []T

  Filter applies a predicate function to each element in a slice and returns a new slice containing only the elements for which the predicate function returned true.

- **Index**[T comparable](vs []T, v T) int

  The Index function returns the index of the first occurrence of the provided element in the given slice, or -1 if the element is not present.

- **Map**[T any, U any](vs []T, f func(T) U) []U

  Map applies a function to every element of the input slice and returns a new slice with the results.

- **Reduce**[T any, U any](vs []T, f func(U, T) U, init U) U

  Reduce combines all elements of a slice into a single value using a provided function.

- **Sort**[T Numerable](v []T, inverse ...bool)

  The Sort function sorts a slice in ascending or descending order, depending on the value of the 'inverse' parameter.

- **Value**[T any](v T, more ...T) T

  The Value function returns the first non-zero value from the parameters or zero value of type T if all parameters are zero values.

- **Zip**[T, U any](a []T, b []U) []Pair[T, U]

  The Zip function takes two slices and returns a slice of pairs, with each pair containing one element from each slice.

- **Distinct**[T comparable](v []T) []T

  Distinct returns a new slice with only unique elements from the input slice, removing any duplicates.

- **Shuffle**[T any](v []T)

  The Shuffle function rearranges the elements of the slice in a random order.

- **Product**[T Numerable](v ...T) T

  The Product function calculates the product of all numeric values in the input slice.

- **Merge**[T Verifiable](a []T, b []T) []T

  Merge takes two sorted slices as input and returns a single sorted slice that contains all the elements from both input slices.