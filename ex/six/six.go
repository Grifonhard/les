package main

import (
	"context"
	"fmt"
	"time"
	"sync"
)

func counter (wg sync.WaitGroup, ctx context.Context, i int) {
	var count int
	wg.Add(1)
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("counter №", i , " = ", count)
			return
		default:
			count++
		}
	}
}

func main() {
	wg := sync.WaitGroup{}


	ch := make(chan int)
	defer close(ch)
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()


	for i:=0; i<5; i++{
		go counter(wg, ctx, i+1)
	}
	
	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			fmt.Println("Bye!")
			return
		default:
			time.Sleep(time.Second)
		}
	}
}