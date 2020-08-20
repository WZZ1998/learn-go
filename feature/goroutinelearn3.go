package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/8/20 10:25
// @description
// @version

func LearnGoroutine3() {
	leftMost := make(chan int)
	var left, right chan int = nil, leftMost
	for i := 0; i < 100000; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	fmt.Println("at most left:", <-leftMost)
}

func f(left, right chan int) {
	left <- (<-right) + 1
}
