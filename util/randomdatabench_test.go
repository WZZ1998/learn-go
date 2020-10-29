package util_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/WZZ1998/learn-go/util"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/19 23:23
// @description
// @version

func BenchmarkGetTimeUnixNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = time.Now().UnixNano()
	}
	// BenchmarkGetTimeUnixNano-8   	14115704	        81.7 ns/op
}

func BenchmarkGetRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = rand.Int()
	}
	//BenchmarkGetRandInt-8   	69230803	        16.8 ns/op
}
func BenchmarkGetRandIntWithLocalRand(b *testing.B) {
	sc := rand.NewSource(time.Now().UnixNano())
	localRa := rand.New(sc)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = localRa.Int()
	}
}

const benchRISliceLen = 1e7

func BenchmarkGetRandIntSliceOfLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.GetRandIntSliceOfLength(benchRISliceLen)
	}
}

func BenchmarkSimpleIRSliceOfLForBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = util.SimpleIRSliceOfLForBench(benchRISliceLen)
	}
}
