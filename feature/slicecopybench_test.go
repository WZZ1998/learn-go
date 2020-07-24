package feature_test

import (
	"learn-go/util"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/14 17:41
// @description
// @version
const oLen = 1024 * 8

func BenchmarkCopySlice(b *testing.B) {
	b.Log("test original slice len:", oLen)
	ori, errGetOri := util.GetRandIntSliceOfLength(oLen)
	if errGetOri != nil {
		b.Fatal("Get origin data failed:", errGetOri)
	}
	dstLoo := make([]int, 0, oLen)
	dstAppend := make([]int, 0, oLen)

	dstCopy := make([]int, oLen) // 先把长度开出来,然后往里面copy
	dstOneToOne := make([]int, oLen)

	b.Run("BenchmarkCopySliceWithLoop",
		func(fb *testing.B) {
			for i := 0; i < fb.N; i++ {
				for j := 0; j < oLen; j++ {
					dstLoo = append(dstLoo, ori[j])
				}
				dstLoo = dstLoo[:0]
			}
		})
	b.Run("BenchmarkCopySliceWithDirectAppend",
		func(fb *testing.B) {
			for i := 0; i < fb.N; i++ {
				dstAppend = append(dstAppend, ori...) // append 打散
				dstAppend = dstAppend[:0]             // 大约0.8ns /op darwin go 1.14
			}
		})

	b.Run("BenchmarkCopySliceWithCopy",
		func(fb *testing.B) {
			for i := 0; i < fb.N; i++ {
				copy(dstCopy, ori) // dstCopy的len必须和ori对应
			}
		})

	b.Run("BenchmarkCopySliceWithOneToOne",
		func(fb *testing.B) {
			for i := 0; i < fb.N; i++ {
				for j := 0; j < oLen; j++ {
					dstOneToOne[j] = ori[j]
				}
			}
		})
}

// 结论: 切片很小时,copy稍有优势;随着切片长度增加,copy和打散append的性能差距基本可以忽略不计
// 这还没有考虑,上面重置切片的时间消耗;因此,打散append和copy,各有各的用途,不用顾忌用哪一个

// 测试环境 darwin go1.14 intel i7-6820HQ MaBook Pro Late 2016
//BenchmarkCopySlice
//    BenchmarkCopySlice: slicecopybench_test.go:15: test original slice len: 64
//BenchmarkCopySlice/BenchmarkCopySliceWithLoop
//BenchmarkCopySlice/BenchmarkCopySliceWithLoop-8         	11169400	       103 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithDirectAppend
//BenchmarkCopySlice/BenchmarkCopySliceWithDirectAppend-8 	84635210	        13.9 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithCopy
//BenchmarkCopySlice/BenchmarkCopySliceWithCopy-8         	90794641	        11.8 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithOneToOne
//BenchmarkCopySlice/BenchmarkCopySliceWithOneToOne-8     	11774038	        99.8 ns/op

//BenchmarkCopySlice
//    BenchmarkCopySlice: slicecopybench_test.go:15: test original slice len: 8192
//BenchmarkCopySlice/BenchmarkCopySliceWithLoop
//BenchmarkCopySlice/BenchmarkCopySliceWithLoop-8         	   77013	     15235 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithDirectAppend
//BenchmarkCopySlice/BenchmarkCopySliceWithDirectAppend-8 	  599180	      1792 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithCopy
//BenchmarkCopySlice/BenchmarkCopySliceWithCopy-8         	  621340	      1810 ns/op
//BenchmarkCopySlice/BenchmarkCopySliceWithOneToOne
//BenchmarkCopySlice/BenchmarkCopySliceWithOneToOne-8     	  111670	     10564 ns/op
