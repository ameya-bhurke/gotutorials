package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot square root a negative number: %f", e)
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	var y = float64(1)
	for math.Abs(x-y*2) > 0.00 {
		y -= ((y * y) - x) / (2 * y)
		y = math.Abs(y)
	}
	return y, nil
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
