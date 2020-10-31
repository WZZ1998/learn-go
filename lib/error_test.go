package lib_test

import (
	"github.com/WZZ1998/learn-go/lib"
	"github.com/pkg/errors"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/29 21:48
// @description
// @version
func TestPlayWithError(t *testing.T) {
	lib.PlayWithError()
}

func TestFailSometime(t *testing.T) {
	_, err := callWithRetry(3)
	if err != nil {
		t.Logf("error: %+v", err)
		ori := errors.Cause(err)
		t.Log("finding cause:", ori) //在errors的嵌套下,寻找到最本源的error
	}
}
func callWithRetry(cnt int) (int, error) {
	result, err := lib.FailSometime()
	if err == nil {
		return result, nil
	}
	if isTem(err) {
		for i := 0; i < cnt; i++ {
			result, err = lib.FailSometime()
			if err == nil {
				return result, nil
			}
		}
		return 0, errors.WithMessagef(err, "retried %d times", cnt)
	}
	return 0, err
}
func isTem(err error) bool {
	var t lib.Temporary
	return errors.As(err, &t) && t.Temporary()
}
