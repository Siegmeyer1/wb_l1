package main

import (
	"fmt"
	"unicode"
)

type set map[rune]struct{}

func UniqueCheck(s string) bool {
	se := set{}
	for _, c := range s { //iterate over symbols
		// if this symb (in lowercase) already in set -> not unique
		if _, ok := se[unicode.ToLower(c)]; ok {
			return false
			// else add in to set (in lowercase)
		} else {
			se[unicode.ToLower(c)] = struct{}{}
		}
	}
	return true
}

func main() {
	tests := []string{"abc", "Aac", "абв", "Ааб", "武士道", "武道道"}
	for i, test := range tests {
		fmt.Printf("test %d (%s): %t\n", i, test, UniqueCheck(test))
	}
}
