package main

// @author  wzz_714105382@icloud.com
// @date  2020/5/31 01:51
// @description
// @version
import (
	"fmt"
	"learn-go/feature"
)

type F func()

func init() { // init函数会在包的最开始执行,这个函数不可以在代码中引用或者调用
	fmt.Println("initializing package main......")
}
func main() {
	//init() 不可以调用
	var mainFunc F = main
	//var initFunc F = init 无法引用到init
	//mainFunc() 引发死递归
	fmt.Printf("function main type v initFunc: %T\n", mainFunc)
	printSeparatingLine()
	feature.LearnString()
	printSeparatingLine()
	feature.LearnTypes()
	printSeparatingLine()
	feature.LearnSlice()
	printSeparatingLine()
	feature.LearnRandom()
	printSeparatingLine()
	feature.LearnTime()
	printSeparatingLine()
	feature.LearnPointer()
}

func printSeparatingLine() {
	fmt.Println("\n\n------------------------")
}
