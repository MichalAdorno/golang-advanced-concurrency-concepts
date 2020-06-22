package main

import (
	"context"
	"fmt"
	"time"
)

/*
	The program shows how to use context cancellation in a sub-task.
	However it contains a logical error worth considering:
	You cannot call `cancel()` twice,
	therfore you cannot use the same context object to cancel multiple goroutines.
*/
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go foo(ctx, "subtask-1")
	go foo(ctx, "subtask-2") //note: one of the goroutines will not be cancelled, the main routine will just cause it to end!
	time.Sleep(100 * time.Millisecond)
	cancel() //note: calling `cancel()` more than once does not yield any results.
	//cancel()
	fmt.Println("Back to the main routine")
}

func foo(ctx context.Context, str string) {
	fmt.Printf("Entering foo() with message: [%s]\n", str)
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Executing foo() with message: [%s]\n", str)
	case <-ctx.Done():
		fmt.Printf("Cancelling foo() with message: [%s]\n", str)
		return
	}
	fmt.Printf("Finishing foo() with message: [%s]\n", str)
}
