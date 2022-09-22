package main

import "fmt"

func main() {
	circle := Circle{10.12}
	square := Square{10}
	fmt.Println(circle.Render())
	fmt.Println(square.Render())
	redCircle := ColoredShape{&circle, "red"}
	fmt.Println(redCircle.Render())
	redCircleTransparent := TransparentShape{&redCircle, 10.2}
	fmt.Println(redCircleTransparent.Render())
}

// Shape is a base interface
type Shape interface {
	Render() string
}

// Circle is a type of shape
type Circle struct {
	Radius float64
}

func (c *Circle) Render() string {
	return fmt.Sprintf("render circle %v", c.Radius)
}

// Square is a type of shape
type Square struct {
	Side int
}

func (s *Square) Render() string {
	return fmt.Sprintf("render square %v", s.Side)
}

// ColoredShape is a decorator that enhance the shape functionality of Render
type ColoredShape struct {
	Shape Shape
	Color string
}

// Render implementation of colored add the color that is required
func (c *ColoredShape) Render() string {
	return fmt.Sprintf("%s with color %s", c.Shape.Render(), c.Color)
}

// TransparentShape is a decorator that enhance the shape functionality of Render
type TransparentShape struct {
	Shape        Shape
	Transparency float64
}

// Render implementation of transparent add the transparency that is required
func (t *TransparentShape) Render() string {
	return fmt.Sprintf("%s with transparency %v", t.Shape.Render(), t.Transparency)
}
