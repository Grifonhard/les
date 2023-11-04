package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arrayInt := randArrayInt(100)								//массив
	
	ch := make(chan int64)
	go square(ch)
	for i:=0; i<10 ; i++{
		ch <-arrayInt[i]
	}
}

func randArrayInt(amOfNum int) []int64 {								//создание массива рандомных чисел
	source := rand.NewSource(time.Now().UnixMicro())				
	sRand := rand.New(source)

	array := make([]int64, amOfNum)
	for i:=0; i<amOfNum; i++ {
		array[i]= sRand.Int63n(1000)
	}
	return array
}

func square (iCh chan int64) {
	i := <-iCh
	fmt.Printf("%d = ", i)
	ch := make(chan int64)
	go showRes(ch)
	i *=i
	ch <- i
} 

func showRes (iCh chan int64) {
	fmt.Println(<-iCh)
}
