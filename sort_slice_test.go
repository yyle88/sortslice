package slicesort

import (
	"testing"
)

func TestSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	SortByIndex(a, func(i, j int) bool { return a[i] < a[j] })
	t.Log(a)
}

func TestSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	SortByValue(a, func(a, b int) bool { return a < b })
	t.Log(a)
}
