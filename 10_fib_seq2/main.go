package main

import (
	"fmt"
	"iter"
)

func Fib(n int) iter.Seq2[int, int] {
	return func(yield func(n, v int) bool) {
		prev, cur := 0, 1
		for i := range n {
			if !yield(i+1, cur) {
				return
			}
			prev, cur = cur, prev+cur
		}
	}
}

func main() {
	for i, v := range Fib(5) {
		fmt.Println(i, v)
	}
}
