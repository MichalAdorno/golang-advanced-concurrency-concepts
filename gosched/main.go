package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func say(s string, wg *sync.WaitGroup, enableGosched bool) {
	for i := 0; i < 5; i++ {
		if enableGosched {
			runtime.Gosched() //change context (yield) - stops and saves the current goroutine, switches to another one
		}
		fmt.Printf("[%s - %d]\n", s, i)
	}
	wg.Done()
}

func main() {
	//`GOMAXPROCS` can also influence goroutine execution.
	//`runtime.Gosched()` results are best visible when `GOMAXPROCS == 1`
	//runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(4)
	fmt.Println("************** With Goroutine Yield (runtime.Gosched()) **************")
	wg.Add(2)
	go say("world", &wg, true) // create a new goroutine
	say("hello", &wg, true)    // current goroutine
	wg.Wait()
	fmt.Println("************** No Goroutine Yield**************")
	wg.Add(2)
	go say("world", &wg, false) // create a new goroutine
	say("hello", &wg, false)    // current goroutine
	wg.Wait()
	fmt.Println("************** END **************")
}
