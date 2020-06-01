package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/1 22:46
// @description golang name and type
// @version
type MI int

func LearnTypes() {
	//_ = 1
	//fmt.Println(_)
	//underscore 不能被使用

	var mii MI = 100
	var it int = 100
	//mii = it 无法通过编译,golang遵守静态类型&强类型
	fmt.Printf("MI(int) type v mii is :%T\n", mii)
	fmt.Printf("int type v it is :%T\n", it)

}
