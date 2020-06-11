package feature

import (
	"fmt"
	"math/rand"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/4 02:01
// @description
// @version

func LearnSlice() {
	// 关键点:slice和slice之间没有任何关系,赋值做值拷贝,但是可能共用底层数组
	fmt.Println("modify with value & modify with pointer:")
	var arr [10]int
	fmt.Println("array now: ", arr)
	a := arr[0:0]
	fmt.Printf("a value %v address %p\n", a, &a)
	fmt.Println("array now: ", arr)
	modifyWithV(a) // 在这之后:数组的数据变了吗? 变了 a作为一个slice,却是没有任何变化,slice是互相独立的!!!
	fmt.Printf("a value %v address %p\n", a, &a)
	fmt.Println("array now: ", arr)
	modifyWithP(&a)
	fmt.Printf("a value %v address %p\n", a, &a)
	fmt.Println("array now: ", arr)

	fmt.Println("切片拷贝:")
	//arr2 := [...]int{91,92,93} 注意了!三个点是让编译器自己推算数据长度,但是arr2还是数组
	//arr2 = append(arr2, 94) 不允许对数组进行append
	ori := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ori2 := ori[0:3:6]
	fmt.Printf("ori2: %v\n", ori2)
	var des []int
	des = append(des, ori...)
	noo := append(ori, 9) // 扩容,换内存
	ori[0] = -9999
	fmt.Printf("ori %v des %v noo %v\n", ori, des, noo) // 因此,三者各不相同

	fmt.Println("copy slice: append v.s. copy:")
	var sr []int
	srLen := 100000
	testTimes := 1000
	rand.Seed(int64(int(time.Now().Nanosecond())))
	for i := 0; i < srLen; i++ {
		sr = append(sr, rand.Int())
	}
	fmt.Println("HEATING:test copy with copy, useless time :", testCopy(sr, copyWithCopy, testTimes))
	fmt.Println("HEATING:test copy with append, useless time :", testCopy(sr, copyWithAppend, testTimes))
	fmt.Printf("testTimes %d srLen %d date&time: %v\n", testTimes, srLen, time.Now().Format(time.RFC822))
	fmt.Println("test copy with append, time :", testCopy(sr, copyWithAppend, testTimes))
	fmt.Println("test copy with copy, time :", testCopy(sr, copyWithCopy, testTimes))
	//testTimes 1000 srLen 128 date&time: 11 Jun 20 23:56 CST
	//test copy with append, time : 841.734µs
	//test copy with copy, time : 1.328477ms

	//testTimes 1000 srLen 100000 date&time: 11 Jun 20 23:58 CST
	//test copy with append, time : 138.704941ms
	//test copy with copy, time : 97.470128ms

	// 数组长度小的时候,append快,较大的时候copy快
	//可能和append的扩容机制有关

}

func modifyWithV(s []int) []int {
	fmt.Println("modify with V")
	s = append(s, 1)
	return s
}
func modifyWithP(s *[]int) {
	fmt.Println("modify with P")
	*s = append(*s, 777) // 注意这里写入了之前已经写入过的位置
}

func copyWithCopy(sr []int) []int {
	var de []int
	de = append(de, sr...)
	return de
}

func copyWithAppend(sr []int) []int {
	de := make([]int, len(sr))
	copy(de, sr)
	return de
}
func testCopy(sr []int, fff func([]int) []int, testTimes int) time.Duration {
	st := time.Now()
	for i := 0; i < testTimes; i++ {
		_ = fff(sr)
	}
	t := time.Since(st)
	return t
}
