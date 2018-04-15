package main

import (
	"fmt"
)

func main() {
	i := 5
	p := &i
	var j int = 5
	var q *int = &j
	fmt.Printf("%d %d %d", *p, j, *q)
}
