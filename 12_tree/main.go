package main

import "fmt"

type Tree[V any] struct {
	v    V
	l, r *Tree[V]
}

func (t *Tree[V]) All(yield func(v V) bool) {
	var process func(t *Tree[V]) bool
	process = func(t *Tree[V]) bool {
		if t == nil {
			return true
		}
		return process(t.l) && yield(t.v) && process(t.r)
	}
	process(t)
}

func main() {
	//        3
	//      /   \
	//    1      5
	//   / \    / \
	//  .   2  4   .
	t := &Tree[int]{
		l: &Tree[int]{
			v: 1,
			r: &Tree[int]{v: 2},
		},
		v: 3,
		r: &Tree[int]{
			l: &Tree[int]{v: 4},
			v: 5,
		},
	}

	for v := range t.All {
		fmt.Println(v)
		// if v == 4 {
		// 	break
		// }
	}
}
