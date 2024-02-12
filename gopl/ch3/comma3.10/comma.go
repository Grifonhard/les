package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	var word string
	fmt.Println("Enter the word: ")
	fmt.Scanf("%s/n", &word)
	for i, r := range word {
		buf.WriteRune(r)
		if (i+1)%3 == 0 {
			buf.WriteRune(',')
		}
	}
	fmt.Println(buf.String())
}
