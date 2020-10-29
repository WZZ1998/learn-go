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
	testLens := []int{1, 3, 10, 40, 100, 1e3, 1e6, 1e7}
	for caseNo, testL := range testLens {
		res, err := util.GetRandIntSliceOfLength(testL)
		if err != nil {
			t.Errorf("case: %d slice len %d error: %v", caseNo, testL, err)
			continue
		}
		if len(res) != testL {
			t.Errorf("case: %d target len %d != res len %d", caseNo, testL, len(res))
			continue
		}
	}
}
