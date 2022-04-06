package main

import (
	"fmt"
	"sync"
	"time"
)

const timeout time.Duration = 2 * time.Second

func main() {
	response := make(chan string, 2)
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		response <- catalogSync("zomato")
	}()
	go func() {
		defer wg.Done()
		response <- catalogSync("swgy")
	}()

	fmt.Println("waiting for both routines ...")
	wg.Wait()

	a := <-response
	b := <-response
	fmt.Println(a)
	fmt.Println(b)

	return
}

func catalogSync(s string) string {
	time.Sleep(time.Duration(len(s)) * time.Second)
	return fmt.Sprintf("> %s_done", s)
}
