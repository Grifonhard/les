package main

import (
	"math/rand"
	"time"
	"fmt"
	"gopl/ch2/popcount"
	"gopl/ch2/popcount23"
	"gopl/ch2/popcount24"
	"gopl/ch2/popcount25"
)

func main() {
	var source = rand.NewSource(time.Now().UnixMicro())														
	var sRand = rand.New(source)
	for {
		x := sRand.Uint64()
		fmt.Println(popcount.PopCount(x))
		fmt.Println(popcount23.PopCount(x))
		fmt.Println(popcount24.PopCount(x))
		fmt.Println(popcount25.PopCount(x))
		time.Sleep(5*time.Second)
	}
}