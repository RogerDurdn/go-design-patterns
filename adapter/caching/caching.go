package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
)

/*
Within an adapter we generally create a bunch of temporary objects in order to adapt the interfaces, this consumes resources
in every use of the adapter, to solve this we can use cache on these middle temporary objects
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

var pointsCache = map[[16]byte][]Point{}

func hash(obj interface{}) [16]byte {
	bytes, _ := json.Marshal(obj)
	return md5.Sum(bytes)
}

func (a *vectorToRasterAdapter) addLine(line Line) {
	h := hash(line)
	if pts, ok := pointsCache[h]; ok {
		for _, pt := range pts {
			a.points = append(a.points, pt)
		}
		return
	}

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

	pointsCache[h] = a.points

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
