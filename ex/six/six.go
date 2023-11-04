package main

import (
	"fmt"
	"context"
)

func counter (ctx context.Context, i int) {
	select {
		case <-ctx
	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel


}