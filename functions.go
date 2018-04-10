package main

import "fmt"

const greeting = "hello"

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
	fmt.Println()
	var c, python, java = true, false, "Hell no!"
	fmt.Println(c, python, java)
	fmt.Println(greeting)
	var sum = 0
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Printf("The sum is %d", sum)
	fmt.Println()
}
