package feature

import (
	"fmt"
	"github.com/WZZ1998/learn-go/util"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/26 03:36
// @description
// @version
func LearnConcurrent() {
	const tCnt int = 5e5
	slBase, errGetOri := util.GetRandIntSliceOfLength(tCnt)
	if errGetOri != nil {
		fmt.Println("get origin data error:", errGetOri)
		return
	}
	slEx := make([]int, tCnt)
	copy(slEx, slBase)
	fmt.Println("origin data generated.")
	st := time.Now()
	sort.Ints(slBase)
	timeUsedBase := time.Since(st)
	st = time.Now()
	util.MyConcurrentQSortWithChannelWait(slEx)
	timeUsedEx := time.Since(st)
	fmt.Printf("base time %v ex time %v ratio: %.4f valid:%t\n",
		timeUsedBase,
		timeUsedEx,
		timeUsedEx.Seconds()/timeUsedBase.Seconds(),
		sort.IntsAreSorted(slEx))
	runAndPProf := false
	if runtime.GOOS == "darwin" && runAndPProf {
		fmt.Println("Run and generate pprof ana file.")
		var pCnt int = 1e7
		slRea, _ := util.GetRandIntSliceOfLength(pCnt)
		pth := "/Users/wangzizhou/Downloads/learn-go-profiles/mysqcpu.prof"
		cf, _ := os.OpenFile(pth,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			0644,
		)
		defer cf.Close()
		errStartProf := pprof.StartCPUProfile(cf)
		if errStartProf != nil {
			fmt.Println("start prof failed:", errStartProf)
			return
		}
		defer pprof.StopCPUProfile()
		util.MyConcurrentQSortWithChannelTaskQueue(slRea)

		//pth2 := "/Users/wangzizhou/Downloads/learn-go-profiles/mysqmem.prof"
		//mf, _ := os.OpenFile(pth2,
		//	os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
		//	0644,
		//)
		//defer mf.Close()
		//runtime.GC() // 手动触发GC
		//errWriteMem := pprof.WriteHeapProfile(mf)

		// 这个地方只能写入当前堆空间的信息,并不能记录分配的历史;
		//如果想看历史,可以使用go test或者http方法

		//if errWriteMem != nil {
		//	fmt.Println("write memory profile failed:", errWriteMem)
		//}
	}

}
