package weighted

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func BenchmarkSW_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := &SW{}
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.Intn(100)+100)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}

func BenchmarkRRW_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := &RRW{}
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.Intn(100)+100)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}

func BenchmarkRandW_Next(b *testing.B) {
	b.ReportAllocs()
	rand.Seed(time.Now().UnixNano())
	w := NewRandW()
	for i := 0; i < 50; i++ {
		w.Add("item-"+strconv.Itoa(i), rand.Intn(100)+100)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		w.Next()
	}
}
