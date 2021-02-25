package testing

import (
	"fmt"
	"testing"
	"time"
)

func TestBenchmarkBuffer(t *testing.T) {
	fmt.Println(" sync", testing.Benchmark(BenchmarkBuffer).String())
}

func BenchmarkBuffer(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			time.Sleep(1e8)
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
