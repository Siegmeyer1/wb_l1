package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(i int, wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()
	for {
		val, ok := <-c //Here workers return when channel is closed. This way we don`t waste time on sending n_of_workers messages to exit channel.
		if !ok {       //It also makes more sense than stopping workers thruogh dedicated exit channel, since we stop using channel c anyways.
			return
		}
		fmt.Println(i, ":", val)

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
	for i := 0; i < n; i++ {
		go worker(i, &wg, c)
	}
	for {
		c <- rand.Intn(100)
		time.Sleep(time.Second / 2)
		select {
		case <-interrupt:
			close(c)
			wg.Wait()
			return
		default:
			continue
		}
	}
}
