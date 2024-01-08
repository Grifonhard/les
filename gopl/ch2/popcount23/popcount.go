package popcount23

var pc[256] byte

func init() {
	for i:=range pc {
		pc[i] = pc[i/2]+byte(i&1)
	}
}

func PopCount(x uint64) int {
	var c int
	for i:=0; i<8; i++{
		c+=int(pc[byte(x>>(i*8))])
	}
	return c
}