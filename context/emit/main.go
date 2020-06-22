package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	/*
		Create a context tree
	*/
	rootCtx := context.Background()

	/*
		Create a child context with its cancellation function (cancel)
	*/
	childCtx, cancel := context.WithCancel(rootCtx)

	// Run two jobs - job-1 is going to fail
	go func() {
		err := job1(childCtx)
		/*
			If this operation returns an error cancel all operations using this context
		*/
		if err != nil {
			cancel()
		}
	}()

	/*
		Run job-2
	*/
	job2(childCtx)
}

func job1(ctx context.Context) error {
	/*
		This job is going to fail after 100 ms.
	*/
	time.Sleep(100 * time.Millisecond)
	return errors.New("Job-1 failed!")
}

func job2(ctx context.Context) {
	/*
		The job-2 depends on the result of job-1,
		so it fails when it receives a cancellation signal
		from the context channel `<-ctx.Done()`
	*/
	select {
	case <-time.After(500 * time.Millisecond): //the key here is that we are waiting for max 500 ms, while job-1 fails in 100 ms.
		fmt.Println("Job-2 done!")
	case <-ctx.Done():
		fmt.Println("Job-2 halted!")
	}
}
