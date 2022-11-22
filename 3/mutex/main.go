package main

import (
	"fmt"
	"sync"
)

func main() {
	var slice []int = []int{2, 4, 6, 8, 10}
	mu := sync.Mutex{}
	res := 0
	wg := sync.WaitGroup{} //using wg to sync gorutines
	wg.Add(len(slice))
	for _, val := range slice {
		go func(_val int) {
			square := _val * _val
			mu.Lock() //now only one gorutine (one that reached mutex first) can execute code below, others have to wait
			res += square
			mu.Unlock() //other gorutines now can run code above
			wg.Done()
		}(val)
	}
	wg.Wait()
	fmt.Println(res)
}
