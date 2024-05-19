package slicesort

import (
	"sort"
	"testing"
)

func TestNewSortByIndex(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(NewSortByIndex(a, func(i, j int) bool { return a[i] < a[j] }))
	t.Log(a)
}

func TestNewSortByValue(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	sort.Sort(NewSortByValue(a, func(a, b int) bool { return a < b }))
	t.Log(a)
}
