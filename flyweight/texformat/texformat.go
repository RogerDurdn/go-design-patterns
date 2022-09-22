package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
In this scenario we capitalize a word having a slice that represents the state of every character
this is bad design because we use memory that is not always used
*/

func main() {
	text := NewTextFormatter("hello world of titans")
	text.Capitalize(6, 10)
	fmt.Println(text)
}

type TextFormatter struct {
	text       string
	capitalize []bool
}

func (t *TextFormatter) Capitalize(start, end int) {
	for i, _ := range t.capitalize {
		if i >= start && i <= end {
			t.capitalize[i] = true
		}
	}
}

func NewTextFormatter(text string) *TextFormatter {
	return &TextFormatter{text: text, capitalize: make([]bool, len(text))}
}

func (t *TextFormatter) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(t.text); i++ {
		c := t.text[i]
		if t.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}
