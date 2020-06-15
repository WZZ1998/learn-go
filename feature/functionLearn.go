package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/15 01:26
// @description
// @version

//func ef(x,y int) int 如果这个函数能够在runtime中被外部定义/实现,
//golang允许这样的申明比如time.Now()内部包裹的now()

func LearnFunction() {
	fmt.Println("getZero return:", getZero(1e6))
	r1, r2 := get2Zero()
	fmt.Println("get2Zero return:", r1, r2)
	fmt.Println("getZeroV2 return:", getZeroV2())
	ar := []int{1, 2, 3}
	arr1, arr2, arr3, arr4 := 5, 6, 7, 8
	fmt.Println("call with slice")
	fmt.Printf("ar = %v address = %p\n", ar, &ar)
	f1(-1, ar...) // 注意:这里是对slice直接进行的值拷贝(slice的本质是增强指针)
	fmt.Printf("ar = %v address = %p\n", ar, &ar)

	fmt.Println("call with arguments")
	fmt.Printf("arr1 %d arr2 %d arr3 %d arr4 %d\n", arr1, arr2, arr3, arr4)
	f1(-1, arr1, arr2, arr3, arr4) // 这里进行的是参数的值拷贝,不会影响原来的值
	fmt.Printf("arr1 %d arr2 %d arr3 %d arr4 %d\n", arr1, arr2, arr3, arr4)

	df1()

}
func getZero(int) int {
	// 函数参数可以不具名
	return 0
}
func get2Zero() (rx1, rx2 int) {
	// 返回值可以具名
	rx1 = 0
	rx2 = 0
	return
}
func getZeroV2() (rr1 int) { // 具名返回值必须加括号
	rr1 = 0
	return
}

func f1(x int, args ...int) {
	fmt.Printf("x : %d args %v address %p \n", x, args, &args)
	args[0] = 1000
	fmt.Printf("in function f1: args %v address %p\n", args, &args)
	y := x * x
	f2(y, args...) //在内部实际上args直接被看做了[]int,要想继续作为...int类型形参的实际参数,还得args...打散
}
func f2(y int, args ...int) {
	return
}

func df1() {
	ti := 10000
	defer ddf1(ti)
	for i := 0; i <= 4; i++ {
		defer ddf1(i) // 注意,注册的defer会倒序执行(先注册的后执行),就像弹栈一样
	}
	ti = -1

}
func ddf1(i int) {
	fmt.Println("in defer function ddf1! i : ", i)
}
