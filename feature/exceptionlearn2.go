package feature

import (
	"fmt"
	"math/rand"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 12:40
// @description
// @version
func LearnException2() {
	errorHandler(myF)(0, 0)
	testNilPerformance()
}
func testNilPerformance() {
	var myNil interface{}
	st := time.Now()
	for i := 0; i < 1e7; i++ {
		t1(myNil)
	}
	fmt.Println("t1:", time.Since(st))
	st = time.Now()
	for i := 0; i < 1e7; i++ {
		t2(myNil)
	}
	fmt.Println("t2:", time.Since(st))
}
func errorHandler(f func(int, int)) func(int, int) {
	// 很遗憾,函数并不会根据参数的类型协变,所以每个signature的函数都得定制一个errorHandler,导致不是太方便
	// 具体方不方便,还得看实际应用的状况
	return func(a, b int) {
		defer func() {
			fmt.Println("handler recovering...")
			// 无论怎么样,都会走一下这个函数,进行recover,检查是否包含错误,然后恢复
			// 影响性能? 尽量在高层调用recover? 如何设计?
			err := recover()
			// 这东西应该还是有意义的,毕竟空接口到接口得检查方法,性能比较下面做了测试,性能差了大约2.5倍
			if err != nil {
				if errT, ok := err.(error); ok { // 做了类型断言,转化成error
					fmt.Println("panic error captured:", errT)
				}
			}
		}()
		f(a, b)
	}
}
func myF(int, int) {
	var err error
	for i := 0; i < 5; i++ {
		_, err = deliberateErrorF()
		check(err)
	}

}
func deliberateErrorF() (int, error) {
	n := rand.Int()
	if n%17 == 0 {
		return 0, fmt.Errorf(" %d mod 17 == 0! deliberate error", n)
	}
	return n, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func t1(s interface{}) {
	if s != nil {
		if _, ok := s.(error); ok {
			fmt.Println("impossible error")
		}
	}

}
func t2(s interface{}) {
	if _, ok := s.(error); ok {
		fmt.Println("impossible error")
	}
}
