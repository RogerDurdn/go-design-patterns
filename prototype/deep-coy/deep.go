package main

import "fmt"

/*
This pattern needs to be able to copy in deep the values of the object

if we want to use this approach we have to copy and create news objects when the case is a pointer
for example slices, maps and references *
*/

func main() {
	pepe := Person{"peep", &Address{"123 street", "cdmx", "mx"}}
	laura := pepe

	laura.Address = &Address{
		StreetAddress: pepe.Address.StreetAddress,
		City:          pepe.Address.City,
		Country:       pepe.Address.Country,
	}
	laura.Address.StreetAddress = "some street"
	fmt.Println(pepe, pepe.Address)
	fmt.Println(laura, laura.Address)
}

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}
