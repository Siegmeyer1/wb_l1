package main

import (
	"fmt"
)

func ReverseString(str string) string {
	runes := []rune(str)
	l := len(runes)
	m := l / 2
	for i := 0; i < m; i++ { // iterate with i over 1-st half of a rune slice
		runes[i], runes[l-i-1] = runes[l-i-1], runes[i] // swap i-th from beginning and i-th from end runes
	}
	return string(runes)
}

func main() {
	str := "武士道"
	reversed := ReverseString(str)
	fmt.Println(reversed)
	str = "武士道Ъ"
	reversed = ReverseString(str)
	fmt.Println(reversed)
}
