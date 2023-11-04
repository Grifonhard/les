package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second*time.Duration(3)))	//время выполнения
	defer cancel()

	
	ch1 := make(chan int64)
	defer close(ch1)
	ch2 := make(chan int64)
	defer close(ch2)
	
	go randArrayInt(ctx, ch1)								
	go square(ctx, ch1, ch2)

	var i int
	for i<1 {
		select {
		case <- ctx.Done():
			i=1
		default:	
			fmt.Println(<-ch2)
		}
	}
}

func randArrayInt(ctx context.Context, ch1 chan int64) {									
	source := rand.NewSource(time.Now().UnixMicro())				
	sRand := rand.New(source)
	for {
		select {
		case <- ctx.Done():
			break
		default:	
			iRand := sRand.Int63n(10000)
			ch1 <- iRand
		}
	}
}

func square (ctx context.Context, ch1 chan int64, ch2 chan int64) {
	for {
		select {
		case <- ctx.Done():
			break
		default:	
			i := <- ch1		
			ch2 <- i*i
		}
	}
} 