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
	td, err := getRandomIntSliceWithL(1e8)
	if err != nil {
		t.Fail()
	}
	utils.MyConcurrentQSort(td)
	if !sort.IntsAreSorted(td) {
		t.Fail()
	}
}

func BenchmarkMyConcurrentQSort(b *testing.B) {
	td, err := getRandomIntSliceWithL(1e7)
	if err != nil {
		b.Fail()
	}
	utils.MyConcurrentQSort(td)
	if !sort.IntsAreSorted(td) {
		b.Fail()
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
