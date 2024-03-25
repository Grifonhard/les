package main

import (
	"time"

	"github.com/Grifonhard/les/tree/main/experiment/namespaces/anothertwo"	
)

func main(){
	
	go anothertwo.One()

	go anothertwo.Two()

	time.Sleep(600*time.Second)
}
