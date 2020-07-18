package utils_test

import (
	"fmt"
	"learn-go/utils"
	"os"
	"sort"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 21:58
// @description
// @version
const slToSortL int = 10000000

var td, modSl []int

func setup() error {
	var err error
	td, err = getRandomIntSliceWithL(slToSortL)
	if err != nil {
		return err
	}
	modSl = make([]int, len(td))
	return nil
}
func TestMain(m *testing.M) {
	fmt.Println("utils_test main started...")
	errSetUp := setup()
	if errSetUp != nil {
		fmt.Println("error while setup:", errSetUp)
		return
	}
	co := m.Run()
	os.Exit(co)
}
func BenchmarkMyConcurrentQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // 把这个时间给扣出去
		copy(modSl, td)
		b.StartTimer()
		utils.MyConcurrentQSort(modSl)
		// wait group
		// channel
	}
}
func BenchmarkMyConcurrentQSortWithWG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer() // 把这个时间给扣出去
		copy(modSl, td)
		b.StartTimer()
		utils.MyConcurrentQSortWithWG(modSl)
	}
}
func BenchmarkStdLibQSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		copy(modSl, td)
		b.StartTimer()
		sort.Ints(modSl)
	}
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
