package main

import (
	"fmt"
	"sync"
)

func main() {
	m := make(map[int]int)
	m[1] = 42
	mu := sync.RWMutex{}   // RWmutex allows to set read lock. Other gorutines will be able to get read lock too, so we enable concurrent reading
	wg := sync.WaitGroup{} // while protecting ourselves from concurrent reading and writing. In case of concurrent R&W map behavior is undefined
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {

			mu.Lock()
			m[0] = i
			mu.Unlock()

			mu.RLock()
			val := m[1]
			mu.RUnlock()

			fmt.Printf("Worker %d now know %d\n", i, val)
			fmt.Printf("Worker %d done\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
