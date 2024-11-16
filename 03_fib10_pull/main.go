package main

import "fmt"

func Pull[V any](pushIter func(yield func(v V))) func() (V, bool) {
	c := make(chan V)

	go func() {
		pushIter(func(v V) {
			c <- v
		})
		close(c)
	}()

	return func() (V, bool) {
		v, ok := <-c
		return v, ok
	}
}

func Fib10(yield func(v int)) {
	prev, cur := 0, 1
	for range 10 {
		yield(cur)
		prev, cur = cur, prev+cur
	}
}

func main() {
	next := Pull(Fib10)

	for {
		v, ok := next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
}
