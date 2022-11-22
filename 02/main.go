package main

import (
	"fmt"
	"sync"
)

func main() {
	var slice []int = []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for i, val := range slice {
		go func(_i, _val int) {
			fmt.Printf("%d (%d): %d\n", _i, _val, _val*_val)
			wg.Done()
		}(i, val)
	}
	wg.Wait()
}
