package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/18 13:18
// @description
// @version

var gn *int

type UST struct { //类型名在包内唯一,大写首字母包外可见,小写首字母包外不可见
	name    string
	age     int
	balance int
	ratio   float64 //小写字母开头,都无法导出
}

func LearnStruct() {
	s0 := UST{"Jack", 40, 60000, 0.34}
	exSl := []UST{s0}       //注意这里对s0进行了值拷贝!
	exSl = append(exSl, s0) // 这里也进行了值拷贝!
	s0.name = "GreatJack"
	exSl[1].name = "StupidJack"
	fmt.Printf("exSl[0].Name == s0.Name: %t\n", s0.name == exSl[0].name)
	// 同理,map和数组也会进行类似的值拷贝
	for _, u := range exSl { // 值拷贝,不能修改切片内实例的属性
		u.name = "Richard"
	}
	fmt.Println(exSl) // 没有修改
	s1 := new(UST)
	s1.name = "John"
	s1.age = 12
	s1.balance = 220
	s1.ratio = 0.002

	var s2 UST
	s2.name = "Tom"

	s3 := UST{
		name:    "Tom",
		age:     33,
		balance: 57033,
		ratio:   0.2,
	}
	changeUSTName(&s3)
	opSomeUTS(100)
	us := genUSTs()
	ttt := 100
	// 注意!这里破例使用了内建的print系列函数,因为print打印变量不会使得这个变量发生逃逸,原因还有待进一步研究!
	println("s1 addr: ", s1, " s2 addr :", &s2, " s3 addr: ", &s3, " ttt addr: ", &ttt)
	for ix := range us {
		if ix%20000 == 0 {
			println("us[", ix, "] (UST addr) ", &us[ix])
		}
	}

	println("gn (addr of hx in opSomeUTS): ", gn)

	stp := int(1e7 / 5)
	var ir [1e7]int // 数组过大就会逃逸到堆内存
	for ix := range ir {
		if ix%stp == 0 {
			println("ir int array ", ix, " addr", &ir[ix])
		}
	}

}
func changeUSTName(u *UST) {
	u.name = "XXX"
}

func opSomeUTS(a int) {
	var n = a
	var hx = 10 // 逃逸到堆内存
	gn = &hx
	var x float64
	x = 3.88
	var sOnS UST
	println("in opsSomeUTS sOnS addr: ", &sOnS)
	sOnS.name = "strange1"
	sOnS.ratio = x
	sOnS.balance = n
}
func genUSTs() []*UST {
	var sl []*UST
	cnt := 80000
	for i := 0; i < cnt; i++ {
		sl = append(sl, new(UST))
	}
	return sl
}
