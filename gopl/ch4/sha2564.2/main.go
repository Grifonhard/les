package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("adsaf"))
	c2 := sha256.Sum256([]byte("adsad"))
	var count int
	for i := 0; i < 32; i++ {
		count += PopCount((uint64(c1[i]) ^ uint64(c2[i])))
	}
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(count)
}

func PopCount(x uint64) int {
	var y uint64
	var c, a int
	for {
		y = x & (x - 1)
		c++
		y = x - y
		for {
			y = y / 2
			a++
			if y == 0 {
				break
			}
		}
		x = uint64(x >> a)
		a = 0
		if x == 0 {
			return c
		}
		if x == 1 {
			c++
			return c
		}
	}
}
