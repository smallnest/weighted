package weighted

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkW_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := &W{}
	for i := 0; i < 10; i++ {
		w.Add("server"+strconv.Itoa(i), rand.Intn(100))
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}
