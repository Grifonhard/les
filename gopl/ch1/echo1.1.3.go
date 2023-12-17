package main

import (
	"testing"
	"os"
	"fmt"
	"strings"
) 

func main () {
	fmt.Println(echo1())
	fmt.Println(echo2())
	fmt.Println(echo3())
}

func echo1() string {
	var s, sep string
	for i:=1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	s, sep :="", ""
	for _, arg := range os.Args[1:] {
		s+=sep +arg
		sep = " "
	}
	return s
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}


func Benchmarkecho1 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}
func Benchmarkecho2 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}
func Benchmarkecho3 (b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}