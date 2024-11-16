package main

import "fmt"

func Fib10(yield func(v int)) {
	prev, cur := 0, 1
	for range 10 {
		yield(cur)
		prev, cur = cur, prev+cur
	}
}

func main() {
	i := 1
	Fib10(func(v int) {
		fmt.Printf("%v: %v\n", i, v)
		i++
	})
}
