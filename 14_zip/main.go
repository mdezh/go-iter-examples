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

func Zip[V any](sequences ...iter.Seq[V]) iter.Seq[[]V] {
	return func(yield func(v []V) bool) {
		n := len(sequences)
		if n < 1 {
			return
		}

		nexts := make([]func() (V, bool), n)

		for i, seq := range sequences {
			next, stop := iter.Pull(seq)
			defer stop()

			nexts[i] = next
		}

		for {
			// make a new tuple every iteration - for thread safety in the yield()
			tuple := make([]V, n)

			for i, next := range nexts {
				v, ok := next()
				if !ok {
					return
				}
				tuple[i] = v
			}

			if !yield(tuple) {
				return
			}
		}
	}
}

func Numerate[V any](seq iter.Seq[V], firstIndex int) iter.Seq2[int, V] {
	return func(yield func(i int, v V) bool) {
		i := firstIndex
		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// infinite sequences

func Fib(yield func(v int) bool) {
	prev, cur := 0, 1
	for yield(cur) {
		prev, cur = cur, prev+cur
	}
}

func Sqr(yield func(v int) bool) {
	for i := 1; yield(i * i); i++ {
	}
}

// print the first 5 pairs from the Fib and Sqr sequences
// numbered from 1 to 5

func main() {
	for i, tuple := range Numerate(First(Zip(Fib, Sqr), 5), 1) {
		fmt.Println(i, tuple)
	}
}
