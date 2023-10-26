package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int64)
	ch2 := make(chan string)
	
	go randArrayInt(ch1)								
	go square(ch1, ch2)

	for i:=0; i<10; i++ {
		fmt.Println(<-ch2)
	}
}

func randArrayInt(ch chan int64) {								
	source := rand.NewSource(time.Now().UnixMicro())				
	sRand := rand.New(source)
	for {
		iRand := sRand.Int63n(10000)
		ch <- iRand
	}
}

func square (ch1 chan int64, ch2 chan string) {
	for {
		i := <- ch1		
		ch2 <- fmt.Sprintf("%d = %d", i, i*i)
	}
} 


