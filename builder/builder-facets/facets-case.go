package main

import "fmt"

/*
In this case scenario we are interested in have multiple builders because the object
is way more complicated
*/

type Person struct {
	// address
	StreetAddress, City string
	// job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

/*
Lives
Utility methods for building address on the main builder
*/
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

/*
PersonAddressBuilder
In order to modularize the builder we need to use different builders for the address and job
a way to do this is to aggregate the personBuilder into the new builder
*/
type PersonAddressBuilder struct {
	PersonBuilder
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (b *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	b.person.CompanyName = companyName
	return b
}

func (b *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	b.person.Position = position
	return b
}

func (b *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	b.person.AnnualIncome = income
	return b
}

func main() {
	pb := NewPersonBuilder().
		Lives().At("apt 123").In("chicago").
		Works().At("FANNG").AsA("doctor").Earning(1222)
	fmt.Println(pb.Build())
}
