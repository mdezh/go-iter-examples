package main

import (
	"fmt"
	"iter"
)

func Fib10(yield func(v int) bool) {
	prev, cur := 0, 1
	for range 10 {
		if !yield(cur) {
			return
		}
		prev, cur = cur, prev+cur
	}
}

func main() {
	next, stop := iter.Pull(Fib10)
	defer stop()

	for v, ok := next(); ok; v, ok = next() {
		fmt.Println(v)
		if v >= 5 {
			// stop()
			break
		}
	}
}
