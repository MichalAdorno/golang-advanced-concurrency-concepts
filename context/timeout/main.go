package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	/*
		Create a new context with a deadline of 100 milliseconds.
		If 100 ms -> most likely you will get a timeout error passed through the context.
		If eg 500 ms -> most likely you will get 200 OK and no context cancellation.
	*/
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 100*time.Millisecond)

	// Make a request, that will call some HTTP server
	req, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
	// Associate the cancellable context we just created to the request
	req = req.WithContext(ctx)
	// Create a new HTTP client and execute the request
	res, err := (&http.Client{}).Do(req)
	// If the request failed, log to the standard output
	if err != nil {
		fmt.Println("Request failed:", err) //`err` contains the info about time limit excess
		return
	}
	// Print the statuscode if the request succeeds
	fmt.Println("Response received with status code:", res.StatusCode)
}
