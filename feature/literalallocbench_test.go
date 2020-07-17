package feature_test

import (
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/17 12:02
// @description
// @version
const benchUniDirPth = "/Users/wangzizhou/Downloads/learn-go-profiles/"
const doCPUProfile = false

func BenchmarkLiteralAlloc(b *testing.B) {
	if runtime.GOOS == "darwin" && doCPUProfile {
		b.Log("profile CPU to", benchUniDirPth)
		pf, err1 := os.Create(path.Join(benchUniDirPth, "literalCPUProfile.prof"))
		if err1 != nil {
			b.Fatal("open bench profile failed:", err1)
		}
		defer pf.Close()

		errProf := pprof.StartCPUProfile(pf)
		if errProf != nil {
			b.Fatal("start profile failed:", errProf)
		}
		defer pprof.StopCPUProfile()
	}
	b.Run("BenchmarkLiteralAllocWithArray",
		func(fb *testing.B) {
			for i := 0; i < fb.N; i++ {
				ar := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //, 9, 10, 11, 12, 13, 14, 15, 16}
				bubbleSort(ar[:8])
			}
		})
	//b.Run("BenchmarkLiteralAllocWithSlice",
	//	func(fb *testing.B) {
	//		for i := 0; i < fb.N; i++ {
	//			sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //, 9, 10, 11, 12, 13, 14, 15, 16}
	//			bubbleSort(sl)
	//		}
	//	})

}

func bubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
}
