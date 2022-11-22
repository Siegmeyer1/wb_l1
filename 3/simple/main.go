package main

import (
	"fmt"
	"sync"
)

func main() {
	var slice []int = []int{2, 4, 6, 8, 10, 0}
	wg := sync.WaitGroup{} //creating wait group to make main gorutine wait for other gorutines
	wg.Add(len(slice))
	for i, val := range slice {
		go func(_i, _val int) { //concurrenly updating values in slice
			slice[_i] = _val * _val
			wg.Done()
		}(i, val)
	}
	wg.Wait()
	wg.Add(len(slice) / 2)
	for i := range slice {
		if i%2 == 0 {
			go func(_i int) { //summing pairs of values concurrently
				slice[_i] += slice[_i+1]
				wg.Done()
			}(i)
		}
	}
	wg.Wait()
	res := 0
	for i := range slice {
		if i%2 == 0 {
			res += slice[i]
		}
	}
	fmt.Println(res)
}
