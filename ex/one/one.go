package main

import(
	"fmt"
	"time"
)

func worker(ch chan int64, i int) {
	for {
		fmt.Println("worker №", i+1, ": ", <-ch)
	}
}

func main(){
	ch := make(chan int64)

	for i:=0; i<10; i++ {
		go worker(ch, i)
	}

	for {
		ch <- time.Now().Unix()
		time.Sleep(time.Second)
	}
}