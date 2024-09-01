package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(a)
	shift := 3
	rotate(&a, shift)

	fmt.Println(a)
}

func rotate(s *[]int, shift int) {
	first := (*s)[:shift]
	(*s) = (*s)[shift:]
	(*s) = append((*s), first...)
}