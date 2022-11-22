package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/constraints" //using only for generics
)

func QuickSort[T constraints.Ordered](arr []T, args ...int) {
	l, r := 0, len(arr)-1
	if len(args) > 0 {
		l, r = args[0], args[1]
	}
	if l < r {
		p := swap(arr, l, r) // we swap elements so that elements in first half is less or equal then elements in second half
		QuickSort(arr, l, p) // and recoursively swap those halfs
		QuickSort(arr, p+1, r)
	}
}

func swap[T constraints.Ordered](arr []T, l, r int) int {
	var v T = arr[(l+r)/2] // we choose one random element to judge our swaps
	i, j := l, r
	for i <= j {
		for arr[i] < v { // iterate from both ends until mismatch found (val on the left is greater then val on the right)
			i++
		}
		for arr[j] > v {
			j--
		}
		if i >= j { // if we crossed closing i and j iterators then nothing more to swap
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return j
}

func main() {
	arr := []int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6}
	QuickSort(arr)
	fmt.Println(arr)
	str := []string{"honor", "bravery", "duty"}
	QuickSort(str)
	fmt.Println(str)
	r := []int{}
	for i := 0; i < 20; i++ {
		r = append(r, rand.Intn(100)-50)
	}
	fmt.Println(r)
	QuickSort(r)
	fmt.Println(r)
}
