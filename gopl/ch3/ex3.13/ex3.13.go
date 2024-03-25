package main

import (
	"bytes"
	"fmt"
)

const(
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main(){
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)

	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d", ZB)  //в лоб функции чтения не принимают слишком большие неименованые инты, надо либо пользоваться пакетом math, либо самому делать функцию преобразующую инты в строки
	fmt.Println(string(YB))
}