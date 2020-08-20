package feature

import (
	"fmt"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/20 18:38
// @description
// @version

func LearnGoroutineBench() {
	fmt.Println("sync:", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered:", testing.Benchmark(BenchmarkChannelBuffered).String())
}
func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}

func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 128)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for range ch {
	}
}
