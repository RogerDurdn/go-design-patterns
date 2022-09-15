package main

/*
this approach is intended to provide objects pre created
*/

func main() {

}

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "Developer", 100}
	case Manager:
		return &Employee{"", "Manager", 333}
	default:
		panic("unsupported")
	}
}
