package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	http.ListenAndServe(":9000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		/*
			This prints to the standard output to show that processing has started
		*/
		fmt.Fprint(os.Stdout, "Request processing has just started!\n")
		/*
			We use `select` to execute a piece of code depending
			on which channel receives a message first
		*/
		select {
		/*
			In the example, we assume that it takes 2 seconds to process the imaginary request correctly
		*/
		case <-time.After(2 * time.Second):
			fmt.Fprint(os.Stdout, "Request processed!\n")
			w.Write([]byte("Request processed!"))

		/*
			If the request gets cancelled by interrupting the REST client that sends the request,
			then, we the `ctx` object informs us by the output from the channel`ctx.Done()`
			that the request was cancelled or processed incorrectly.
		*/
		case <-ctx.Done():
			fmt.Fprint(os.Stderr, "Request cancelled or processed incorrectly!\n")
		}
	}))
}
