package feature

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/4 02:01
// @description
// @version

func LearnSlice() {
	//var exBigArr [1 << 44]int 无法分配这么大的数组,会OOM
	// 千万注意,数组是一种值类
	var arr2dim [5][8]int
	fmt.Println("2 dim array ", arr2dim)
	ar0 := new([3]int)
	fmt.Printf("ar0: type %T value %v \n", ar0, ar0)
	ar1 := [...]int{2: 1, 3: 2, 4: 3} // 指定索引初始化
	ar2 := ar1                        //注意,这一步发生了数组拷贝,因为数组是值类型;注意将数组直接作为参数也会导致数组拷贝
	ar2[0] = 9
	ar2[1] = 8
	fmt.Printf("ar0 %v addr %p ar1 %v addr %p ar2 %v addr %p\n", *ar0, ar0, ar1, &ar1, ar2, &ar2)
	// ar1 ar2 修改互不干扰

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

	s1 := make([]int, 0, 1024)
	fmt.Println("s1 len before [:cap(s1)]:", len(s1))
	s1 = s1[:cap(s1)]
	// s1 = s1[:2*cap(s1)] panic
	fmt.Println("s1 len after [:cap(s1)]:", len(s1))

	s2 := make([][]int, 2)
	for ix := range s2 {
		s2[ix] = make([]int, 3)
	}
	fmt.Printf("构造多维切片:s2 type %T value %v\n", s2, s2)

	bys := []byte("冰清玉洁")
	sss := string(bys)
	fmt.Printf("bys %v sss %s\n", bys, sss)
	var poemCh rune
	for _, poemCh = range sss { // 注意,这里承接range的可以是外部的变量
		fmt.Printf("%U ", poemCh)
	}
	fmt.Println()

	nsLatter := []int{1, 2, 3}
	nsFormer := make([]int, 8)
	ns1 := make([]int, len(nsFormer)+len(nsLatter))
	copy(ns1, nsFormer)
	copy(ns1[len(nsFormer):], nsLatter)
	ns2 := append(nsFormer, nsLatter...)
	fmt.Println("ns1(copy): ", ns1, "ns2(append):", ns2) // 使用copy实现的拼接,肯定不如append方便快捷
	chs := []byte{'a', 'b', 'c', 'd', 'e'}
	fmt.Println("chs: ", chs)
	fmt.Println("重叠拷贝:")
	cCnt := copy(chs[2:5], chs[0:3])
	fmt.Println("chs: ", chs) // 注意,copy做的是这种整段的拷贝,而不是一个值一个值的拷贝
	fmt.Println("elements copied:", cCnt)
	// 看起来真的是整段拷贝,可能做了特殊处理, 或者直接用了额外的空间来腾挪拷贝
	originS := []int{12, 45, 33, 44, 2, 1, 0, 888}
	fmt.Printf("筛选slice元素 偶数: %v \n", sliceFilter(originS, func(i int) bool { return i%2 == 0 }))

	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	copy(s3[3:], s3[5:]) //移除了索引为3和4的元素
	//for ii, ll := 8, len(s3); ii < ll; ii++ { 加上这个可以把之前写过的值清零 有无必要呢?
	//	s3[ii] = 0
	//}
	s3 = s3[0:8]
	s3 = s3[:cap(s3)] // 但是,这里如果做切片重组,会发现原来的值还是在的
	// 如果默认从底层数组扩展出来的新空间初始化为0的话,这就比较麻烦
	// 不知道这个算不算是个需要考虑的问题
	fmt.Println(s3)

	unsortedNums := []int{12, 49, 230, 2349, 394, 9, 129, 133, 1, 2}
	sort.Ints(unsortedNums) //切片排序
	fmt.Printf("unsortedNums now : %v, isSorted: %t\n", unsortedNums, sort.IntsAreSorted(unsortedNums))
	fmt.Printf("index of 394 : %d\n", sort.SearchInts(unsortedNums, 394))
	// 注意,search做二分查找,切片必须是排好序的!必须排好序! 如果没有查到,就返回slice的长度
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

func sliceFilter(origin []int, f func(int) bool) []int {
	var res []int
	for _, n := range origin {
		if f(n) {
			res = append(res, n)
		}
	}
	return res
}
