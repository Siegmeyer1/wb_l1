package main

import "fmt"

func RemoveElement[T any](s []T, i int) []T {
	res := []T{}                  //create new slice
	res = append(res, s[:i]...)   // add everything before i-th elem to it
	res = append(res, s[i+1:]...) // and everything after i-th element too
	return res
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	res := RemoveElement(s, 2)
	fmt.Println(res, s) //not breaking initial slice
}
