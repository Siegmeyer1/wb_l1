package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

var justString string

// problem is that when taking substring like this (v[:100]) we take subarray of byte array that lays under string, and some symbols take more place than 1 byte
func someFunc() {
	var b bytes.Buffer
	for i := 0; i < 1<<10; i++ {
		b.WriteString("◉")
	}
	v := b.String()

	justString = v[:100]
}

// to fix it we first cast string to rune slice, that take subslice of last 100 elements
func fixedFunc() {
	var b bytes.Buffer
	for i := 0; i < 1<<10; i++ {
		b.WriteString("◉")
	}
	v := b.String()
	fmt.Println("huge string byte size: ", len(v))
	fmt.Println("huge string symbols count: ", utf8.RuneCountInString(v))
	justString = string([]rune(v)[:100]) // before getting substring we cast string to rune slice
}

func main() {
	someFunc()
	fmt.Println(justString)
	// Expected 100 "◉" symbols, got much less and undefined symbol at the end
	fixedFunc()
	fmt.Println(justString)
}
