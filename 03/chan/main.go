package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{2, 4, 6, 8, 10}
	c := make(chan int, len(slice)) //creating buffered chan
	res := 0
	for _, val := range slice {
		go func(_val int) {
			c <- _val * _val // every gorutine writes in channel
			return
		}(val)
	}
	for range slice {
		res += <-c // we read from channel len(slice) times. If there`s no message, main gorutine will wait until it appears
	}
	fmt.Println(res)
}
