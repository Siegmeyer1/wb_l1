package main

import (
	"fmt"
	"math"
)

type Dot struct {
	x float64
	y float64
}

func NewDot(x, y float64) *Dot {
	return &Dot{x: x, y: y}
}

func Dist(A, B *Dot) float64 { // Just school math
	a := math.Abs(A.x - B.x)
	b := math.Abs(A.y - B.y)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	A := NewDot(-1.0, -1.0)
	B := NewDot(2.0, 2.0)
	fmt.Println(Dist(A, B))
}
