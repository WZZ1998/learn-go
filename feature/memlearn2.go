package feature

import (
	"fmt"
)

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
	showInterfaceArg(&ce) // 传入指针,创建接口类型,不做值拷贝
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
	foobar(myC)

	var inx interface{} = myC // 发生了值拷贝!如果myC是个很大的对象,问题就很严重
	println("direct convert:", interface{}(myC))
	// 即使是这样直接转换,还是会做值拷贝
	print("inx:")
	println(inx)
	print("type assert addr:")
	con, _ := inx.(int) // 这里自然也是值拷贝
	println(&con)
	con = 1
	print("myCf:")
	println(myC)
	fmt.Printf("myC: %d\n", myC)

	// 千万注意值拷贝问题,不要直接把值赋给接口
	// 调用方法,如果receiver是值类型,那么也会进行值拷贝

	escF(ST{t: 20200714, name: "esc"}) // 这个地方逃逸了吗?
	// 函数调用的字面量实参,确实在栈里存在了,栈中还有堆中形参的地址

}
func escF(st ST) *ST {
	ttt := 0
	print("ttt addr:")
	println(&ttt)
	print("st addr:")
	println(&st)
	//p := uintptr(unsafe.Pointer(&ttt))
	//dumpMemoryWithUnsafe(p, p+64, 8)

	//ttt addr:0xc0002bbca8
	//st addr:0xc00000d060
	//addr v [c0002bbcd0] deref int        18717215 hexi 11d9a1f
	//addr v [c0002bbcc8] deref int        20200714 hexi 1343d0a  这里! !
	//addr v [c0002bbcc0] deref int        18301630 hexi 11742be
	//addr v [c0002bbcb8] deref int    824636587528 hexi c0002bbe08
	//addr v [c0002bbcb0] deref int    824633774176 hexi c00000d060
	//addr v [c0002bbca8] deref int               0 hexi 0
	return &st
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
	concrete, _ := in.(ST) // 这里类型断言,也做了结构体拷贝! // concrete并不是原来创建的那个值了
	// concrete本身就是个值类型,值类型赋给值类型,自然是值拷贝
	concrete.name = "XXX"
	print("concrete addr:")
	println(&concrete)
}

type NT struct {
	i   interface{}
	idp *int
}

func foobar(i interface{}) *NT {
	print("in foobar():")
	println(i)
	ntp := &NT{}
	ntp.i = i
	//这逃逸分析没有想象中的厉害呀!
	// 154行 ntp其实并没有逃逸,函数结束ntp就成垃圾了
	// 但是由于155行赋值,编译器还是把i的实际数据放到堆了
	// 可能开启inline之后能优化掉吧

	x := 1000 // 这种情况下,x 必须分配到堆了
	// 注意,如果NT的属性是值类型,那么拷贝的时候可能原始的变量并不需要分配到堆
	println("x addr:", &x)
	ntp.idp = &x
	// 同样的, 这里的x本身也可以分配在栈上,因为ntp函数结束即成为垃圾
	// 但是目前的go 1.14 的逃逸分析 darwin环境,还是把x放到了堆上
	return nil
}
