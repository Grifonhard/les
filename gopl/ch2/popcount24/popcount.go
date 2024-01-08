package popcount24

func PopCount(x uint64) int {
	var c int
	for i := 0; i<=64; i++ {
		c+=int(byte(x>>i)&1)
	}
	return c
}