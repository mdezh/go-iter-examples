package main

import (
	"fmt"
	"iter"
)

// seq utils

func First[V any](seq iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(v V) bool) {
		if n < 1 {
			return
		}

		for v := range seq {
			if !yield(v) {
				return
			}
			n--
			if n == 0 {
				return
			}
		}
	}
}

func While[V any](seq iter.Seq[V], p func(v V) bool) iter.Seq[V] {
	return func(yield func(v V) bool) {
		for v := range seq {
			if !p(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}

func Filter[V any](seq iter.Seq[V], p func(v V) bool) iter.Seq[V] {
	return func(yield func(v V) bool) {
		for v := range seq {
			if p(v) && !yield(v) {
				return
			}
		}
	}
}

func Map[V1 any, V2 any](seq iter.Seq[V1], m func(v V1) V2) iter.Seq[V2] {
	return func(yield func(v V2) bool) {
		for v := range seq {
			if !yield(m(v)) {
				return
			}
		}
	}
}

func Reduce[V any, A any](seq iter.Seq[V], acc A, r func(acc A, v V) A) A {
	for v := range seq {
		acc = r(acc, v)
	}
	return acc
}

// infinite sequence

func Fib(yield func(v int) bool) {
	prev, cur := 0, 1
	for yield(cur) {
		prev, cur = cur, prev+cur
	}
}

// predicates, mappers and reducers

func isEven(v int) bool {
	return v&1 == 0
}

func intoBrackets(v int) string {
	return fmt.Sprintf("(%v)", v)
}

func shorterThan(n int) func(v string) bool {
	return func(v string) bool {
		return len(v) < n
	}
}

func concat(acc, v string) string {
	return acc + v
}

// transform first 90 members of the Fib sequence to the string like "(N1)(N2)(N3)..."
// use only even members
// don't process N with length of "(N)" >= 7

func main() {
	head := First(Fib, 90)
	evens := Filter(head, isEven)
	inBrackets := Map(evens, intoBrackets)
	shorts := While(inBrackets, shorterThan(7))

	// Reduce with concat works non optimal way, for demo purposes only
	fmt.Println(Reduce(shorts, "", concat))
}
