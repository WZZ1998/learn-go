package main

// @author  wzz_714105382@icloud.com
// @date  2020/5/31 01:51
// @description
// @version
import (
	"fmt"
	"learn-go/feature"
	//"net/http"
	//_ "net/http/pprof" 可以动态监测pprof
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

	//go func() {
	//	_ = http.ListenAndServe("localhost:8085", nil)
	//}() // 开一个goroutine来监测程序的状态

	//init() 不可以调用
	var mainFunc F = main
	//var initFunc F = init 无法引用到init
	//mainFunc() 引发死递归
	fmt.Printf("function main type v initFunc: %T\n\n", mainFunc)
	learnFunctions := []func(){
		feature.LearnMapOP,
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
		//feature.LearnMem,
		feature.LearnMem2,
		//feature.LearnMem3,
		feature.LearnGoroutine,
		feature.LearnAlgUse,
		feature.LearnIterate,
		//golang 编译器的一个bug,已经从github上拿到了1.14.7的源代码,改了path, 不再有这个问题
		// 关于后续的优化: https://github.com/golang/go/issues/40502
		feature.LearnUnsafe,
		feature.LearnGoroutine2,
		feature.LearnChan,
		feature.LearnChan2,
		feature.LearnTimer,
		feature.LearnChan3,
		feature.LearnServer,
		feature.LearnGoroutine3,
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
