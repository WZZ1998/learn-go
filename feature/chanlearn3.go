package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/8/19 14:59
// @description
// @version

func LearnChan3() {
	// 这样写不会爆栈, 但是实际上golang的堆栈是可以扩展的
	fibonacciFunc := func(st1, st2 Any) (Any, Any, Any) {
		ist1, ist2 := st1.(int), st2.(int)
		sv := ist1 + ist2
		return ist1, ist2, sv
	}
	e := buildLazyIntEvaluator(fibonacciFunc, 1, 1)
	for i := 0; i < 5; i++ {
		fmt.Println("lazy generator generate:", e())
	}
}

type Any interface{}
type evalFunc func(Any, Any) (Any, Any, Any)

func buildLazyEvaluator(evf evalFunc, initState1, initState2 Any) func() Any {
	retVCh := make(chan Any)
	loopF := func() {
		st1 := initState1
		st2 := initState2
		var retV Any
		for {
			retV, st1, st2 = evf(st1, st2)
			retVCh <- retV
		}
	}
	evaluator := func() Any {
		return <-retVCh
	}
	go loopF()
	return evaluator
}

func buildLazyIntEvaluator(evf evalFunc, initState1, initState2 Any) func() int {
	evaluator := buildLazyEvaluator(evf, initState1, initState2)
	return func() int {
		return evaluator().(int)
	}
}
