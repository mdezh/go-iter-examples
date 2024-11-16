package main

import "fmt"

func Fib10() {
	prev, cur := 0, 1
	for range 10 {
		fmt.Println(cur)
		prev, cur = cur, prev+cur
	}
}

func main() {
	Fib10()
}
