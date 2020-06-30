package feature

import (
	"fmt"
	"learn-go/utils"
	"math/rand"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/26 03:36
// @description
// @version
func LearnConcurrent() {
	const tCnt int = 1e6
	slBase := make([]int, tCnt)
	slEx := make([]int, tCnt)
	var v int
	for i := 0; i < tCnt; i++ {
		v = rand.Int()
		slBase[i] = v
		slEx[i] = v
	}
	fmt.Println("origin data generated.")
	st := time.Now()
	//sort.Ints(slBase)
	//time.Sleep(2 * time.Second)
	timeUsedBase := time.Since(st)
	st = time.Now()
	utils.MyConcurrentQSort(slEx)
	timeUsedEx := time.Since(st)
	flg := true
	for i := 0; i < tCnt; i++ {
		if slEx[i] != slBase[i] {
			flg = false
			break
		}
	}
	fmt.Printf("base time %v ex time %v ratio: %.4f valid:%t\n",
		timeUsedBase,
		timeUsedEx,
		timeUsedEx.Seconds()/timeUsedBase.Seconds(),
		flg)

}
