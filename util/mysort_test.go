package util_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/WZZ1998/learn-go/util"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:29
// @description
// @version
var _testEleLens = []int{1, 5, 100, 1e3, 1e4, 1e6, 1e7}

func TestMyConcurrentQSortWithChannelWait(t *testing.T) {
	t.Log("testing validation of my quick sort implementation")
	tests, genErr := generateRandomSlice(_testEleLens)
	if genErr != nil {
		t.Fatal(genErr)
	}
	for ix, tt := range tests {
		util.MyConcurrentQSortWithChannelWait(tt)
		verify(t, ix, tt)
	}
}
func TestMyConcurrentQSortWithWG(t *testing.T) {
	t.Log("testing validation of my quick sort implementation with wait group")
	tests, genErr := generateRandomSlice(_testEleLens)
	if genErr != nil {
		t.Fatal(genErr)
	}
	for ix, tt := range tests {
		util.MyConcurrentQSortWithWG(tt)
		verify(t, ix, tt)
	}
}
func TestMyConcurrentQSortWithChannelTaskQueue(t *testing.T) {
	t.Log("testing validation of my quick sort implementation with channel work queue")
	tests, genErr := generateRandomSlice(_testEleLens)
	if genErr != nil {
		t.Fatal(genErr)
	}
	for ix, tt := range tests {
		util.MyConcurrentQSortWithChannelTaskQueue(tt)
		verify(t, ix, tt)
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

func generateRandomSlice(eleLens []int) ([][]int, error) {
	var res [][]int
	for _, l := range eleLens {
		d, err := util.GetRandIntSliceOfLength(l)
		if err != nil {
			return nil, fmt.Errorf("generateRandomSlice : %w", err)
		}
		res = append(res, d)
	}
	return res, nil
}
