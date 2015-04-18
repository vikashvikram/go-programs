package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	fmt.Println("My favorite number is", rand.Intn(100))
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Println("My favorite number is", s1)
	fmt.Println("My favorite number is", r1.Intn(100))
	fmt.Println("My favorite number is", r1.Intn(100))
	fmt.Println("My favorite number is", r1.Intn(100))
}

