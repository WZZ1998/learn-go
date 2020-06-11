package feature

import (
	"fmt"
	"strings"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/11 19:45
// @description
// @version

var xG int = 50000

func LearnPointer() {
	x := 1000
	fmt.Printf("x type %T value %v address %p\n", x, x, &x)
	fmt.Printf("xG type %T value %v address %p\n", xG, xG, &xG)
	s1 := `"pointer is complex! Right?"`
	//s1 = append(s1, "Appending something!") 字符串是value类型,不许再append了
	s2 := s1
	s1p := &s1
	fmt.Printf("s1 value %s s1p value %p &s1 value %p s2 value %s &s2 value %p\n", s1, s1p, &s1, s2, &s2)
	*s1p = strings.Repeat("No!", 15000)
	fmt.Printf("s1 value too long! s1!=s2:%t s1p value %p &s1 value %p s2 value %s &s2 value %p\n",
		s1 != s2, s1p, &s1, s2, &s2)
	/*
		说明: golang中的字符串本质上是一个指向内存空间的slice,应该是包含了地址,长度,容量信息
		对字符串进行的=拷贝,进行的是slice信息的值拷贝,但是因为字符串是不可变类型,所以没什么影响
		s1 和 s2 简单理解,就是栈空间上两个大小固定的结构体
		期初他们指向同一段内存空间,修改时候各自指向自己的内存空间
		结构体大小是固定的,因此,s1和s2在栈空间上的位置,是不会变的
	*/
	//var inP *int = nil
	//*inP = 1 对空指针解引用会导致panic
}
