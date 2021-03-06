package feature

import (
	"fmt"
	"sort"
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
	ori2 := ori[0:3:6] // 截取的cap,被append满了之后,会扩容并获得新的空间,不会使用原来的空间
	// 因为cap是在截取的时候定的,仅以这个设定的cap为准(即使原来的底层数组还有空间)
	// 如果截取的cap比原来的slice的cap大,那么直接panic
	fmt.Printf("ori2: %v\n", ori2)
	var des []int
	des = append(des, ori...)
	noo := append(ori, 9) // 扩容,换内存
	ori[0] = -9999
	fmt.Printf("ori %v des %v noo %v\n", ori, des, noo) // 因此,三者各不相同

	s1 := make([]int, 0, 1024)
	fmt.Println("s1 len before [:cap(s1)]:", len(s1))
	s1 = s1[:cap(s1)] // 直接增长
	// s1 = s1[:2*cap(s1)] // panic
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

	stBytes := []byte("") //  现在stBytes的cap是0了,历史上cap=32的情况被干掉了
	//fmt.Println(stBytes)  // 一句话让stBytes分配到堆
	fmt.Println("empty string to []byte cap:", cap(stBytes))
	ss1 := append(stBytes, 'a')
	ss2 := append(stBytes, 'b')
	print("s11 ss2 :")
	println(ss1, ss2)
	// ss1 和 ss2 指向不同的底层空间

	myArr1 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("myArr1:", myArr1)
	ssl1 := myArr1[0:2] // len 2 cap 6
	// ssl1 := myArr1[0:2:2]
	//len 2 cap 2 ,这样写下面就会panic,对切片进行切片,不能超过原来切片的cap范围!
	fmt.Println("ssl1 myArr1[0:2] :", ssl1)
	ssl2 := ssl1[0:6] // 可以直接切出来,即使ssl1的长度为2
	fmt.Println("ssl2 ssl1[0:6] :", ssl2)

	sToCut := []int{1, 2, 3}
	fmt.Println("sToCut:", sToCut)
	fmt.Println("sToCut[3:]:", sToCut[3:]) // 这样的切片可以,返回空切片,但是[4:]就不行了
	// 注意,冒号之前空默认值是0,冒号后默认值是原切片的len
	// 切片只能向后伸展,在底层数组大的情况下,
	// 切片: 首地址 长度 容量
	// re-slice 就是重新调整三个变量(以[切片头索引:切片尾索引:容量尾索引]为方法)
	// 必须满足: 1. 三个变量<=关系递增 2.三者 >= 0 而且 <= 原切片的cap
	sToCut2 := make([]int, 20, 100)
	fmt.Println("sToCut2 len cap:", len(sToCut2), cap(sToCut2))
	sCut := sToCut2[40:45:60] // 三个值都可以超过原来的len,但是都不能超过原来的cap
	fmt.Println("sToCut2[40:45:60] len cap:", len(sCut), cap(sCut))
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

func sliceFilter(origin []int, f func(int) bool) []int {
	var res []int
	for _, n := range origin {
		if f(n) {
			res = append(res, n)
		}
	}
	return res
}
