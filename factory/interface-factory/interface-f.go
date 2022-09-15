package main

import "fmt"

/*
This implementation allows us to create a new person without exposing the person struct
with this the encapsulation is preserved.
*/
func main() {
	p := NewPerson("Roger", 30)
	p.SayHello()
	pOld := NewPerson("OldMan", 100)
	fmt.Println()
	pOld.SayHello()
}

// Person is the interface that wraps the SayHello method.
type Person interface {
	SayHello()
}

// NewPerson is the only way to create a new person
// now the person is chosen by the age
func NewPerson(name string, age int) Person {
	if age < 100 {
		return &person{name, age}
	}
	return &tiredPerson{name, age}
}

// person is a private struct that encapsulates the data
type person struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hello, my name is %s and I am %d years old", p.name, p.age)
}

// tiredPerson is an optional implementation of the Person interface
type tiredPerson struct {
	name string
	age  int
}

func (p *tiredPerson) SayHello() {
	fmt.Printf("Hello, my name is %s and I am tired", p.name)
}
