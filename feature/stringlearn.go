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
	"strconv"
	"strings"
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
	fmt.Println()
	fmt.Print(`非解释字符串 \n直接输出 可以直接换行
行2
行3`)
	fmt.Println("123###" + // 加号必须放在上一行,因为会自动补齐分号
		"456")

	poemLine := "春江潮水连海平,海上明月共潮生\n"
	index := 4
	fmt.Printf("the [%d] of poemline: %c\n", index, poemLine[index])
	rPoemLine := []rune(poemLine)
	fmt.Printf("the [%d] of []rune rPoemline: %c\n", index, rPoemLine[index])
	re := strings.NewReader(poemLine)
	fmt.Println(re.Len())
	spc := make([]byte, 22)
	n, _ := re.Read(spc)
	fmt.Printf("bytes read from reader: %v contents %v\n", n, spc)
	var sb strings.Builder
	sb.Write(spc)
	fmt.Printf("the string read from reader is %v\n", sb.String())
	in2 := 1567
	in2S := strconv.Itoa(in2)
	fl2 := 3.445579
	fl2S := strconv.FormatFloat(fl2, 'e', -1, 64)
	fmt.Printf("in2 type %T value %v convert to %T value %v\n", in2, in2, in2S, in2S)
	fmt.Printf("fl2 type %T value %v binary convert to %T value %v\n", fl2, fl2, fl2S, fl2S)
	wrongIntS := "happy!"
	resI, errI := strconv.Atoi(wrongIntS)
	resF, errF := strconv.ParseFloat(fl2S, 64)

	if errI != nil {

		fmt.Printf("转换失败:%v\n", errI)
	} else {
		fmt.Printf("转换成功,结果:%d\n", resI)
	}
	fmt.Printf("浮点数转化: type %T value %v error %v", resF, resF, errF)

}
