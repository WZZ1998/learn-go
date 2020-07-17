package feature

import (
	"fmt"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/16 17:55
// @description
// @version
func LearnMem3() {
	fmt.Println("array:")
	sort1() //数组
	fmt.Println("slice:")
	sort2() // 切片
	//sort3() // 堆逃逸的切片
}

func sort3() {
	start := time.Now().UnixNano()
	var sli []int
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		sli = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		bubbleSort(sli)
	}
	fmt.Println(time.Now().UnixNano() - start)
}
func sort2() {
	start := time.Now().UnixNano()
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		bubbleSort(sli)
	}
	fmt.Println(time.Now().UnixNano() - start)
}
func sort1() {
	start := time.Now().UnixNano()
	const NUM int = 100000000
	for i := 0; i < NUM; i++ {
		base := [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		bubbleSort(base[:9])
	}
	fmt.Println(time.Now().UnixNano() - start)
	// 9 7102903000
	// 10 5611907000
	// 11 5442214000
	// 12 5574813000
	// 13 5025680000
	// 14 5206263000
	// 15 4989521000
	// 16 5213174000
	// 17 5034452000
	// 18 5375320000
	//打印消耗时间
}

//冒泡
func bubbleSort(arr []int) {
	//println(&arr)
	//st := uintptr(unsafe.Pointer(&arr))
	//dumpMemoryWithUnsafe(st, st+4096, 8)
	for j := 0; j < len(arr)-1; j++ {
		for k := 0; k < len(arr)-1-j; k++ {
			if arr[k] < arr[k+1] {
				arr[k], arr[k+1] = arr[k+1], arr[k]
			}
		}
	}
}
