package main

type AnInterface interface {
	DoSomething(int) error
}

type OtherInterface interface {
	DoSomethingElse(int) error
}
