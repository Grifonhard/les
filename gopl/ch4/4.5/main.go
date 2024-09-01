package main

import "fmt"

func main() {
	sl_str := []string{"aaa", "aa", "aa", "qerqwfa", "afdiovz", "dddd", "dddd", "dddd", "orweq"}
	fmt.Println(sl_str)
	delDubl(&sl_str)
	fmt.Println(sl_str)
}

func delDubl(s *[]string){
	for i:=0; i<len(*s)-1; i++{
		j := i+1
		if (*s)[i] == (*s)[j]{
			for j <= len(*s) - 2{					
				(*s)[j] = (*s)[j + 1]
				j++
			}
			(*s) = (*s)[:len(*s) - 1]
			i--
		}
	}
}