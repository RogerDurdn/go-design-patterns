package main

import "fmt"

type Color int

const (
	green Color = iota
	blue
	red
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	Name string
	Color
	Size
}

type Filter struct {
}

// ByColor use pointer because return the same data that is given
func (f Filter) ByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// BySize case example of a bad design, if now is needed by size, not OCP
func (f Filter) BySize(products []Product, color Color) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if product.Color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// using "Specification pattern"

type Specification interface {
	IsSatisfied(product *Product) bool
}

type ColorSpecification struct {
	Color
}

func (c ColorSpecification) IsSatisfied(product *Product) bool {
	return product.Color == c.Color
}

// SizeSpecification if we need a new type of filter we create a new specification
type SizeSpecification struct {
	Size
}

func (s SizeSpecification) IsSatisfied(product *Product) bool {
	return product.Size == s.Size
}

// AndSpecification multi spec using a composite approach
type AndSpecification struct {
	one, second Specification
}

func (cs AndSpecification) IsSatisfied(product *Product) bool {
	return cs.one.IsSatisfied(product) && cs.second.IsSatisfied(product)
}

type GoodFilter struct {
}

func (gf GoodFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if spec.IsSatisfied(&product) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	products := []Product{
		Product{Name: "cup", Color: green, Size: small},
		Product{Name: "bowl", Color: green, Size: medium},
		Product{Name: "shirt", Color: blue, Size: medium},
	}

	filter := Filter{}

	onlyGreen := filter.ByColor(products, green)
	fmt.Println("Result of filter with bad design:")
	for _, product := range onlyGreen {
		fmt.Printf("- %s product is green \n", product.Name)
	}

	goodFilter := GoodFilter{}
	greenSpec := ColorSpecification{green}
	mediumSpec := SizeSpecification{medium}
	greenMediumSpec := AndSpecification{greenSpec, mediumSpec}

	fmt.Println("Result Good filter:")
	fmt.Println("- Only green")
	PrintArr(goodFilter.Filter(products, greenSpec))
	fmt.Println("- Only Medium")
	PrintArr(goodFilter.Filter(products, mediumSpec))
	fmt.Println("- Only Green Medium")
	PrintArr(goodFilter.Filter(products, greenMediumSpec))
}

func PrintArr(products []*Product) {
	for _, product := range products {
		fmt.Printf("- - :%v\n", product.Name)
	}
}
