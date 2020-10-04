package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

// a mock for AnImplementation of AnInterface
type anImplementationMock struct {
	mock.Mock
}

func (m *anImplementationMock) DoSomething(x int) error {
	fmt.Println("Calling a mocked method")
	args := m.Called(x)
	fmt.Println(args)
	return nil
}

func TestDoComplexStuff(t *testing.T) {
	//Prepare a mock
	service1 := new(anImplementationMock)

	// we then define what should be returned from SendChargeNotification
	// when we pass in the value 100 to it. In this case, we want to return nil
	service1.On("DoSomething", 100).Return(nil)
	service1.On("DoSomething", 200).Return(nil)
	// next we want to define the service we wish to test
	sut := MyService{service1, new(OtherImplementation)}
	// and call said method
	sut.DoComplexStuff(100)
	sut.DoComplexStuff(200)
	// at the end, we verify that our myService.ChargeCustomer
	// method called our mocked SendChargeNotification method
	service1.AssertExpectations(t)
}
