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

func Sqr5(yield func(v int)) {
	for i := 1; i <= 5; i++ {
		yield(i * i)
	}
}

func main() {
	nextFib := Pull(Fib10)
	nextSqr := Pull(Sqr5)

	vFib, okFib := nextFib()
	vSqr, okSqr := nextSqr()

	for i := 1; okFib || okSqr; i++ {
		fmt.Printf("%v: fib = %v, sqr = %v\n", i, vFib, vSqr)

		vFib, okFib = nextFib()
		vSqr, okSqr = nextSqr()
	}
}
