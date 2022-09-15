package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
This approach allows us to create a copy of an object and everything inside in the right way,
in order to do this we have to use some packages like bytes and encoding/gob
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

// DeepCopy serialize and deserialize the object into a new one
func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)
	fmt.Println(string(b.Bytes()))
	d := gob.NewDecoder(&b)
	var result Person
	_ = d.Decode(&result)
	return &result
}

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}
