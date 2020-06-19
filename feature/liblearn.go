package feature

import (
	"container/list"
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/17 17:02
// @description
// @version
func LearnLib() {
	linkL := list.New()
	linkL.PushFront(101)
	linkL.PushFront(100)
	linkL.PushFront(99)
	for p := linkL.Front(); p != nil; p = p.Next() {
		fmt.Printf("addr %p value %d\n", p, p.Value.(int)) // 双向链表,interface{}的问题开始暴露了,没泛型啊!
	}
	n := 100
	fmt.Printf("get mem size with unsafe : int : size %d\n", unsafe.Sizeof(n)) // 8个字节
	var sl []int
	fmt.Printf("get mem size with unsafe: slice []int : size %d\n", unsafe.Sizeof(sl)) // slice三个部分.24字节

	pat := "^[0-9]+..[0-9]+..[0-9]+$" // 正则表达式匹配:注意开头结尾^$匹配;注意编译pattern对象
	s1 := "2..4.555.333.......3"
	s2 := "23..44..12345"
	isS1Match, err1 := regexp.MatchString(pat, s1) //注意!这个方法内部还是编译了
	isS2Match, err2 := regexp.MatchString(pat, s2)
	if err1 != nil || err2 != nil {
		fmt.Println("reg match string error:", err1, err2)
		return
	}
	fmt.Printf("s1 match %t s2 match %t\n", isS1Match, isS2Match)
	re := regexp.MustCompile(pat)
	fmt.Printf("re match 23.444 %t\n", re.MatchString("23.444"))
	rem := regexp.MustCompile("[0-9]+\\.\\.") // 两次反义
	x := rem.FindAllString(s2, -1)
	fmt.Println(x)
	s2m1 := rem.ReplaceAllString(s2, "XX..")
	s2m2 := rem.ReplaceAllStringFunc(s2, func(s string) string {
		numP := s[:len(s)-2] //如果不是英文这里是不对的! 截取切片是按照字节进行的,而不是rune
		num, _ := strconv.Atoi(numP)
		num *= 2
		return strconv.Itoa(num) + "&&"
	})
	fmt.Printf("s2m1 :%s s2m2 : %s\n", s2m1, s2m2)

	// 大数运算包
	bi1 := big.NewInt(math.MaxInt64)
	bi2 := big.NewInt(4)
	bi1.Mul(bi1, bi2)
	bi3 := big.NewInt(math.MaxInt64)
	bbi := big.NewInt(1)
	bbi.Div(bi1, bi3)
	fmt.Println("bbi : ", bbi)
}
