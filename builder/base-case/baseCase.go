package main

import (
	"fmt"
	"strings"
)

/* in this base case scenario we use the string builder in order
to create code for html
*/
func main() {
	writeWord("hey")
	writeListWords([]string{"hey", "you", "dummy"})
}

// If we have a simple use case it doesn't matter
func writeWord(word string) {
	sb := strings.Builder{}
	sb.WriteString("<p>")
	sb.WriteString(word)
	sb.WriteString("</p>")
	fmt.Println(sb.String())
}

// With a more complex use this start to be a mess
func writeListWords(words []string) {
	sb := strings.Builder{}
	sb.WriteString("<ul>")
	for _, word := range words {
		sb.WriteString("<li>")
		sb.WriteString(word)
		sb.WriteString("</li>")
	}
	sb.WriteString("</ul>")
	fmt.Println(sb.String())
}
