package main

import (
	"fmt"
	"iter"
)

func Fib(n int) iter.Seq[int] {
	return func(yield func(v int) bool) {
		prev, cur := 0, 1
		for range n {
			if !yield(cur) {
				return
			}
			prev, cur = cur, prev+cur
		}
	}
}

func main() {
	for v := range Fib(5) {
		fmt.Println(v)
	}
}
