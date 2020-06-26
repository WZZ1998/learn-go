package feature

import (
	"fmt"
	"math/rand"
	"sort"
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
	sort.Ints(slBase)
	//time.Sleep(2 * time.Second)
	timeUsedBase := time.Since(st)
	st = time.Now()
	myConcurrentQSort(slEx)
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
func myConcurrentQSort(a []int) {
	ll := len(a)
	CQSort(0, ll-1, a)
}

func CQSort(start, endIncluded int, a []int) {
	l := endIncluded - start + 1
	if l <= 1 {
		return
	} else if l == 2 {
		if a[start] > a[endIncluded] {
			a[start], a[endIncluded] = a[endIncluded], a[start]
		}
		return
	} else {
		// 原地快排
		middleIx := start + (l / 2)
		fv := a[middleIx]
		a[middleIx], a[endIncluded] = a[endIncluded], a[middleIx]
		mp := start
		for i := start; i <= endIncluded-1; i++ {
			if a[i] < fv {
				if i != mp {
					a[mp], a[i] = a[i], a[mp]
				}
				mp++
			}
		}
		a[mp], a[endIncluded] = a[endIncluded], a[mp]
		splitPo := mp
		if l >= 1e4 {
			c1 := make(chan bool)
			c2 := make(chan bool)

			go func() {
				CQSort(start, splitPo-1, a)
				c1 <- true
			}()
			go func() {
				CQSort(splitPo+1, endIncluded, a)
				c2 <- true
			}()
			<-c1
			<-c2
			return
		} else {
			CQSort(start, splitPo-1, a)
			CQSort(splitPo+1, endIncluded, a)
			return
		}
	}
}
