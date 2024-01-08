package popcount25

func PopCount(x uint64) int {
	var y uint64
	var c, a int
	for {
		if x = 0 {
			return c
		}
		y = x & (x-1)
		c++
		a = 1 + int(x-y)/2
		x = uint64(x>>a)
	}
}