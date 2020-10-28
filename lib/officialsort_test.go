package lib_test

import (
	"github.com/WZZ1998/learn-go/lib"
	"sort"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/26 00:44
// @description
// @version

func TestSortBlas(t *testing.T) {
	blas := []*lib.BlaDO{{1, 16}, {2, 15}, {0, 1}}
	lib.SortBlas(blas)
	t.Logf("sorted: %v", blas)
	xs := make([]int, 0, len(blas))
	for _, bla := range blas {
		xs = append(xs, bla.X)
	}
	if !sort.IntsAreSorted(xs) {
		t.Fatal("sort blas failed.")
	}
}
