package main

import "fmt"

func main() {
	s := "marshmellow"
	fmt.Println(s)
	fmt.Println([]byte(s))
	reverse(&s)
	fmt.Println(s)
	fmt.Println([]byte(s))
}

func reverse(str *string) {
	s := []byte(*str)
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	*str = string(s)
}