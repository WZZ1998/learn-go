package main

// @author  wzz_714105382@icloud.com
// @date  2020/5/31 01:51
// @description
// @version
import (
	"fmt"
	"learn-go/feature"
	"reflect"
	"runtime"
	"strings"
)

var sepL = func() string {
	fmt.Println("in main.go initializing var sepL")
	return strings.Repeat("\n", 2) + strings.Repeat("-", 40)
}()

type F func()

func init() { // init函数会在包的最开始执行,这个函数不可以在代码中引用或者调用
	fmt.Println("initializing package main......")
}
func main() {
	//init() 不可以调用
	var mainFunc F = main
	//var initFunc F = init 无法引用到init
	//mainFunc() 引发死递归
	fmt.Printf("function main type v initFunc: %T\n\n", mainFunc)
	learnFunctions := []func(){
		feature.LearnString,
		feature.LearnTypes,
		feature.LearnRandom,
		feature.LearnTime,
		feature.LearnPointer,
		feature.LearnControl,
		feature.LearnFunction,
		feature.LearnSlice,
		feature.LearnMap,
		feature.LearnLib,
		feature.LearnStruct,
		feature.LearnStruct2,
		feature.LearnMethod,
		feature.LearnMethod2,
		feature.LearnInterface,
		feature.LearnInterface2,
		feature.LearnReflect,
		feature.LearnReflect2,
		feature.LearnSimpleFP,
		//feature.LearnIO,
		//feature.LearnIO2,
		//feature.LearnIO3,
		feature.LearnConcurrent,
		feature.LearnJson,
		feature.LearnCrypto,
		feature.LearnException,
		feature.LearnException2,
		//feature.LearnSchedule, // 有for{} 除非trace,不要调用了
		feature.LearnMem,
		feature.LearnMem2,
		//feature.LearnMem3,
		feature.LearnGoroutine,
	}
	for _, lf := range learnFunctions {
		fmt.Println(getFuncNameWithFV(lf))
		fmt.Println()
		lf()
		printSeparatingLine()
	}

}

func printSeparatingLine() {
	fmt.Println(sepL)
}
func getFuncNameWithFV(fToGetName func()) string {
	vPtr := reflect.ValueOf(fToGetName).Pointer()
	pf := runtime.FuncForPC(vPtr)
	return pf.Name()
}
