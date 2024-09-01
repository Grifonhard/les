package main

import (
	"fmt"
	"unicode"
)

func main() {
	sb := []byte("asfv   asdfasdf     asdfv   jklsdfg")
	fmt.Println(sb)
	fmt.Println(string(sb))
	delDubl(&sb)
	fmt.Println(sb)
	fmt.Println(string(sb))
}

func delDubl(sb *[]byte){
	for i:=0; i<len(*sb)-1; i++{
		j := i+1
		if unicode.IsSpace(rune((*sb)[i])) && unicode.IsSpace(rune((*sb)[j])){
			for j <= len(*sb) - 2{					
				(*sb)[j] = (*sb)[j + 1]
				j++
			}
			(*sb) = (*sb)[:len(*sb) - 1]
			i--
		}
	}
}