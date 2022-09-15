package main

import (
	"fmt"
	"strings"
)

/*
The adapter is intended to provide and interface desired and not only what is given by an API

>> A construct which adapts an existing interface X to conform to the required interface Y

*/

/*
Scenario in which is useful the adapter:
Let's suppose that we are using a third party library of figures
*/

func main() {
	rc := NewRectangle(5, 10)
	adapter := NewVectorToRaster(rc)
	fmt.Println(DrawPoints(adapter))
}

// adapter to use DrawPoints with lines

type vectorToRasterAdapter struct {
	points []Point
}

func (a *vectorToRasterAdapter) GetPoints() []Point {
	return a.points
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	fmt.Println("generated", len(a.points), "points")
}

func NewVectorToRaster(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLine(line)
	}
	return &adapter
}

// interface we need, this interfaces need points instead of lines

type Point struct {
	X, Y int
}

type RasterImage interface {
	GetPoints() []Point
}

func DrawPoints(owner RasterImage) string {
	maxX, maxY := 0, 0
	points := owner.GetPoints()
	for _, pixel := range points {
		if pixel.X > maxX {
			maxX = pixel.X
		}
		if pixel.Y > maxY {
			maxY = pixel.Y
		}
	}
	maxX += 1
	maxY += 1

	// preallocate

	data := make([][]rune, maxY)
	for i := 0; i < maxY; i++ {
		data[i] = make([]rune, maxX)
		for j := range data[i] {
			data[i][j] = ' '
		}
	}

	for _, point := range points {
		data[point.Y][point.X] = '*'
	}

	b := strings.Builder{}
	for _, line := range data {
		b.WriteString(string(line))
		b.WriteRune('\n')
	}

	return b.String()
}

// library, the library exist in terms of lines

type Line struct {
	X1, X2, Y1, Y2 int
}

type VectorImage struct {
	Lines []Line
}

func NewRectangle(width, height int) *VectorImage {
	width -= 1
	height -= 1
	return &VectorImage{[]Line{
		{0, 0, width, 0},
		{0, 0, 0, height},
		{width, 0, width, height},
		{0, height, width, height},
	}}
}

// library

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
