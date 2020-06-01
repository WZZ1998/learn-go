/* Package feature contains feature study of golang.

   This package contains my study and research of basic
   feature in golang, such as string, type, standard libs,
   and so on.
*/
package feature

// @author  wzz_714105382@icloud.com
// @date  2020/5/30 23:04
// @description
// @version

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	fmt.Println("initializing package feature......")
}

// LearnString includes some go string feature study, printing some lines.
func LearnString() { // 注意，导出对象的注释一定要用该对象的标识符开始
	fmt.Printf("***string len analysis by %s***:\n", "wzz")
	sl := []string{"昨天我吃了一大碗过水面", "hello", "昨天我去park"}
	for _, s := range sl {
		fmt.Printf("context:%s,len:%d, chars:%d\n",
			s, len(s), utf8.RuneCountInString(s))
	}
	s := "hello!!!" // 使用%T来填入变量的类型
	fmt.Printf("variable type s string: %T\n", s)

	foreignS := "Καλημέρα κόσμε; or こんにちは 世界"
	fmt.Printf("Golang支持完全国际化:%s\n", foreignS)

	fmt.Println("\n***iterating string***")
	for _, c := range "小轩窗正梳妆" {
		fmt.Printf("%c ", c)
	}
}

func prettyLen(name string, l int) {
	fmt.Printf("%s len = %d\n", name, l)
}
