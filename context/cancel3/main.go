package main

import (
	"context"
	"fmt"
)

/*
	See godoc: https://golang.org/pkg/context/#WithCancel
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for n := range yield(1, ctx) {
		fmt.Println(n)
		if n == 5 { //We want to consume only this number of elements from the channel
			// (1)
			//break
			// cancel()
			break
		}
	}

	// (2)
	cancel() //or alternatively: defer cancel() // cancel when we are finished consuming integers
}

/*
	`yield` generates integers in a separate goroutine and
	sends them to the returned channel.
	The callers of gen need to cancel the context once
	they are done consuming generated integers not to
	leak the internal goroutine started by gen.
*/
func yield(start int, ctx context.Context) <-chan int {
	dst := make(chan int)
	n := start
	go func() {
		for {
			select {
			case <-ctx.Done():
				// (3)
				// close(dst)
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
