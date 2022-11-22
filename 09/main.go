package main

import "fmt"

func StartChannel(arr []int) <-chan int {
	out := make(chan int) // This is the "first" channel, we write numbers in it
	go func() {
		for _, val := range arr {
			out <- val
		}
		close(out)
	}()
	return out
}

func Double(input <-chan int) <-chan int {
	out := make(chan int) // This is the "second" channel, we write doubled numbers in it
	go func() {
		for val := range input {
			out <- val * 2
		}
		close(out)
	}()
	return out
}

func main() {
	arr := []int{1, 2, 3, 4}
	ch1 := StartChannel(arr)
	ch2 := Double(ch1)
	for val := range ch2 {
		fmt.Println(val)
	}
}
