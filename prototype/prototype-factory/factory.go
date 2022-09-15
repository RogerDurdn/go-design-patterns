package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
In this approach we create an API that allows the client to create (factory) a copy (prototype) with some custom choices
*/

func main() {

	roger := NewMainOfficeEmployee("roger", 1)
	laura := NewRemoteOfficeEmployee("laura", 2)

	fmt.Println(roger)
	fmt.Println(laura)
}

type Address struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address
}

// these are base objects take serves as prototypes
var (
	mainOffice   = Employee{"", Address{0, "av 123", "DF"}}
	remoteOffice = Employee{"", Address{0, "palmas 22", "GG"}}
)

// utility function to create a copy with the desire prototype
func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

/*
functions to actually create an object with a desire preset
*/

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewRemoteOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&remoteOffice, name, suite)
}

// DeepCopy is the same as serialization
func (p *Employee) DeepCopy() *Employee {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	d := gob.NewDecoder(&b)
	var result Employee
	_ = d.Decode(&result)
	return &result
}
