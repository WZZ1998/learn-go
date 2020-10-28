package lib_test

import (
	"github.com/WZZ1998/learn-go/lib"
	"testing"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/27 22:45
// @description
// @version
func TestAbstracts(t *testing.T) {
	t.Log("test abstract fake implementation.")
	var x lib.FullABAble = lib.NewConcreteAbSum1(20, 30)
	t.Logf("interface call: A() %v B() %v Sum() %v", x.A(), x.B(), x.Sum())
	t.Log("abstract fake function OK.")
}
