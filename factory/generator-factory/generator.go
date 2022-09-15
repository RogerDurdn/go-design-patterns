package main

import "fmt"

/*
This example shows how to create a factory that can determine what type of object to create based on the input.
this can be done with 2 approaches:
* using functional approach: this returns a function that creates the object
* using the struct approach: this returns an instance of a factory that can create the object

*/

func main() {
	// using the functional approach
	developerFactory := NewEmployeeFactory("Developer", 30000)
	managerFactory := NewEmployeeFactory("Manager", 500)

	developer := developerFactory("Developer Roger")
	manager := managerFactory("Manager Tom")
	fmt.Println(developer)
	fmt.Println(manager)
	devFactory := NewEmployeeFactorys("devs", 1000)
	developa := devFactory.Create("juan")
	fmt.Println(developa)
}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// NewEmployeeFactory is the functional approach
func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, income}
	}
}

// with this approach is needed to crete a struct that can hold the data
// also is needed to have a function that creates the object
// also we need a constructor "NewEmployeeFactorys" which helps to create the factory

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func (ef *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, ef.Position, ef.AnnualIncome}
}

func NewEmployeeFactorys(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}
