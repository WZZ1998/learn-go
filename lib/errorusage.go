package lib

import (
	"errors"
	"fmt"
	"time"

	perrors "github.com/pkg/errors"
	"golang.org/x/xerrors"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/29 21:10
// @description
// @version

func PlayWithError() {
	check(foobar)
}

type customArError struct {
}

func newCustomArError() error {
	return &customArError{}
}
func (e *customArError) Error() string {
	return "my own custom error"
}
func check(f func()) {
	defer func() {
		if errRecover := recover(); errRecover != nil {
			fmt.Print("[total recover in foobar]")
			fmt.Printf(" recovered: %+v \n", errRecover)
			err := errRecover.(error)
			fmt.Println(perrors.Cause(err))
		}
	}()
	f()
}
func foobar() {
	err := bar()
	if err != nil {
		fmt.Printf("%+v\n", err)
		if myErr := new(customArError); xerrors.Is(err, myErr) {
			fmt.Println("It is a custom error")
			var x *customArError
			xerrors.As(err, &x)
			fmt.Println(x)
		}
	}
}
func foobar2() {
	bar2()
}
func bar2() {
	barX, barY := _pRand.Intn(100), _pRand.Intn(60)
	defer func() {
		if errThrowing := recover(); errThrowing != nil {
			capErr := errThrowing.(error)
			panic(perrors.WithMessage(capErr,
				fmt.Sprintf("rec-pa at bar: X %d Y %d", barX, barY)))
		}
	}()
	foo2()

}
func foo2() {
	oriErr := getRError()
	if oriErr != nil {
		wrapped := perrors.Wrap(oriErr, "CAN'T HANDLE!")
		wrapped = perrors.WithMessage(wrapped, "foo2")
		panic(wrapped)
	}
}
func bar() error {
	err := foo()
	if err != nil {
		return xerrors.Errorf("bar : %w", err)
	}
	return err
}
func foo() error {
	originErr := getRError()
	return xerrors.Errorf("foo : %w", originErr)

}

func getRError() error {
	errByNew := errors.New("arbitrary error in func foo")
	errCustom := newCustomArError()
	errWithFmt := fmt.Errorf("temporary error created at %v", time.Now().Format(time.RFC3339))
	opts := [6]error{errByNew, errCustom, errWithFmt} // some nils
	rix := _pRand.Intn(len(opts))
	return opts[rix]
}
