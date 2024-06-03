package slicesort

import (
	"sort"

	"github.com/pkg/errors"
)

type Slice[V any] struct {
	data  []V
	iLess func(i, j int) bool
	vLess func(a, b V) bool
}

func NewSortByIndex[V any](a []V, iLess func(i, j int) bool) sort.Interface {
	return &Slice[V]{data: a, iLess: iLess}
}

func NewSortByValue[V any](a []V, vLess func(a, b V) bool) sort.Interface {
	return &Slice[V]{data: a, vLess: vLess}
}

func (T *Slice[V]) Len() int {
	return len(T.data)
}

func (T *Slice[V]) Less(i, j int) bool {
	switch {
	case T.iLess != nil:
		return T.iLess(i, j)
	case T.vLess != nil:
		return T.vLess(T.data[i], T.data[j])
	default:
		panic(errors.New("missing less function"))
	}
}

func (T *Slice[V]) Swap(i, j int) {
	T.data[i], T.data[j] = T.data[j], T.data[i]
}
