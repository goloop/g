[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/do)](https://goreportcard.com/report/github.com/goloop/do) [![License](https://img.shields.io/badge/godoc-A+-brightgreen)](https://godoc.org/github.com/goloop/do) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/do/blob/master/LICENSE)

# Do

The **do** package is a powerful utility library for Go programming language, extending Go's functionality with a suite of helper functions that simplify software development. As of Go 1.20, we leverage generics to provide flexible and type-safe functions that can operate across many data types.

Some of the key features of **do** package are:

 - **Generics**: The package uses Go's new generics features extensively, allowing for flexible and type-safe operations across a variety of data types.

 - **Simplified Conditionals**: Includes an If function, providing a shorthand similar to the ternary operator found in languages like C/C++ and Python.
   ```
   Python |  max = a if a > b else b
   C/C++  |  int max = (a > b) ? a : b;
   Go     |  max := do.If(a > b, a, b)
   ```

 - **Zero-Value Checks**: Includes `IsEmpty`, `All`, `Any`, and `Value` functions that perform checks for zero-values in various contexts.

 - **Numeric Operations**: The package provides flexible and efficient functions like `Min`, `Max`, `Sum`, `Average`, `Median` that work on all numeric types.

 - **Functional Programming**: It also offers functional programming helpers like `Sort`, `Map`, `Filter`, `Reduce` that apply functions to slices and return new slices.

 - **Element Operations**: The `Contains` and `Index` functions to check if a slice contains a certain element and to get the index of an element.

 - **Pairing**: A `Zip` function is provided that pairs elements from two slices together into a slice of Pair structures.

Please note that usage of the do package requires **Go 1.20** or later, due to its use of generics.
