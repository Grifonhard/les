package popcount25


func PopCount(x uint64) int {
	var y uint64
	var c, a int
	for {
		y = x & (x-1)
		c++
		y = x - y
		for {
			y = y/2
			a++
			if y == 0 {
				break
			}
		}	
		x = uint64(x>>a)
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