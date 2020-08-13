package util

import (
	"runtime"
	"sync"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:23
// @description
// @version
const concurrentLenBoundForMyQSort = 2e3

// MyConcurrentQSortWithChannelWait sorts the slice with quick sort by goroutines.
func MyConcurrentQSortWithChannelWait(a []int) {
	if a == nil {
		return
	}
	cQSortSyncWithChannel(a)
}

func cQSortSyncWithChannel(a []int) {
	mp := conquerSortSplitSlice(a)
	if mp == -1 {
		return
	}
	lSli := len(a)
	if lSli >= concurrentLenBoundForMyQSort {
		done := make(chan bool)
		go func() {
			cQSortSyncWithChannel(a[:mp])
			done <- true
		}()
		go func() {
			cQSortSyncWithChannel(a[mp+1:])
			done <- true
		}()
		<-done
		<-done
		return
	} else {
		qSortSeq(a[:mp])
		qSortSeq(a[mp+1:])
		return
	}
}

func MyConcurrentQSortWithWG(a []int) {
	if a == nil {
		return
	}
	cQSortWithWG(a)
}

func cQSortWithWG(a []int) {
	mp := conquerSortSplitSlice(a)
	if mp == -1 {
		return
	}
	lSli := len(a)
	if lSli >= concurrentLenBoundForMyQSort {
		wg := new(sync.WaitGroup)
		wg.Add(2)
		go func() {
			cQSortWithWG(a[:mp])
			wg.Done()
		}()
		go func() {
			cQSortWithWG(a[mp+1:])
			wg.Done()
		}()
		wg.Wait()
	} else {
		qSortSeq(a[:mp])
		qSortSeq(a[mp+1:])
		return
	}
}

type arrSEPair struct {
	start int
	end   int
}

func MyConcurrentQSortWithChannelTaskQueue(a []int) {
	if a == nil {
		return
	}
	wn := runtime.NumCPU()
	tq := make(chan arrSEPair, 5000) // 这里有可能爆栈的bug, 注意
	awg := new(sync.WaitGroup)
	for i := 0; i < wn; i++ {
		go func() {
			for pa := range tq {
				cLen := pa.end - pa.start
				if cLen < 2000 {
					qSortSeq(a[pa.start:pa.end])
				} else {
					mp := conquerSortSplitSlice(a[pa.start:pa.end])
					awg.Add(2)
					tq <- arrSEPair{pa.start, pa.start + mp}
					tq <- arrSEPair{pa.start + mp + 1, pa.end}
				}
				awg.Done()
			}
		}()
	}
	awg.Add(1)
	tq <- arrSEPair{0, len(a)}
	awg.Wait()
	close(tq)
}

// 分治,求中间值,有边界
func conquerSortSplitSlice(a []int) int {
	lSli := len(a)
	if lSli <= 1 {
		return -1
	} else if lSli == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return -1
	} else {
		endIncluded := lSli - 1
		// 原地快排
		middleIx := 0 // 首个元素
		fv := a[middleIx]
		a[middleIx], a[endIncluded] = a[endIncluded], fv
		mp := 0
		for i := 0; i < endIncluded; i++ {
			if a[i] < fv {
				if i != mp {
					a[mp], a[i] = a[i], a[mp]
				}
				mp++
			}
		}
		a[mp], a[endIncluded] = fv, a[mp]
		return mp
	}
}

// 串行快排
func qSortSeq(a []int) {
	mp := conquerSortSplitSlice(a)
	if mp == -1 {
		return
	}
	qSortSeq(a[:mp])
	qSortSeq(a[mp+1:])
}
