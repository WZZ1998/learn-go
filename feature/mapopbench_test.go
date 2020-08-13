package feature_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/26 17:25
// @description
// @version
func BenchmarkMapOpContentInsert(b *testing.B) {
	b.Run("WithFixedByteSliceToString", func(sb *testing.B) {
		for i := 0; i < sb.N; i++ {
			mm := make(map[string]int64, 1000000)
			kBuf := make([]byte, 0, 32)
			var j int64
			for j = 0; j < 1000000; j++ {
				v := time.Now().Unix()
				kBuf = strconv.AppendInt(kBuf, j, 10)
				kBuf = append(kBuf, '_')
				kBuf = strconv.AppendInt(kBuf, v, 10)
				mm[string(kBuf)] = v
				kBuf = kBuf[:0]
			}
		}
	})

	b.Run("WithNewSliceAndUnsafe", func(sb *testing.B) {
		for i := 0; i < sb.N; i++ {
			mm := make(map[string]int64, 1000001)
			var j int64
			for j = 0; j < 1000000; j++ {
				kBuf := make([]byte, 0, 32)
				v := time.Now().Unix()
				kBuf = strconv.AppendInt(kBuf, j, 10)
				kBuf = append(kBuf, '_')
				kBuf = strconv.AppendInt(kBuf, v, 10)
				k := *(*string)(unsafe.Pointer(&kBuf))
				mm[k] = v
			}
		}
	})
	b.Run("WithSPrintf", func(sb *testing.B) {
		for i := 0; i < sb.N; i++ {
			mm := make(map[string]int64, 1000001)
			var j int64
			for j = 0; j < 1000000; j++ {
				v := time.Now().Unix()
				k := fmt.Sprintf("%d_%d", j, v)
				mm[k] = v
			}
		}
	})

}
