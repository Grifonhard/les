package main

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/Grifonhard/les/gopl/ch2/popcount"
	"github.com/Grifonhard/les/gopl/ch2/popcount23"
	"github.com/Grifonhard/les/gopl/ch2/popcount24"
	"github.com/Grifonhard/les/gopl/ch2/popcount25"
)

main() {
	var source = rand.NewSource(time.Now().UnixMicro())														
	var sRand = rand.New(source)
	for {
		x := sRand.Uint64()
		fmt.Println(popcount.PopCount(x))
		fmt.Println(popcount23.PopCount(x))
		fmt.Println(popcount24.PopCount(x))
		fmt.Println(popcount25.PopCount(x))
		time.Sleep(time.Second)
	}
}