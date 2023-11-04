package main

import(
	"fmt"
	"time"
	"flag"
)

func worker(ch chan int64, i int) {
	for {
		fmt.Println("worker №", i+1, ": ", <-ch)
	}
}

func main(){
	ch := make(chan int64)
	wn := flag.Int("wn", 15,"an int")


	flag.Parse()
	for i:=0; i<*wn; i++ {
		go worker(ch, i)
	}

	for {
		ch <- time.Now().Unix()
		time.Sleep(time.Second)
	}
}