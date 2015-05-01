package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	next := 1
	current := 0
	return func() int{
		temp := current
		current = next
		next = temp + current
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("fib(%v) = %v\n", i, f())
	}
}

