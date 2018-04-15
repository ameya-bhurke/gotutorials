package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var y = float64(1)
	for math.Abs(x-y*2) > 0.00 {
		y -= ((y * y) - x) / (2 * y)
		y = math.Abs(y)
	}
	return y
}

func main() {
	fmt.Printf("The square root of 4 is %f", sqrt(4))
}
