package main

import (
	"fmt"
)

type MyService struct {
	service1 AnInterface
	service2 OtherInterface
}

func (a MyService) DoComplexStuff(value int) error {
	a.service1.DoSomething(value)
	fmt.Printf("Inside realDoComplexStuff with [%v]\n", value)
	a.service2.DoSomethingElse(value)
	return nil
}

func main() {
	fmt.Println("Production")

	service1 := AnImplementation{}
	service2 := OtherImplementation{}
	myService := MyService{service1, service2}
	myService.DoComplexStuff(100)
}
