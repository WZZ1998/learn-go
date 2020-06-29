package utils

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 16:23
// @description
// @version

// MyConcurrentQSort sorts the slice with quick sort by goroutines.
func MyConcurrentQSort(a []int) {
	ll := len(a)
	cQSort(0, ll-1, a)
}

func cQSort(start, endIncluded int, a []int) {
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
		if l >= 1<<16 {
			c1 := make(chan bool)
			c2 := make(chan bool)

			go func() {
				cQSort(start, splitPo-1, a)
				c1 <- true
			}()
			go func() {
				cQSort(splitPo+1, endIncluded, a)
				c2 <- true
			}()
			<-c1
			<-c2
			return
		} else {
			cQSort(start, splitPo-1, a)
			cQSort(splitPo+1, endIncluded, a)
			return
		}
	}
}
