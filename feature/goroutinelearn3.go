package feature

import (
	"fmt"
	"time"
)

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

	// 流水线化处理
	var dataN = 20
	si := make(chan int, dataN)
	sOut := make(chan int, dataN)
	ci := make(chan int, dataN)
	cOut := make(chan int, dataN)
	st := time.Now()
	go seqProcessData(si, sOut)
	for i := 0; i < dataN; i++ {
		si <- i
	}
	for i := 0; i < dataN; i++ {
		pv := <-sOut
		if pv != i+4 {
			panic("wrong process res!")
		}
	}
	fmt.Println("seq process, time:", time.Since(st))
	close(si)

	stc := time.Now()
	go cProcessData(ci, cOut)
	for i := 0; i < dataN; i++ {
		ci <- i
	}
	for i := 0; i < dataN; i++ {
		pv := <-cOut
		if pv != i+4 {
			panic("wrong process res!")
		}
	}
	fmt.Println("pipeline process, time", time.Since(stc))
	close(ci)
}

func f(left, right chan int) {
	left <- (<-right) + 1
}

func preProcess(a int) int {
	time.Sleep(4 * time.Millisecond)
	return a + 1
}

func processA(a int) int {
	time.Sleep(8 * time.Millisecond)
	return a + 1
}

func processB(a int) int {
	time.Sleep(4 * time.Millisecond)
	return a + 1
}
func processC(a int) int {
	time.Sleep(7 * time.Millisecond)
	return a + 1
}

func cProcessWrap(processF func(int) int, in <-chan int, aOut chan<- int) {
	for d := range in {
		aOut <- processF(d)
	}
	close(aOut)
}

func seqProcessData(in <-chan int, out chan<- int) {
	for d := range in {
		tp := preProcess(d)
		ta := processA(tp)
		tb := processB(ta)
		tc := processC(tb)
		out <- tc
	}
	close(out)
}

func cProcessData(in <-chan int, out chan<- int) {
	preOut := make(chan int, 100)
	aOut := make(chan int, 100)
	bOut := make(chan int, 100)
	cOut := make(chan int, 100)
	go cProcessWrap(preProcess, in, preOut)
	go cProcessWrap(processA, preOut, aOut)
	go cProcessWrap(processB, aOut, bOut)
	go cProcessWrap(processC, bOut, cOut)
	go cProcessWrap(func(a int) int { return a }, cOut, out)
}
