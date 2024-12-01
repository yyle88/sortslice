[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/sortslice/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/sortslice/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/sortslice)](https://pkg.go.dev/github.com/yyle88/sortslice)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/sortslice/master.svg)](https://coveralls.io/github/yyle88/sortslice?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/sortslice.svg)](https://github.com/yyle88/sortslice/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/sortslice)](https://goreportcard.com/report/github.com/yyle88/sortslice)

# sortslice

`sortslice` is a Go package that provides a simple and flexible way to sort slices using custom comparison functions. It leverages Go's generics and the `sort.Interface` to avoid repeating the implementation of sorting logic for different types.

## CHINESE README

[中文说明](README.zh.md)

## Installation

To install the `sortslice` package, you can use the following command:

```bash
go get github.com/yyle88/sortslice
```

## Usage

The package offers several functions for sorting slices with different comparison strategies. Below are the key functions available:

### `SortByIndex`

Sorts the slice `a` using an index-based comparison function `iLess`.

```go
sortslice.SortByIndex(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place using the provided index-based comparison function.

### `SortByValue`

Sorts the slice `a` using a value-based comparison function `vLess`.

```go
sortslice.SortByValue(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place using the provided value-based comparison function.

### `SortIStable`

Sorts the slice `a` using an index-based comparison function `iLess` and preserves the order of equal elements (stable sort).

```go
sortslice.SortIStable(a []V, iLess func(i, j int) bool)
```

- `a`: The slice to be sorted.
- `iLess`: The function that compares the indices of two elements in the slice.
- Sorts the slice in place while maintaining the original order of equal elements (stable sort).

### `SortVStable`

Sorts the slice `a` using a value-based comparison function `vLess` and preserves the order of equal elements (stable sort).

```go
sortslice.SortVStable(a []V, vLess func(a, b V) bool)
```

- `a`: The slice to be sorted.
- `vLess`: The function that compares the values of two elements in the slice.
- Sorts the slice in place while maintaining the original order of equal elements (stable sort).

## Example

Here's a basic example of how to use `SortByIndex` and `SortByValue`:

```go
package main

import (
	"fmt"
	"github.com/yyle88/sortslice"
)

func main() {
	// Example 1: Sorting by index
	numbers := []int{5, 3, 8, 1, 4}
	sortslice.SortByIndex(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // Compare by values at indices
	})
	fmt.Println("Sorted by index:", numbers)

	// Example 2: Sorting by value
	strings := []string{"apple", "banana", "cherry", "date"}
	sortslice.SortByValue(strings, func(a, b string) bool {
		return a < b // Compare by string values
	})
	fmt.Println("Sorted by value:", strings)
}
```

## License

`sortslice` is open-source and released under the MIT License. See the LICENSE file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `sortslice`!** 🎉

Give me stars. Thank you!!!

## See stars
[![see stars](https://starchart.cc/yyle88/sortslice.svg?variant=adaptive)](https://starchart.cc/yyle88/sortslice)
