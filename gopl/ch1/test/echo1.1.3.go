package echo

import (
	"os"
	"strings"
) 

func Echo1() string {
	var s, sep string
	for i:=1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func Echo2() string {
	s, sep :="", ""
	for _, arg := range os.Args[1:] {
		s+=sep +arg
		sep = " "
	}
	return s
}

func Echo3() string {
	return strings.Join(os.Args[1:], " ")
}