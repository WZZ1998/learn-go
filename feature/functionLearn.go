package feature

import (
	"fmt"
	"log"
	"strings"
)

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

	fmt.Println("内置的new函数:") //返回的是指针
	v := 3
	vv := new(int)
	*vv = 4
	fmt.Printf("v %d addr %p vv %d addr %p\n", v, &v, *vv, vv)

	res := calFibonacciTable(100)
	n := 26
	fmt.Printf("fibonachi %d :%d\n", n, res[n])
	fmt.Printf("cal fibonachi with lambda %d : %d\n", n, calFibonacciWithLambda(n))

	rawS := "杨柳岸dex@@@ql&&&,晓1aa风daw;daw;d残!!!!月"
	ss := strings.Map(keepUCharAndComma, rawS) //如果rune在map函数返回了负值,那么新的字符串将会丢弃这个字符
	fmt.Printf("map去除杂乱符号: rawS: %s ss: %s\n", rawS, ss)

	func(a, b int) int { // 直接调用匿名函数
		res := a + b
		fmt.Printf("a = %d b = %d res = %d\n", a, b, res)
		return res
	}(3, 4)

	// 闭包
	fa := adder()
	fmt.Printf("first +100:%d  ", fa(100))
	fmt.Printf("second +100:%d  ", fa(100))
	fmt.Printf("third +100:%d\n", fa(100))

	log.SetFlags(log.Llongfile)
	where := log.Print
	where("Here!")

}
func keepUCharAndComma(rc rune) rune {
	if rc == ',' || rc > 255 {
		return rc
	}
	return -1
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
	defer myUnTrace(myTrace("df1")) // 注意,这里会先调用trace
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
func myTrace(s string) string {
	fmt.Println("entering :", s)
	return s
}
func myUnTrace(s string) {
	fmt.Println("leaving :", s)
}

func calFibonacciTable(ma int) []int {
	t := make([]int, 0, ma+1)
	t = append(t, 0, 1, 1)
	for i := 3; i <= ma; i++ {
		t = append(t, t[i-1]+t[i-2])
	}
	return t
}

func adder() func(int) int {
	var x int
	var y int
	y = 0
	defer fmt.Printf("leaving adder :x addr %p y addr %p\n", &x, &y)
	return func(delta int) int {
		x += delta
		return x
	}
}

func calFibonacciWithLambda(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	p, q := 1, 1
	cf := func() { p, q = q, p+q }
	cnt := n - 2
	for i := 0; i < cnt; i++ {
		cf()
	}
	return q
}
