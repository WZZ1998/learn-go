package feature

// @author  wzz_714105382@icloud.com
// @date  2020/7/24 21:58
// @description
// @version
func LearnIterate() {
	problemSt()
	rs := []int{1, 2, 3}
	for ix := range rs {
		if ix+1 < 1 {
			panic("Impossible Panic???")
		}
	}
}

func problemSt() {
	i := 0
top:
	if i < 2 {
		i++
		if i < 1 {
			return
		}
		println(i)
		goto top
	}
}
