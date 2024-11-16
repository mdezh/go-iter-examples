package main

import "fmt"

func Fib10(yield func(v int) bool) {
	fmt.Println("start")
	prev, cur := 0, 1
	for range 10 {
		if !yield(cur) {
			break
		}
		prev, cur = cur, prev+cur
	}
	fmt.Println("finish")
}

func main() {
	fmt.Println("main start")
	for v := range Fib10 {
		fmt.Println(v)
		if v >= 5 {
			return
		}
	}
	fmt.Println("main finish")
}
