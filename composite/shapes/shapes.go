package main

import (
	"fmt"
	"strings"
)

/*
In this implementation we are using a struct that have a collection of items that are the same,
and we use a recursive method that iterates through the collections of every single item
in this sense we can apply the changes (draw) to all the composite
*/

func main() {
	item1 := NewGraphicItem("it1", "red")
	item1.Draw()
	item2 := NewGraphicItem("it2", "blue")
	item3 := NewGraphicItem("it3", "green")
	item3.child = append(item3.child, item1)
	item3.child = append(item3.child, item2)
	item3.Draw()

	item4 := NewGraphicItem("it4", "yellow")
	item4.child = append(item4.child, item3)
	item4.Draw()

}

// GraphicItem is the base struct that hase a collection of themselves
type GraphicItem struct {
	Name, Color string
	child       []*GraphicItem
}

func NewGraphicItem(name, color string) *GraphicItem {
	return &GraphicItem{name, color, nil}
}

func (gi *GraphicItem) Draw() {
	gi.draw(0)
}

// draw is a
func (gi *GraphicItem) draw(deep int) {
	text := fmt.Sprintf("%s %s %s", strings.Repeat("*", deep), gi.Name, gi.Color)
	fmt.Println(text)
	for _, item := range gi.child {
		item.draw(deep + 1)
	}
}
