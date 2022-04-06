package main

import (
	"fmt"
	"time"
)

// use cases
// timers: you wait for a program to execute
// tickers: you wanna repeatedly execute after x duration

func main() {
	// timerFunc()
	tickerFunc()
}

func timerFunc() {
	t := time.NewTimer(2 * time.Second)
	start := time.Now()
	<-t.C
	fmt.Println(">> click after: ", time.Since(start))
}

func tickerFunc() {
	interval := 1 * time.Second

	t := time.NewTicker(interval)

	go func() {
		for t := range t.C {
			fmt.Println("tick at: ", t)
		}
	}()

	time.Sleep(4 * time.Second) //wait for 4 seconds
	t.Stop()                    // stop the ticker

	time.Sleep(3 * time.Second)
	fmt.Println(">> exiting ticker func")
}
