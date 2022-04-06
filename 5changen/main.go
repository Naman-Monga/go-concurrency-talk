package main

import (
	"fmt"
	"time"
)

func networkCall() bool {
	time.Sleep(4 * time.Second)
	return true
}

// channel generator
func myRoutine() chan bool {
	resp := make(chan bool, 1)
	go func() {
		resp <- networkCall()
		close(resp)
		return
	}()
	return resp
}

func main() {
	start := time.Now()

	resp := myRoutine()
	response := <-resp // wait

	fmt.Println(response, time.Since(start))
	return
}
