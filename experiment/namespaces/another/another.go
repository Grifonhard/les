package another

import (
	"github.com/Grifonhard/les/tree/main/experiment/namespaces/count"
	"github.com/Grifonhard/les/tree/main/experiment/namespaces/funct"
)

func Set(n int){
	count.Count = n
	if count.Count < 20{
			count.Count-- 
	}
	funct.Show()
}

func IncrDec(){
	if count.Count < 20{
		count.Count-- 
	} else {
		count.Count++
	}
	funct.Show()
}