package feature_test

import "testing"
import "github.com/bradfitz/iter"

// @author  wzz_714105382@icloud.com
// @date  2020/7/24 16:56
// @description
// @version
const iterN = 50000

func BenchmarkIterWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < iterN; j++ {
			_ = j
		}
	}
}

func BenchmarkIterWithBradIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for ix := range iter.N(iterN) {
			_ = ix
		}
	}

}
