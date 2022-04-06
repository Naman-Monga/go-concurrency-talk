package main

import (
	"fmt"
	"time"
)

type Partner struct {
	Name string
	Eta  time.Duration
}

// generic return list of responses within the time out
func getAllResponses(timeout time.Duration, list []interface{}, fc func(it interface{}) interface{}) ([]interface{}, error) {
	responses := make(chan interface{}, len(list))

	for _, i := range list {
		item := i
		go func() {
			responses <- fc(item)
		}()
	}

	var outputList []interface{}
	t := time.NewTimer(timeout)
	for range list {
		select {
		case response := <-responses:
			outputList = append(outputList, response)
		case <-t.C:
			return outputList, fmt.Errorf("func_timed_out")
		}
	}

	return outputList, nil
}

func main() {
	// input data
	start := time.Now()
	timeout := 6 * time.Second
	partners := []interface{}{
		Partner{Name: "dunzo", Eta: 2 * time.Second},
		Partner{Name: "shadowfax", Eta: 3 * time.Second},
		Partner{Name: "rapido", Eta: 5 * time.Second},
	}

	// get all responses
	resp, err := getAllResponses(timeout, partners, checkConfig)
	if err != nil {
		fmt.Println("Error Message: ", err)
	}

	fmt.Print("final: [")
	for _, r := range resp {
		str := r.(string)
		fmt.Print(str, ", ")
	}
	fmt.Println("]")

	elapsed := time.Since(start)
	fmt.Println(">> time elapsed: ", elapsed)
	time.Sleep(5 * time.Second) // waiting just to see everything afterwards
	fmt.Println("*** all done ***")
}

// input func
func checkConfig(partner interface{}) interface{} {
	p := partner.(Partner)
	time.Sleep(p.Eta)
	p.Name = p.Name + "|x" // appending x means successfuly processed
	return p.Name
}
