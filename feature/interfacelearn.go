package feature

import (
	"fmt"
	"math/rand"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/21 13:42
// @description
// @version
type shaper interface {
	area() float64
}

type square struct {
	shape
	a float64
}

type shape struct{}
type uls struct{}

func (sh *shape) area() float64 {
	// 一个抽象类,不过接口这么强大,抽象类貌似没啥亮点了
	// 本身OO现在也强调接口而不是抽象类的使用
	return 0.0
}

func (s *square) area() float64 {
	return s.a * s.a
}
func LearnInterface() {

	// var _ shaper = (*uls)(nil)
	//用编译器检验某个struct是否实现了某个接口,有IDE的话用处不大

	sq1 := &square{shape{}, 5}
	var areaFl shaper
	dynamicBindingInterface(&areaFl, sq1, &shape{})

	// 安全的类型断言,判断areaFl中是不是含有相应类型的变量
	if v, ok := areaFl.(*square); ok {
		// 注意,这里不能把类型写作.(square)
		// 接口做类型断言的时候,对类型的要求很严格,指针是指针,值类型是值类型,它们两个不相等
		// 注意,如果是值类型,那么会新建结构体的零值,如果结构体很大,这里可能就有点问题;而指针会设成nil
		fmt.Println("assert ok after assert: ", v)
		fmt.Println("areaFl area() :", areaFl.area())
	} else {
		fmt.Printf("assert ok: %t v value %v ", ok, v)
		// 转换之后,如果不是对应的类型,那么转换的值v会被置为相应目标类型(*square)的零值
		// 实际上这个时候areaFl包含的是一个shape类型的变量
		//fmt.Println("calling areaFl.area(): ", areaFl.area())
	}

	//type-switch 格式具有特殊的形式
	fmt.Println("type switch")
	switch t := areaFl.(type) {
	case *square:
		fmt.Printf("type *square value %v \n", t)
		//fallthrough 禁止在type-switch中使用fallthrough
	case *shape:
		fmt.Printf("type *shape value %v \n", t)
	case nil:
		fmt.Printf("type %T value %v \n", t, t)
		// 转换不成 t的类型还是保持为shaper,但是其实值是nil
	default:
		fmt.Printf("unknown type %T value %v \n", t, t)
	}
	if interV, ok := interface{}(sq1).(shaper); ok { // 不加interface{}是没办法编译通过的
		fmt.Printf("var sq1 implements interface shaper: %t\n", ok)
		fmt.Println("call with interV: ", interV.area())
	}

	//tsh := shape{}
	//var a2 shaper = tsh
	// 很有意思, 和之前一样,接口定义的方法绑定在结构体,receiver为值类型,而变量类型为指针,这样赋值可以
	// receiver 类型为指针类型,而变量为值类型的时候就不能通过编译
	// 原因: 考虑结构是某内置类型的别名, 这个结构的变量是可以被设置为常量的,这个时候怎么取地址?

}

func dynamicBindingInterface(inter *shaper, s1 *square, s2 *shape) {
	// 在程序运行之前,能知道areaFl绑定了哪一个实例和方法指针表吗?
	//显然是不能的,这只能在运行的时候,动态确定,因此interface是一个动态类型
	rr := rand.Int()
	fmt.Println("rr :", rr)
	if rr%3 == 0 {
		*inter = s2
	} else if rr%3 == 1 {
		*inter = s1
	} else {
		*inter = nil
	}
}
