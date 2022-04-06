package main

import (
	"fmt"
	"sync"
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

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	start := time.Now()
	in := gen(1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Distribute the sq work across two goroutines that both read from in.
	// fan-out
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	// fan-in
	singlestream := merge(c1, c2)

	for n := range singlestream {
		fmt.Println(n)
	}
	fmt.Print(">> total time: ", time.Since(start))
}
