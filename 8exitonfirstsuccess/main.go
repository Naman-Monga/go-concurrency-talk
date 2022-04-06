package main

import (
	"fmt"
	"time"
)

type Partner struct {
	Name string
	Eta  time.Duration
}

// channel generator
func gen(list []interface{}, fc func(it interface{}) interface{}) <-chan interface{} {
	resp := make(chan interface{}, len(list))
	for _, i := range list {
		item := i
		go func() {
			resp <- fc(item)
		}()
	}
	return resp
}

// generic first success return
func getFirstSuccess(timeout time.Duration, list []interface{}, fc func(it interface{}) interface{}) interface{} {

	// create channel buffer
	responses := gen(list, fc)

	t := time.NewTimer(timeout)
	for i := len(list); i > 0; i-- {
		select {
		case resp := <-responses:
			if resp != nil {
				return resp
			}
		case <-t.C:
			fmt.Println("timeout")
			return nil
		}
	}
	return nil
}

func main() {
	// input data (config of a store)
	start := time.Now()
	timeout := 4 * time.Second
	partners := []interface{}{
		Partner{Name: "dunzo", Eta: 2 * time.Second},
		Partner{Name: "shadowfax", Eta: 3 * time.Second},
		Partner{Name: "rapido", Eta: 10 * time.Second},
	}

	// in our case: to check if order is serviceable for a store
	if resp := getFirstSuccess(timeout, partners, checkServiceability); resp != nil {
		fmt.Println("Final: ", resp)
	} else {
		fmt.Println("Final: Failed")
	}

	elapsed := time.Since(start)
	fmt.Println(">> time elapsed: ", elapsed)
	time.Sleep(10 * time.Second) // waiting just to see everything afterwards
	fmt.Println("*** all done ***")
}

// input func
func checkServiceability(partner interface{}) interface{} {
	p := partner.(Partner)
	time.Sleep(p.Eta)
	if p.Name == "dunzo" {
		return nil
	}
	return p
}
