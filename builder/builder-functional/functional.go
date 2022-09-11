package main

import "fmt"

/*
this approach allows us to extend the functionality of a existing builder
in a functional way
*/

func main() {
	b := PersonBuilder{}
	b.Called("roger")
	fmt.Println(*b.Build())
}

type Person struct {
	name, position string
}

// personMod is a function used to change the person
type personMod func(*Person)

// PersonBuilder has a slice of actions (personMod) that change the person
type PersonBuilder struct {
	actions []personMod
}

// Called is a normal but inside it creates an action with personMod
func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(person *Person) {
		person.name = name
	})
	return b
}

// Build is a function that iterates through the actions and returns a person
func (b *PersonBuilder) Build() *Person {
	person := &Person{}
	for _, action := range b.actions {
		action(person)
	}
	return person
}
