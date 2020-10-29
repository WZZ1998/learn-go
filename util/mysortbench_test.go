package util_test

import (
	"testing"

	"github.com/WZZ1998/learn-go/util"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 21:58
// @description
// @version
const slToSortL int = 1e6

func BenchmarkAllMySort(b *testing.B) {
	benchOriginData, errGetData := util.GetRandIntSliceOfLength(slToSortL)
	if errGetData != nil {
		b.Fatal("prepare data failed, error:", errGetData)
	}
	modSl := make([]int, len(benchOriginData))
	b.Run("BenchmarkMyConcurrentQSortWithChannelWait", func(subB *testing.B) {
		for i := 0; i < subB.N; i++ {
			subB.StopTimer() // 把这个时间给扣出去
			copy(modSl, benchOriginData)
			subB.StartTimer()
			util.MyConcurrentQSortWithChannelWait(modSl)
		}
	})
	b.Run("BenchmarkMyConcurrentQSortWithWG", func(subB *testing.B) {
		for i := 0; i < subB.N; i++ {
			subB.StopTimer() // 把这个时间给扣出去
			copy(modSl, benchOriginData)
			subB.StartTimer()
			util.MyConcurrentQSortWithWG(modSl)
		}
	})
	//b.Run("BenchmarkStdLibQSort", func(subB *testing.B) {
	//	for i := 0; i < subB.N; i++ {
	//		subB.StopTimer()
	//		copy(modSl, benchOriginData)
	//		subB.StartTimer()
	//		sort.Ints(modSl)
	//	}
	//})
	b.Run("BenchmarkMyConcurrentQSortWithChannelTaskQueue", func(subB *testing.B) {
		for i := 0; i < subB.N; i++ {
			subB.StopTimer()
			copy(modSl, benchOriginData)
			subB.StartTimer()
			util.MyConcurrentQSortWithChannelTaskQueue(modSl)
		}
	})

}

// 跑来跑去还是用channel的快一点,大概有10%的性能优势
//分析发现,waitGroup同步器分配到了堆上
// 但是按理说,channel的共享内存也应该是分配到堆,分配到栈上怎么实现共享data呢?
//(只是猜测,也可能就是有这种操作)
//或者可能golang对channel的管理有所优化
//BenchmarkMyConcurrentQSort
//BenchmarkMyConcurrentQSort-8                  18         320727575 ns/op
//BenchmarkMyConcurrentQSortWithWG
//BenchmarkMyConcurrentQSortWithWG-8            13         363510666 ns/op
