package utils_test

import (
	"fmt"
	"learn-go/utils"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:29
// @description
// @version
func TestMyConcurrentQSort(t *testing.T) {
	t.Log("testing validation of my quick sort implementation")
	var testSlLens = []int{1, 5, 100, 1e3, 1e4, 1e6, 1e7}
	var testCases [][]int
	for _, ll := range testSlLens {
		td, err := getRandomIntSliceWithL(ll)
		if err != nil {
			t.Fatal(err) // fatal会立刻终止,error会报错但是不会终止
		}
		testCases = append(testCases, td)
	}
	for ix, sl := range testCases {
		utils.MyConcurrentQSort(sl)
		verify(t, ix, sl)
	}
}
func TestMyConcurrentQSortWithWG(t *testing.T) {
	t.Log("testing validation of my quick sort implementation with wait group")
	var testSlLens = []int{1, 5, 100, 1e3, 1e4, 1e6, 1e7}
	var testCases [][]int
	for _, ll := range testSlLens {
		td, err := getRandomIntSliceWithL(ll)
		if err != nil {
			t.Fatal(err) // fatal会立刻终止,error会报错但是不会终止
		}
		testCases = append(testCases, td)
	}
	for ix, sl := range testCases {
		utils.MyConcurrentQSortWithWG(sl)
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
func getRandomIntSliceWithL(lengthOfS int) ([]int, error) {
	if lengthOfS > 1e9 {
		return nil, fmt.Errorf("length of int slice too big")
	}
	td := make([]int, lengthOfS)
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < lengthOfS; i++ {
		td[i] = rand.Int()
	}
	return td, nil
}
