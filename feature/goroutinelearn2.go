package feature

import (
	"fmt"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/10 21:51
// @description
// @version
func LearnGoroutine2() {
	myCh1 := make(chan bool)
	go checkGetThenCheck(myCh1)
	//myCh1 <- true
	//myCh1 <- true
	for i := 0; i < 3; i++ {
		myCh1 <- true
		time.Sleep(100 * time.Millisecond) // 等一下
	}
	close(myCh1)
	time.Sleep(500 * time.Millisecond)

}

func checkGetThenCheck(ch <-chan bool) {
	for {
		var x bool
		x = <-ch
		fmt.Println("received:", x)
		if x == false {
			break
		}
	}
	// 在通道关闭之后,接收方从通道只能拿到对应类型的零值,并且会立即返回
	fmt.Println("After break from loop, get :", <-ch)
	// 注意,向已经关闭的channel发送,或者关闭一个已经关闭的channel,都会导致panic
	//可以包装一层,在内部recover一下
}
