package feature

import (
	"fmt"
	"runtime"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/20 22:53
// @description
// @version
const (
	MonL  = "Monday"
	TueL  = "Tuesday"
	WedL  = "Wednesday"
	ThurL = "Thursday"
	FriL  = "Friday"
	SatL  = "Saturday"
	SunL  = "Sunday"
)

var ix2LiteralWeekDay = [7]string{SunL, MonL, TueL, WedL, ThurL, FriL, SatL}

// 并不存在所谓数组常量或者切片常量
const (
	sunday day = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

type A struct {
	r int
	s int
}

func (receiver *A) String() string {
	return fmt.Sprintf("~~~ r = %d s = %d ", receiver.r, receiver.s)
}

func LearnMethod2() {
	md := tuesday
	fmt.Printf("md: %v\n", md)
	ma := A{1, 2}
	fmt.Printf("ma: %v\n", ma)
	// 非常有意思,当String()方法定义在指针上,而给出的是值类型实例时,并不会调用String()方法
	// 可能是因为,当使用类型别名时,如果类型支持常量, 事实上无法对类型进行取地址操作

	//runtime.GC() 触发GC

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d kb \n", m.Alloc/(1<<10))
}

type day int

func (receiver *day) String() string {
	return ix2LiteralWeekDay[*receiver]
}
