package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//Method 1: we use channel to signal gorutine that it have to stop.
	quit := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Method 1 gorutine started")
		for {
			select {
			case <-quit:
				fmt.Println("Method 1 gorutine stopped")
				wg.Done()
				return
			default:
				continue
			}
		}
	}()
	quit <- true
	wg.Wait()
	fmt.Println("|===============================================================|")

	//Method 2: context cancelling
	wg.Add(1)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func(ctx context.Context) {
		fmt.Println("Method 2 gorutine started")
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Method 2 gorutine err: %s\n", err)
			}
			fmt.Println("Method 2 gorutine stopped")
			wg.Done()
			return
		}
	}(ctx)
	cancel()
	wg.Wait()
	fmt.Println("|===============================================================|")

	//Method 3: channel closing
	wg.Add(1)
	ch := make(chan bool)
	go func() {
		fmt.Println("Method 3 gorutine started")
		for {
			if _, ok := <-ch; !ok {
				fmt.Println("Method 3 gorutine stopped")
				wg.Done()
				return
			}
		}
	}()
	close(ch)
	wg.Wait()
	fmt.Println("|===============================================================|")

	//Method 4: set timer with time.After(). Simillar to method 1, but message to time.After channel is sent automatically after set time
	wg.Add(1)
	go func() {
		fmt.Println("Method 4 gorutine started")
		for {
			select {
			case <-time.After(time.Second / 2):
				fmt.Println("Method 4 gorutine stopped")
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
