package main

import (
	"fmt"
	"strings"
)

/*
The main idea of builders is to have Steps in the creation of
an object, in this sense you can
*/

func main() {
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChild("li", "goo morning")
	b.AddChild("li", "bye")
	fmt.Println(b.String())

	bf := NewHtmlBuilder("ul").
		AddChildFluent("li", "hello f").
		AddChildFluent("li", "goo morning f").
		AddChildFluent("li", "bye f")
	fmt.Println(bf.String())
}

// HtmlBuilder is the actual builder
type HtmlBuilder struct {
	rootTag string
	root    HtmlElement
}

func NewHtmlBuilder(rootTag string) *HtmlBuilder {
	return &HtmlBuilder{rootTag: rootTag,
		root: HtmlElement{rootTag, "",
			[]HtmlElement{}}}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childTag, childText string) {
	e := HtmlElement{childTag, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
}

// AddChildFluent allows to concat multiple calls
func (b *HtmlBuilder) AddChildFluent(childTag, childText string) *HtmlBuilder {
	e := HtmlElement{childTag, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}

const (
	indentationSize = 2
)

// HtmlElement whe use a struct to represent the data that our object could have
type HtmlElement struct {
	tag, text string
	elements  []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

// create the string with recursive indentation
func (e *HtmlElement) string(indentation int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentationSize*indentation)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.tag))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentationSize*(indentation+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}
	for _, element := range e.elements {
		sb.WriteString(element.string(indentation + 1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.tag))
	return sb.String()
}
