package feature

import (
	"flag"
	"fmt"
	"os"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/26 17:07
// @description
// @version
func LearnIO3() {
	//fmt.Println("os args:", os.Args, "len:", len(os.Args))
	newL := flag.Bool("n", false, "print newline") // 默认值为false
	// flag 的话 -n 和 --n 都是有效的
	flag.PrintDefaults()
	flag.Parse()
	s := ""
	fmt.Println("args cnt:", flag.NArg())   // 参数个数,不含被解析了的flag
	fmt.Println("flags cnt:", flag.NFlag()) // flag 个数
	for i := 0; i < flag.NArg(); i++ {      //flag.Arg 的 0号元素是第一个真正意义上的参数,而不是程序的名称
		if i > 0 {
			s += " "
			if *newL {
				s += "\n"
			}
		}
		s += flag.Arg(i)
	}
	_, err := os.Stdout.WriteString(s)
	if err != nil {
		panic(err)
	}

}
