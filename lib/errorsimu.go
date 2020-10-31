package lib

import (
	"github.com/pkg/errors"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/30 18:11
// @description
// @version

type Temporary interface { // 行为契约
	Temporary() bool
}

type sometimeFailError struct {
}

func (e *sometimeFailError) Error() string {
	return "sometime... It just failed"
}
func (e *sometimeFailError) Temporary() bool {
	return true
}
func FailSometime() (int, error) {
	r := _pRand.Intn(100)
	if r < 45 {
		return 0, errors.WithMessage(errors.WithStack(&sometimeFailError{}), "FailSomeTime")
	}
	return 19260817, nil
}
