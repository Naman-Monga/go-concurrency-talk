package main

import (
	"fmt"
	"time"
)

func main() {
	name := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		name <- "naman"
		time.Sleep(1 * time.Second)
		fmt.Println("func 1 exited")
	}()

	go func() {
		n := <-name
		fmt.Println("func 2 exited, name: ", n)
	}()

	fmt.Println("** go routines created **")

	// wait for everything to complete
	time.Sleep(5 * time.Second)
	fmt.Println("** exited main **")
}
