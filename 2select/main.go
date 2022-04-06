package main

import (
	"fmt"
	"time"
)

const timeout time.Duration = 2 * time.Second

func main() {
	response := make(chan string, 1)
	go func() {
		response <- networkCall()
	}()

	// what if we wanna implement a timeout?

Loop:
	for {
		select {
		case resp := <-response:
			fmt.Println(resp)
			break Loop
		case <-time.After(timeout):
			fmt.Println("timeout")
			break Loop
		}
	}

	return
}

func networkCall() string {
	time.Sleep(1 * time.Second)
	return "naman"
}
