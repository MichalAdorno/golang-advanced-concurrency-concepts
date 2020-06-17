package main

import (
	"fmt"
	"sync"
)

var doOnce sync.Once

func main() {
	doSth()
	doSth()
	doOnce.Do(lambda) //the instruction in `lambda` will be run only once, regardless of where it is placed
}

var lambda = func() {
	fmt.Println("Run this: [only once]")
}

func doSth() {
	doOnce.Do(lambda)                    //the instruction in `lambda` will be run only once, regardless of where it is placed
	fmt.Println("Run this: [each time]") //this line is run each time
}
