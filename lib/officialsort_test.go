package lib_test

import (
	"sort"
	"testing"

	"github.com/WZZ1998/learn-go/lib"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/26 00:44
// @description
// @version

func TestSortBlas(t *testing.T) {
	blas := []*lib.BlaDO{{1, 16}, {2, 15}, {0, 1}}
	lib.SortBlas(blas)
	t.Logf("sorted: %v", blas)
	blaXs := make([]int, 0, len(blas))
	for _, bla := range blas {
		blaXs = append(blaXs, bla.X)
	}
	if !sort.IntsAreSorted(blaXs) {
		t.Fatal("sort blas failed.")
	}
}
