package util_test

import (
	"sort"
	"testing"

	"github.com/WZZ1998/learn-go/util"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:29
// @description
// @version
var testSlLens = []int{1, 5, 100, 1e3, 1e4, 1e6, 1e7}

func TestMyConcurrentQSortWithChannelWait(t *testing.T) {
	t.Log("testing validation of my quick sort implementation")
	var testCases [][]int
	for _, ll := range testSlLens {
		td, err := util.GetRandIntSliceOfLength(ll)
		if err != nil {
			t.Fatal(err) // fatal会立刻终止,error会报错但是不会终止
		}
		testCases = append(testCases, td)
	}
	for ix, sl := range testCases {
		util.MyConcurrentQSortWithChannelWait(sl)
		verify(t, ix, sl)
	}
}
func TestMyConcurrentQSortWithWG(t *testing.T) {
	t.Log("testing validation of my quick sort implementation with wait group")
	var testCases [][]int
	for _, ll := range testSlLens {
		td, err := util.GetRandIntSliceOfLength(ll)
		if err != nil {
			t.Fatal(err) // fatal会立刻终止,error会报错但是不会终止
		}
		testCases = append(testCases, td)
	}
	for ix, sl := range testCases {
		util.MyConcurrentQSortWithWG(sl)
		verify(t, ix, sl)
	}
}
func TestMyConcurrentQSortWithChannelTaskQueue(t *testing.T) {
	t.Log("testing validation of my quick sort implementation with channel work queue")
	var testCases [][]int
	for _, ll := range testSlLens {
		td, err := util.GetRandIntSliceOfLength(ll)
		if err != nil {
			t.Fatal(err) // fatal会立刻终止,error会报错但是不会终止
		}
		testCases = append(testCases, td)
	}
	for ix, sl := range testCases {
		util.MyConcurrentQSortWithChannelTaskQueue(sl)
		verify(t, ix, sl)
	}
}
func verify(t *testing.T, caseNo int, sliceToVerify []int) {
	t.Helper()
	// 该方法能够标记某个测试方法是一个helper函数
	//当一个测试包在输出测试的文件和行号信息时，
	//将会输出调用help函数的调用者的信息，而不是输出helper函数的内部信息
	if !sort.IntsAreSorted(sliceToVerify) { // 没排序好
		t.Errorf("case %d: length %d sort invalid.", caseNo, len(sliceToVerify))
	}
}
