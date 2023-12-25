package main

import "fmt"

func main() {
	var f = boilingF
	var c = (f - 32)*5/9
	fmt.Printf("Температура кипения = %g*F или %g*C\n", f, c)
}