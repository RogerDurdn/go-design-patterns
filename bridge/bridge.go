package main

import "fmt"

/*
This pattern prevents the "Cartesian product"
complexity explosion
Example:
  - Common type threadScheduler
  - Can be preemtive or cooperative
  - Can run on Windows of Unix
  - End up with a 2x2 scenario: WindowsPTS, UnixPTS
    WindowsCTS, UnixCTS

# Bridge pattern avoids the entity explosion

>> A mechanism that decouples an interface (hierarchy)

	from an implementation (hierarchy)

we avoid the problem by introduce a kind of dependency on the class that needs to be capable
of different behavior
*/
func main() {
	raster := RenderRaster{10}
	vector := &RenderVector{}
	circle := NewCircle(&raster, 30)
	circle.Draw()
	circle.render = vector
	circle.Resize(20)
	circle.Draw()

}

/* The use case is based on a shape that has to be renderer in a way, this kind of shapes are
different and needs a different wa y to be handled

types of shapes: circle, square
types fo renders: raster, vector

this can can generate a relation like:
rasterCircle
rasterSquare
vectorCircle
vectorSquare
...
so how we simplify this: Bridge design pattern
*/

type Circle struct {
	render Renderer
	radius float32
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{render: renderer, radius: radius}
}

func (c *Circle) Draw() {
	c.render.RenderCircle(c.radius)
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

type Renderer interface {
	RenderCircle(radius float32)
}

type RenderVector struct {
	//
}

func (v *RenderVector) RenderCircle(radius float32) {
	fmt.Println("Draw circle with vector:", radius)
}

type RenderRaster struct {
	Dpi int
}

func (v *RenderRaster) RenderCircle(radius float32) {
	fmt.Println("Draw circle with raster:", radius)
}
