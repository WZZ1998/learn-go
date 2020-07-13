package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/7/11 01:18
// @description
// @version

type emp struct{}

type ST struct {
	t    int
	name string
}

func (sp ST) foo() string { // 绑定到值了
	fmt.Printf("in foo sp addr %p \n", &sp)
	return "bar ST!"
}

func (sp *ST) setName(n string) {
	fmt.Printf("in method: sp %p \n", sp)
	sp.name = n
}

type myI interface {
	foo() string
}

func LearnMem2() {

	myS := ST{name: "origin"}
	fmt.Printf("myS ST value addr %p \n", &myS)
	myS.setName("peter")
	// 绑定到指针的方法,也可以用结构体的值类型进行调用,但是并不会导致拷贝结构体,而是进行自动的取指针
	// 所以拷贝什么关键还是看形式参数
	fmt.Println("after set name to peter:")
	fmt.Println("now myS", myS)

	fmt.Println("escape and empty struct analysis")
	e1 := &emp{}
	e2 := &emp{}
	// 使用fmt print 两个都逃逸到堆内存,占用空间为零,使用公共zero base
	fmt.Printf("e1 addr : %p, e2 addr: %p, e1 addr == e2 addr %t \n", e1, e2, e1 == e2)

	x := 0            // 没有逃逸
	var s myI = &ST{} // 逃逸到堆
	//绑定到值定义的方法,实现了接口,可以使用具体类型的指针赋值给接口
	// 但是接口和值对调不行
	s.foo() // 调用接口方法,会导致被分配到堆
	// 注意这里发生结构体拷贝
	print("(on heap) s:")
	println(s)
	print("(on stack) %x:")
	println(&x)

	bar() // 把地址泄露出去,分配到堆
	ipp := &ip
	fmt.Println("ip addr:", ipp)

	ce := ST{} // 值类型
	print("ce addr:")
	println(&ce)
	showInterfaceArg(&ce)
	fmt.Println()

	var myI2 myI = ce // 呵呵,这里居然发生了结构体拷贝!
	print("myI2:")
	println(myI2)
	showInterfaceArg(myI2)
	fmt.Println()

	showInterfaceArg(ce) //第一步,先转成接口,需要做结构体拷贝!

	// 空接口呢?
	myC := 2
	print("myC addr:")
	println(&myC)
	var inx interface{} = myC // 发生了值拷贝!如果myC是个很大的对象,问题就很严重
	print(inx)
	println(inx)

	// 千万注意值拷贝问题,不要直接把值赋给接口
	// 调用方法,如果receiver是值类型,那么也会进行值拷贝

}

var ip *int // 必须得在堆上

func recordIntAddr(ad *int) {
	ip = ad
}

func bar() {
	myX := 1
	recordIntAddr(&myX)
}

func showInterfaceArg(in myI) {
	println("in showInterfaceArg:")
	print("arg interface:")
	println(in)
	in.foo()
	concrete, _ := in.(ST) // 这里类型断言,也做了结构体拷贝,concrete并不是原来的那个值了
	concrete.name = "XXX"
	print("concrete addr:")
	println(&concrete)
}
