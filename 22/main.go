package main

import (
	"fmt"
	"math/big"
)

// Using wrappers over math.big functions
func Add(A, B int64) {
	a := big.NewInt(A)
	b := big.NewInt(B)
	a.Add(a, b)
	fmt.Println(a)
}
func Sub(A, B int64) {
	a := big.NewInt(A)
	b := big.NewInt(B)
	a.Sub(a, b)
	fmt.Println(a)
}

func Mul(A, B int64) {
	a := big.NewInt(A)
	b := big.NewInt(B)
	a.Mul(a, b)
	fmt.Println(a)
}

func Div(A, B int64) {
	a := big.NewInt(A)
	b := big.NewInt(B)
	a.Quo(a, b)
	fmt.Println(a)
}

func main() {
	var a int64 = 1<<63 - 1
	var b int64 = 1<<63 - 2
	Add(a, b)
	Sub(a, b)
	Mul(a, b)
	Div(a, b)

}
