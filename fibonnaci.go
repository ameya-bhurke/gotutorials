package main

import "fmt"

var pre int = 0
var post int = 1

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	return func() int {
		pre, post = post, pre+post
		return pre
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
