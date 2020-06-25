package feature

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/23 01:48
// @description
// @version
func LearnIO() {
	var li2 string
	var li3 string
	n, inputErr := fmt.Scan(&li2, &li3)
	// scan,当遇到换行符号,执行扫描, 按照个数进行匹配,将换行视为空格,直到匹配全所有的目标
	fmt.Printf("ele cnt: %d, err: %v content:%s\n", n, inputErr, li2)

	var li string // 直接扫入一行
	fmt.Println("Please enter a line")
	n, inputErr = fmt.Scanln(&li)
	// 机制和scan基本一致,区别扫描过程中遇到换行符,会直接结束扫描,如果不够,
	// 或者多了(在最后一个扫上的目标之后还有目标), 会有相应的error
	fmt.Printf("ele cnt: %d, err: %v, content:%s\n", n, inputErr, li)

	var fName, lName string // 扫入两个变量
	fmt.Println("Please enter your name")
	n, inputErr = fmt.Scanf("%s\n  %s\n", &fName, &lName)
	// scanf 不会把换行符视为空格,会在键入换行符之后开始扫描;
	// 多余的连续空格符会被视为一个空格
	// %s 后面除了空格和换行符,什么都不要跟;这样的字符是不可能被扫描到的:千万不要写%se这种,后面的e是扫不到的

	// 匹配规则: %s从第一个非空格字符开始匹配,如果是换行直接不匹配;否则写入变量到遇到第一个换行或者空格停止
	// 模式串空格使得模式串指针移到下一个非空格字符,匹配串指针移动到下一个非空格字符
	// 模式串换行先使得匹配串指针移动到下一个非空格字符,然后做严格匹配
	// 其余做严格一对一匹配
	fmt.Printf("ele cnt: %d, err: %v\n", n, inputErr)
	fmt.Println("welcome", fName, lName)

	inputReader := bufio.NewReader(os.Stdin)
	i2 := bufio.NewReader(os.Stdin)
	fmt.Println("enter something:")
	str, err := inputReader.ReadString('\n')
	nStr := strings.TrimRight(str, "\n") // 注意, 读进来的字符串是带着换行符的
	fmt.Printf("str : %s err %v \n", nStr, err)
	str, err = i2.ReadString('\n')
	str = strings.TrimRight(str, "\n")
	fmt.Printf("str : %s err %v \n", str, err)
}
