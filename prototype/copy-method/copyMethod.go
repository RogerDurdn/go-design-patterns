package main

import "fmt"

/*
this approach focus on make a copy of an object in an easy way
not recommended because the complexity and maintainability
*/

func main() {
	mar := Person{"mar",
		&Address{"123 street", "df", "mx"},
		[]string{"john", "mike"}}
	tono := mar.DeepCopy()
	tono.Name = "tono"
	tono.Friends = append(tono.Friends, "yisus")
	fmt.Println(mar, mar.Address)
	fmt.Println(tono, tono.Address)

}

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{a.StreetAddress, a.City, a.Country}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

// DeepCopy is important to copy like value in the first instance
func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}
