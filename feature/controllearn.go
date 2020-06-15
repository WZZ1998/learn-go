package feature

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/12 10:32
// @description learn control flow golang
// @version

func LearnControl() {
	x := 100
	if x = -1; x > 0 {
		fmt.Println("x is positive")
	} // 可以在条件处进行初始化或者复赋值,但是在此处声明的变量会遮蔽外部的同名变量!
	intS := "ABC"
	in, err := strconv.Atoi(intS)
	if err != nil {
		fmt.Printf("program error %v\n", err)
		//os.Exit(1)
	} else {
		fmt.Println("convert success. integer:", in)
	}
	nn := 100
	switch nn { // 注意! golang的switch不需要写break!
	case 100, 200, 300:
		fmt.Printf("hundred <= 300!\n")
		//fallthrough 使用fallthrough会导致滑落下一个case分支
	case 400, 500, 600:
		fmt.Printf("hundred <= 600 >= 400!\n")
	default:
		fmt.Printf("not recognizable hundred!")
	}
	switch {
	case alwaysFalse():
		fmt.Println("it is not me! always false!")
	case alwaysTrue():
		fmt.Println("it is me! always true!")
	case 100 < 400: // 依然具有顺序性
		fmt.Println("it is NOT me! 200 < 400")
	default:
		fmt.Println("I am default case!")
	}

	switch rrr := getRandInt200(); {
	case rrr <= 50:
		fmt.Println("rrr <= 50!")
	case rrr > 50 && rrr <= 100:
		fmt.Println(" 50 < rrr <= 100!")
	default:
		fmt.Println("rrr > 100")
	}
	for i, j := 0, 5; i < j; i, j = i+1, j-1 { // 利用平行赋值同时使用多个循环计数器
		fmt.Printf("i = %d, j = %d\n", i, j)
	}
	var sbForMatrix strings.Builder
	var middleSymbol string
	for i := 1; i <= 10; i++ { // 打印一个20*10 星号矩阵
		switch i {
		case 1, 10:
			middleSymbol = "*"
		default:
			middleSymbol = " "
		}
		sbForMatrix.WriteString("*")
		sbForMatrix.WriteString(strings.Repeat(middleSymbol, 18))
		sbForMatrix.WriteString("*")
		fmt.Println(sbForMatrix.String())
		sbForMatrix.Reset()
	}
	iii := 5
	for iii > 0 { //类似while循环
		fmt.Printf("%d~", iii)
		iii--
	}
	fmt.Println()
	names := []string{"Jack", "Nick", "Tom"}
	for ix, name := range names {
		fmt.Printf("No.%d ! %s!\n", ix, name) // 注意,for range循环对原来集合的元素进行值拷贝
	}
	pL := "锦瑟无端五十弦,一弦一柱思华年"
	for ix, c := range pL {
		fmt.Printf("%c(%d)^", c, ix) // 千万注意,这个index是字节数组的index,也就是每个rune的字节数组位置
	}
	fmt.Println()

}
func alwaysTrue() bool  { return true }
func alwaysFalse() bool { return false }
func getRandInt200() int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(200)
}
