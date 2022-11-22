package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup, c <-chan int, exit <-chan bool) {
	defer wg.Done()
	for {
		select {
		case val := <-c:
			fmt.Println(i, ":", val)
		case <-exit: //Here we stop workers using exit channel. This is not most appropriate way for this exact task, but I implemented it
			return //as a trainnig anyways.
		}
	}
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt) //here we make "interrupt" channel "catch" ^C signals
	c := make(chan int)
	var n int
	fmt.Print("Enter num of workers: ")
	fmt.Scanln(&n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	exit := make(chan bool, n)
	for i := 0; i < n; i++ {
		go worker(i, &wg, c, exit)
	}
	for {
		c <- rand.Intn(100)
		time.Sleep(time.Second / 2)
		select {
		case <-interrupt:
			for i := 0; i < n; i++ {
				exit <- true
			}
			wg.Wait()
			return
		default:
			continue
		}
	}
}
