package popcount24

import "testing"
import "math/rand"
import "time"


var source = rand.NewSource(time.Now().UnixMicro())														
var sRand = rand.New(source)

func BenchmarkPopCount (b *testing.B) {
	for i :=0 ; i<=b.N; i++ {
		PopCount(sRand.Uint64())
	}
}