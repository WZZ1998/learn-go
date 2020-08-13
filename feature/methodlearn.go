package feature

import (
	"fmt"
	"strconv"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/19 15:51
// @description
// @version
type dataS struct {
	a int
	b int
}

func (receiver *dataS) addThem() int {
	return receiver.a + receiver.b
}
func (receiver *dataS) clearThem() {
	receiver.a = 0
	receiver.b = 0
}

type intVector struct {
	n  int
	sl []int
}

//func (receiver intVector) add(e int) {
//这样定义的方法,对receiver做的是值拷贝,并不能修改到原来的内容
//	receiver = append(receiver, e)
//}
func (receiver *intVector) add(e int) {
	receiver.sl = append(receiver.sl, e)
}
func (receiver intVector) show() {
	fmt.Printf("show receiver addr %p\n", &receiver)
	receiver.n = -1
	receiver.sl = append(receiver.sl, -1)
	fmt.Println("receiver:", receiver)
}
func LearnMethod() {
	d1 := &dataS{3, 6}
	d2 := dataS{1, 2}
	d1.clearThem()
	d2.clearThem() // 虽然method的receiver是指针,但是好像用值类型直接调用也可以
	fmt.Println("d1 :", d1)
	fmt.Println("d2 :", d2)
	fmt.Println("d1.addThem: ", d1.addThem())
	fmt.Println("d2 add Them: ", d2.addThem())
	nums := intVector{100, []int{1, 2, 3}}
	nums.add(4) // 调用时,自动取地址,隐含的
	nums.add(5)
	s := (*intVector).add
	s(&nums, 6) // 真有意思,拆成func expression就不给做自动取地址了
	fmt.Println("nums :", nums)
	numsP := &nums
	fmt.Printf("unmsp :%p\n", numsP)
	fmt.Printf("unmsp value :%#v \n", numsP)
	// 很有意思,如果receiver是指针,那么即使你用值类型调用,依然会给你取地址,能够修改原来的实例
	// 如果receiver是值类型,那么即使用指针调用,还是会解引用,做值传递,不能修改修改原来的实例
	// 因此,方法的行为,受控于receiver的类型!
	fmt.Println("play mercedes:")
	myMercedes := &mercedes{car{4, &fusionEngine{1746}}}
	// interface 的本质是指针
	myMercedes.start()
	myMercedes.sayHi()
	myMercedes.stop()

	ee := &employee{
		person{
			base{101193},
			"Jack",
			"Ma",
		},
		170000.3,
	}
	fmt.Println("ee id:", ee.Id())

	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()

	mi := new(myInt)
	*mi = 5
	fmt.Println("mi:", mi.String())

}

type engine interface {
	start()
	stop()
}
type car struct {
	wheelCnt uint
	engine
}

func (receiver *car) numberOfWheels() uint {
	return receiver.wheelCnt
}

type mercedes struct {
	car
}

func (receiver *mercedes) sayHi() {
	fmt.Printf("Hi from mercedes at %p !\n", receiver)
}

type fusionEngine struct {
	id int
}

func (receiver *fusionEngine) start() {
	fmt.Println("fusion engine start!")
}
func (receiver *fusionEngine) stop() {
	fmt.Println("fusion engine stop!")
}

type base struct {
	id int
}

func (receiver *base) Id() int {
	return receiver.id
}
func (receiver *base) setId(id int) {
	receiver.id = id
}

type person struct {
	base
	firstName string
	lastName  string
}
type employee struct {
	person
	salary float64
}

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (receiver Base) MoreMagic() {
	receiver.Magic()
	receiver.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type myInt int

func (receiver *myInt) String() string {
	return "myInt:" + strconv.Itoa(int(*receiver))
}
