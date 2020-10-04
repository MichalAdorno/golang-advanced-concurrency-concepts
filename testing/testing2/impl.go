package main

import (
	"errors"
	"fmt"
)

//----------------------------------------------------------------
type AnImplementation struct{}

// This is the method we are going to mock
func (i AnImplementation) DoSomething(x int) error {
	if x > 1000 {
		return errors.New("An error")
	}
	fmt.Printf("Print it: \"Doing the real thing with: [%v]\"", x)
	return nil
}

//----------------------------------------------------------------
type OtherImplementation struct{}

func (i OtherImplementation) DoSomethingElse(x int) error {
	if x > 10000 {
		return errors.New("An error")
	}
	fmt.Printf("Print it: \"Doing the OTHER real thing with: [%v]\"", x)
	return nil
}

//----------------------------------------------------------------
