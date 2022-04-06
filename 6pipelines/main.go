package main

import (
	"fmt"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline.
	start := time.Now()
	c := gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	out := sq(c)

	// Consume the output.
	for o := range out {
		fmt.Println(o)
	}
	fmt.Print(">> total time: ", time.Since(start))
}
