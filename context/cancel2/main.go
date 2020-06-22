package main

import (
	"fmt"
	"time"
)

/*
	The program shows how to cancel dependent goroutines.
	It is the counterpart of the `context/cancel1` module.
	In this program, we use a buffered channel
*/
func main() {
	//Create a cancellation channel
	cancelChannel := make(chan bool, 2)
	//and a completion channel
	completionChannel := make(chan bool, 2)
	//Run goroutines
	go foo(cancelChannel, completionChannel, "subtask-1")
	go foo(cancelChannel, completionChannel, "subtask-2") //note: one of the goroutines will not be cancelled, the main routine will just cause it to end!

	time.Sleep(100 * time.Millisecond)
	//Load the buffer with cancellation signals
	cancelChannel <- true
	cancelChannel <- true
	//Synchronize main with the goroutines
	<-completionChannel
	<-completionChannel
	fmt.Println("Back to the main routine")
}

func foo(cancel <-chan bool, completionChannel chan bool, str string) <-chan bool {
	fmt.Printf("Entering foo() with message: [%s]\n", str)
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Executing foo() with message: [%s]\n", str)
	case completionChannel <- <-cancel:
		fmt.Printf("Cancelling foo() with message: [%s]\n", str)
		return completionChannel
	}
	fmt.Printf("Finishing foo() with message: [%s]\n", str)
	completionChannel <- true
	return completionChannel
}
