package utils

import "sync"

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:23
// @description
// @version
const concurrentLenBoundForMyQSort = 2e3

// MyConcurrentQSort sorts the slice with quick sort by goroutines.
func MyConcurrentQSort(a []int) {
	if a == nil {
		return
	}
	cQSort(a)
}

func cQSort(a []int) {
	lSli := len(a)
	if lSli <= 1 {
		return
	} else if lSli == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return
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
		if lSli >= concurrentLenBoundForMyQSort {
			done := make(chan bool)
			go func() {
				cQSort(a[:mp])
				done <- true
			}()
			go func() {
				cQSort(a[mp+1:])
				done <- true
			}()
			<-done
			<-done
			return
		} else {
			cQSort(a[:mp])
			cQSort(a[mp+1:])
			return
		}
	}
}

func MyConcurrentQSortWithWG(a []int) {
	if a == nil {
		return
	}
	ll := len(a)
	cQSortWithWG(0, ll-1, a)
}

func cQSortWithWG(start, endIncluded int, a []int) {
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
		middleIx := start //+ (l / 2)
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
		if l >= concurrentLenBoundForMyQSort {
			wg := new(sync.WaitGroup)
			wg.Add(2)
			go func() {
				cQSortWithWG(start, splitPo-1, a)
				wg.Done()
			}()
			go func() {
				cQSortWithWG(splitPo+1, endIncluded, a)
				wg.Done()
			}()
			wg.Wait()
			return
		} else {
			cQSortWithWG(start, splitPo-1, a)
			cQSortWithWG(splitPo+1, endIncluded, a)
			return
		}
	}
}
