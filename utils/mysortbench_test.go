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
const slToSortL int = 1000000

var td, modSl []int

func setup() error {
	fmt.Println("set up for myQSort benchmark")
	var err error
	td, err = getRandomIntSliceWithL(slToSortL)
	if err != nil {
		return err
	}
	modSl = make([]int, len(td))
	return nil
}
func TestMain(m *testing.M) {
	fmt.Println("从这里开始的?")
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
