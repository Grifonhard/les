package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"flag"
)

func main() {
	sec := flag.Int("sec", 15,"an int")															//получение времени работы из флага
	flag.Parse()


	ctx, cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Second*time.Duration(*sec)))	//создание контекста
	defer cancel()

	
	source := rand.NewSource(time.Now().UnixMicro())											//отправляемые данные			
	sRand := rand.New(source)


	ch := make(chan int)
	defer close(ch)
	go showInt(ctx, ch)


	var i int
	for i<1 {
		select {
		case <-ctx.Done():																		//break не срабатывает
			i = 1
		default:
			ch <- sRand.Int()
		}
	}
}

func showInt (ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case r := <-ch:
			fmt.Println(r)
		}
	}
}