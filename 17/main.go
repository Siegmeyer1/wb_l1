package main

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints" //using only for generics
)

func BinSearch[T constraints.Ordered](arr []T, key T) (int, error) {
	l, r := -1, len(arr)
	for l < r-1 { // repeat until there`s nothing between boundaries
		m := (l + r) / 2  // set index of the middle element of segment
		if arr[m] > key { // if middle element greater than key, then key must be on the left, so change boundaries accordingly and repeat
			r = m
		} else {
			l = m
		}
	}
	if arr[l] == key { // if found key, return index, else return -1 and error
		return l, nil
	} else {
		return -1, errors.New("nonexist")
	}
}

func main() {
	arr := []int{-50, -44, -39, -39, -32, -25, -22, -13, -10, -5, -3, 6, 9, 12, 24, 31, 31, 37, 39, 44}
	fmt.Println(BinSearch(arr, -39))
	arr2 := []int{1, 2, 2, 2, 4}
	fmt.Println(BinSearch(arr2, 1))
	fmt.Println(BinSearch(arr2, 2))
	fmt.Println(BinSearch(arr2, 4))
	fmt.Println(BinSearch(arr2, 3))
	arr3 := []int{1, 1, 1, 2, 2}
	fmt.Println(BinSearch(arr3, 1))
	fmt.Println(BinSearch(arr3, 2))
}
