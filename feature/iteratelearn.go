package feature

import (
	"fmt"
	"os"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/24 21:58
// @description
// @version
func LearnIterate() {
	problemSt()
	problemSt2()
	problemSt3()
	problemSt4()
	problemSt5()
	problemSt6()
	os.Exit(20)
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

func problemSt2() {
	rates := []int32{1, 2, 3, 4, 5, 6}
	var sink [6]int
	j := len(sink)
	for star := range rates {
		if star+1 < 1 {
			panic("")
		}
		j--
		sink[j] = j
	}

	fmt.Println(sink)
}
func problemSt3() {
	rs := []int{1, 2, 3}
	for ix := range rs {
		if ix+1 < 1 {
			panic("Impossible Panic???")
		}
	}
}

func problemSt4() {
	i := 5
top:
	i++
	if i < 6 {
		return
	}
	if i > 6 {
		return
	}
	println(i)
	goto top
}

func problemSt5() {
	i := 5
top:
	i++
	println(i)
	if !(i < 6 || i > 6) {
		goto top
	}

}

func problemSt6() {
	i := 0
	var sink [10]int
	j := len(sink)
top:
	j--
	sink[j] = j
	if i < 2 {
		i++
		if i < 1 {
			return
		}
		goto top
	}
}
