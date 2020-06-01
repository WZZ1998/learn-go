package feature

// @author  wzz_714105382@icloud.com
// @date  2020/5/30 23:04
// @description
// @version

import(
	"fmt"
	"unicode/utf8"
)
func prettyLen(name string, l int) {
	fmt.Printf("%s len = %d\n", name, l)

}
func LearnString() {
	fmt.Printf("***string len analysis by %s***:\n\n", "wzz")
	sl := []string{"昨天我吃了一大碗过水面", "hello", "昨天我去park"}
	for _, s := range sl {
		fmt.Printf("context:%s,len:%d, chars:%d\n",
			s, len(s), utf8.RuneCountInString(s))
	}
	println("\n***iterating string***\n")
	for _, c := range "小轩窗正梳妆" {
		fmt.Printf("%c ", c)
	}
	println()
}