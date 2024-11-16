package main

import "fmt"

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
	Fib10(func(v int) bool {
		fmt.Println(v)
		return v < 5
	})
}
