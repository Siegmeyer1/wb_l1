package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	fmt.Print("Enter N of seconds to stop program after: ")
	fmt.Scanln(&n)
	ch := make(chan int, 1)
	exit := make(chan bool, 1)

	go func(n int) { //running gorutine with timer, that will send to exit channel when timer runs out
		time.Sleep(time.Duration(n) * time.Second)
		exit <- true
	}(n)

	i := 0
	for {
		select {
		case <-exit: //checking for value in exit channel. If there is value, then terminate
			return
		case val := <-ch: //print value in channel ch if there is any
			fmt.Println(val)
		default: //if not exiting and have nothing to print from channel ch, send next value to that channel
			ch <- i
			i++
			time.Sleep(time.Second / 10)
		}
	}
}
