package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Square root of 7 is %g", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(1, 2))
}
