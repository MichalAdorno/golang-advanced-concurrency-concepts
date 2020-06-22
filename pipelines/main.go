package main

import "fmt"

func main() {
	// Set up the pipeline and consume the output.
	input := []int{1, 2, 3, 4, 5}
	for n := range double(squareRoot(channelizeInput(input...))) {
		fmt.Println(n)
	}
}

/**** Data source ****/
/*
	Note: the input data are pipelined into a first channel,
	which is then passed to functions performing subsequent operations.
*/
func channelizeInput(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

/**** Operations/transformations ****/
/*
	Note: each operation is represented by a function of the following signature:
	`func (in <-chan int) <-chan int`
	where:
	* `in <-chan int` is the result of the previous operation,
	* `<-chan int` is the result of the current operation.
	Instead of returning a result, it returns an "out"-channel.
*/
func squareRoot(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- 2 * n
		}
		close(out)
	}()
	return out
}
