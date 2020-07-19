package utils_test

import (
	"learn-go/utils"
	"math/rand"
	"testing"
	"time"
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

const benchRISliceLen = 1e7

func BenchmarkGetRandIntSliceOfLength(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = utils.GetRandIntSliceOfLength(benchRISliceLen)
	}
}

func BenchmarkSimpleIRSliceOfLForBench(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = utils.SimpleIRSliceOfLForBench(benchRISliceLen)
	}
}
