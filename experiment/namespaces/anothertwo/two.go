package anothertwo

import (
	"time"

	
	"github.com/Grifonhard/les/tree/main/experiment/namespaces/another"
)

func One(){
	another.Set(10)
	for{
		another.IncrDec()
		time.Sleep(1*time.Second)
	}
}

func Two(){
	another.Set(100)
	for{
		another.IncrDec()
		time.Sleep(1*time.Second)
	}
}