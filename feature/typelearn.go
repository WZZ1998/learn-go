package feature

import (
	"fmt"
	"os"
	"runtime"
	"unicode"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/1 22:46
// @description golang name and type
// @version
type MI int32
type Rope string

var mii MI = 100

const (
	ccc = -1
	//veryBig int64 = 9999999999999999999 溢出的常量将无法通过编译
	veryPrecise float64 = 0.1234567898765432123456 // 精度丢失
)

func LearnTypes() {
	//_ = 1
	//fmt.Println(_)
	//underscore 不能被使用

	fmt.Printf("const v type : %T ccc = %v\n", ccc, ccc)
	//fmt.Printf("const v type : %T veryBig = %v\n", veryBig, veryBig)
	fmt.Printf("const v type %T veryPrecise = %v\n", veryPrecise, veryPrecise)
	var it int32 = 100
	//mii = it 无法通过编译,golang遵守静态类型&强类型,可以用强制类型转换
	fmt.Printf("MI(int) type v mii is :%T\n", mii)
	fmt.Printf("int32 type v it is :%T\n", it)
	fmt.Println("casting:")
	a := 0.6666666

	fmt.Printf("var a type : %T a = %v\n", a, a)
	fmt.Printf("保留三位小数:%.3f\n", a)
	aCast := float32(a) // 精度丢失
	fmt.Printf("aCast = %v\n", aCast)
	b := 2359
	fmt.Printf("二进制表示:%b "+
		"八进制:%o 十进制:%d 十六进制:%x, Unicode字符:%c\n",
		b, b, b, b, b)
	var cStart int32 = 47000
	var ll int32 = 25
	fmt.Printf("start unicode num = %d, length = %d\n", cStart, ll)
	for i := cStart; i <= cStart+25; i++ {
		fmt.Printf("%c ", i)
	}
	fmt.Println()
	goos := runtime.GOOS
	fmt.Printf("The OS is %s\n", goos)
	path := os.Getenv("GOPATH")
	fmt.Printf("GO Path is %s\n", path)

	p, q := 12, 56
	fmt.Printf("p = %d q = % d then swap \n", p, q)
	p, q = q, p
	fmt.Printf("p = %d, q = %d now \n", p, q)

	j := 8712
	fmt.Printf("the number is %v with 5 numbers %05d\n ", j, j)
	// 规定输出的最小长度

	t := 51324.3344594
	fmt.Printf("the number is %v, with precision format %8.6f\n", t, t)
	// m.ng m表示最小宽度(没有数字用空格补齐), n表示保留小数的位数;
	// 若是%g则m代表最小宽度(同样用空格补齐),n代表有效数字位数

	cx := 3.535 + 5.22i
	cx = cx * cx
	cxr := real(cx)
	cxi := imag(cx)
	fmt.Printf("cx type : %T value = %v\n", cx, cx)
	fmt.Printf("cx type : %T .4f value = %.4f\n", cx, cx)
	fmt.Printf("cxr cxi type : %T %T value : %v %v\n", cxr, cxi, cxr, cxi)
	var nu int = 127
	fmt.Printf("nu type %T binary value: %b\n", nu, nu)
	fmt.Printf("^ with int nu %v bit value %b\n", ^nu, ^nu)
	//这个value的二进制表示很有意思,绕开了位数的信息
	//但是^还是按位取反,只不过是按照补码的表示按位取反

	var myRope Rope = "Rope is a alias for string"
	fmt.Println(myRope)

	ch := '\u0534'
	fmt.Printf("ch unicode point:\n"+
		"type %T unicode %U value %v char %c\n", ch, ch, ch, ch)
	fmt.Printf("\\u0534 isLetter: %t isDigit:%t isSpace:%t",
		unicode.IsLetter(ch), unicode.IsDigit(ch), unicode.IsSpace(ch))
}
