package feature_test

import (
	"github.com/bradfitz/iter"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/24 16:56
// @description 测试一下iter包中迭代函数的性能
// @version
const iterN = 50000

func myIter(n int) []struct{} {
	return make([]struct{}, n)
}

func BenchmarkIterWithBradIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range iter.N(iterN) {
		}
	}

}

func BenchmarkIterWithMyFuncIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for range myIter(iterN) {
		}
	}
}

func BenchmarkIterWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ix := 0
		for ; ix < iterN; ix++ {
		}
	}
}

//goos: darwin
//goarch: amd64
//BenchmarkIterWithBradIter-8                75512             15541 ns/op
//BenchmarkIterWithMyFuncIter-8              71478             16026 ns/op
//BenchmarkIterWithLoop-8                    74888             15962 ns/op
//PASS
//ok      command-line-arguments  4.140s
