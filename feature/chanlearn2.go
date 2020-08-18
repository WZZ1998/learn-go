package feature

import (
	"fmt"
	"math/rand"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/15 01:54
// @description
// @version

func LearnChan2() {
	numCh := make(chan int)
	go func() {
		for i := 2; i < 10000; i++ {
			numCh <- i
		}
		close(numCh)
	}()
	max := 0
	var ch <-chan int = numCh
	for {
		pr, valid := <-ch
		if valid == false {
			break
		}
		max = pr
		ch = addPrimeRoutineFilter(ch, pr)
	}
	fmt.Println("max prime:", max)

	tk := time.NewTicker(100 * time.Millisecond)
	defer tk.Stop() // 记得要Stop()回收资源,否则会一直占用
	tm := time.NewTimer(300 * time.Millisecond)
	defer tm.Stop()
	lFlg := true
	for lFlg {
		select { // 如果有多个channel就绪,那么select会随机选择一个
		case tv := <-tk.C:
			fmt.Println(tv)
		case <-tm.C:
			lFlg = false
		}
	}
	//for {
	//	rCh := make(chan int, 1)
	//	go simulateCallTimeout(rCh)
	//	callTimer := time.NewTimer(10 * time.Millisecond)
	//	select {
	//	case res := <-rCh:
	//		_ = res
	//	case <-callTimer.C:
	//		fmt.Println("call failed. timeout!")
	//		panic("impossible timeout!")
	//	}
	//	callTimer.Stop()
	//}
}

func addPrimeRoutineFilter(in <-chan int, p int) <-chan int {
	ich := make(chan int)
	go func() {
		for v := range in {
			if v%p != 0 {
				ich <- v
			}
		}
		close(ich)
		//fmt.Println("prime ", p, "channel closed!")
	}()
	return ich
}

func timeCostCall(string) int {
	time.Sleep(20 * time.Microsecond)
	return rand.Intn(200)
}

func simulateCallTimeout(resCh chan<- int) {
	resCh <- timeCostCall("call!")
	close(resCh)
}
