package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	i            int
	sync.RWMutex //embed RWMutex for lock and read lock
}

func (c *Counter) Read() int {
	c.RLock()
	defer c.RUnlock() //using read lock on reading just in case
	return c.i
}

func (c *Counter) Inc() {
	c.Lock() //lock before incrementing, so no race
	c.i++
	c.Unlock()
}

// tested with -race key
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)
	cnt := Counter{}
	for i := 0; i < 100; i++ {
		go func() {
			cnt.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(cnt.Read())
}
