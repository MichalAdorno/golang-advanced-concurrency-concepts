package main

import (
	"errors"
	"fmt"
	"sync"
)

func main() {

	doneChannel := make(chan bool)
	errorChannel := make(chan error)

	var wg sync.WaitGroup
	wg.Add(4)
	go routine("ONE", false, errorChannel, &wg)
	go routine("TWO", true, errorChannel, &wg) //change to `false` to see how the error and done signals are processed in this approach
	go routine("THREE", true, errorChannel, &wg)
	go routine("FOUR", false, errorChannel, &wg) //note only the first error is recorded
	/*
		Using a WaitGroup demands calling `Wait()` upon it in order to create a memory barrier.
		We call `wg.Wait()` in a separate goroutine
		in which we also close the channel for the done signal.

		There are two barriers:
		* `wg` for a pool for goroutines
		* the `select` clause to process results depending on which signal will come from the goroutines

	*/

	go func() {
		wg.Wait()          //a barrier for goroutines
		close(doneChannel) //a barrier for the select statement, see below:
	}()

	select {
	case x := <-doneChannel: //the flow for correct execution of goroutines (done signal)
		//note that we do not send any value to the `doneChannel`,
		//we just wait for the nil value of the channel type (bool): false
		fmt.Printf("-> %v\n", x)
		carryOn()
		break
	case err := <-errorChannel: //the flow for execution with errors (error signal)
		close(errorChannel) //the channel is no longer needed, so it should be closed
		fmt.Printf("[Received error from one of the goroutines: [%v]]\n", err)
	}

}

func routine(routineName string, isError bool, errorChannel chan error, wg *sync.WaitGroup) {
	fmt.Printf("[Message from: %s]\n", routineName)
	if isError {
		err := causeError(routineName)
		if err != nil {
			errorChannel <- err
		}
	}
	wg.Done()
}

func causeError(routineName string) error {
	return errors.New(fmt.Sprintf("Some new error from: [%s]", routineName))
}

func carryOn() {
	//a marker for the instructions that will be called after the waitgroup of goroutines
	fmt.Println("Carry on: exit `select` and do next instructions")
}
