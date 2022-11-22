package main

import (
	"fmt"
	"time"
)

func MySleep(t time.Duration) {
	start := time.Now()         // remember when we started
	for time.Since(start) < t { // loop until needed time since start passed
	}
}

func main() {
	start := time.Now()
	MySleep(time.Second)
	fmt.Println("slept for:", time.Since(start))
}
