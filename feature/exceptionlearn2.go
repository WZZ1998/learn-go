package feature

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 12:40
// @description
// @version
func LearnException2() {
	errorHandler(myF)(0, 0)
	testNilPerformance()
	//runOuterCommand()
}
func runOuterCommand() {
	llCommand := exec.Command("top")
	err2 := llCommand.Run() // run的时候,当前goroutine会wait这个进程结束
	//如果进程被kill,那么err2会包含  signal: killed
	// 这里需要手动kill
	if err2 != nil {
		fmt.Println("err2 while run llCommand:", err2)
	}
}
func testNilPerformance() {
	var myNil interface{}
	st := time.Now()
	for i := 0; i < 1e7; i++ {
		t1(myNil)
	}
	fmt.Println("!= nil then type assert-t1:", time.Since(st))
	st = time.Now()
	for i := 0; i < 1e7; i++ {
		t2(myNil)
	}
	fmt.Println("directly type assert-t2:   ", time.Since(st))
	// 肯定是有差距,但是貌似不大
	// != nil then type assert-t1: 27.619883ms
	//directly type assert-t2:    40.377385ms

	// 网上有文章说, 作为参数传递然后调用,具体类型最快,空接口转换至具体类型次之,直接传递接口又次之,空接口转换至接口最慢
	// 不过权衡很重要, 性能数据也可能随着优化发生变化,有待考证

	// 另外,使用指针实现接口性能通常比较好
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
