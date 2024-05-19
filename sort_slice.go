package slicesort

import "sort"

func SortByIndex[V any](a []V, iLess func(i, j int) bool) {
	sort.Sort(NewSortByIndex(a, iLess))
}

func SortByValue[V any](a []V, vLess func(a, b V) bool) {
	sort.Sort(NewSortByValue(a, vLess))
}

func SortIStable[V any](a []V, iLess func(i, j int) bool) {
	sort.Stable(NewSortByIndex(a, iLess))
}

func SortVStable[V any](a []V, vLess func(a, b V) bool) {
	sort.Stable(NewSortByValue(a, vLess))
}
