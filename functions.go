package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y int) (a, b int) {
	a = y
	b = x
	return
}

func main() {
	fmt.Printf("add: %d", add(1, 3))
	fmt.Println()
	a, b := swap(2, 3)
	fmt.Printf("swap 2 and 3 returns %d and %d", a, b)
}
