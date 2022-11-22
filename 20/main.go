package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	words := strings.Fields(str)
	for i, l := 0, len(words); i < l/2; i++ { // iterate with i over 1-st half of a words
		words[i], words[l-i-1] = words[l-i-1], words[i] // swap i-th from beginning and i-th from end words
	}
	return strings.Join(words, " ")
}

func main() {
	str := "snow dog sun"
	reverse := ReverseWords(str)
	fmt.Println(reverse)

	str = "snow 道士武 хелло ворлд"
	reverse = ReverseWords(str)
	fmt.Println(reverse)
}
