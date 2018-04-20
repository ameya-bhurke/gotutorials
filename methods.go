package main

import (
	"fmt"
	"math"
)

// Vertex A vertex on a 2 dimensional cartesian coordinate plane.
type Vertex struct {
	X float64
	Y float64
}

// Abs returns the distance of the Vertex form the centre.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale scales the vertex by a factor f
func (v* Vertex) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v = Vertex{1, 2}
	fmt.Println(v.Abs())
	v.Scale(1.2)
	fmt.Println(v)
}
