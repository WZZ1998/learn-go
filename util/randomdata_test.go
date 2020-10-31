package util_test

import (
	"testing"

	"github.com/WZZ1998/learn-go/util"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/20 00:09
// @description
// @version

func TestGetRandIntSliceOfLength(t *testing.T) {
	tests := []int{1, 3, 10, 40, 100, 1e3, 1e6, 1e7}
	for caseNo, tt := range tests {
		ranSlice, err := util.GetRandIntSliceOfLength(tt)
		if err != nil {
			t.Errorf("case %d : slice len %d get error: %v", caseNo, tt, err)
			continue
		}
		if len(ranSlice) != tt {
			t.Errorf("case %d : want len %d != ranSlice len %d", caseNo, tt, len(ranSlice))
			continue
		}
	}
}
