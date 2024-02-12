package main

import (
	"fmt"
)

func main() {
	var word1, word2 string
	var isAgrmma bool = true
	fmt.Println("Enter the word 1: ")
	fmt.Scanf("%s/n", &word1)
	fmt.Println("Enter the word 2: ")
	fmt.Scanf("%s/n", &word2)
	w2 := []rune(word2)
	l := len(w2) - 1
	for i, r := range word1 {
		if r != w2[l-i] {
			isAgrmma = false
			break
		}
	}
	fmt.Println(isAgrmma)
}
