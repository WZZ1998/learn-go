package feature

import "fmt"

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
