package feature

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"unicode/utf8"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/21 15:22
// @description
// @version

func LearnCode() {
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	substr := str[1:3] // 这样还是按照字节截取的
	fmt.Println(s2, substr)

	// 注意截取字符时需要转化为[]rune再截取,然后再转回来
	myS := "一二三四五六七八九十"
	par := string([]rune(myS)[5:7])
	fmt.Println(par, utf8.RuneCountInString(par))

	isa := [9]int{1: 2, 4: 6, 0: -99} // 直接用索引指定初始元素
	is := isa[:]

	fmt.Println(is)

	mm := make(map[int]string)
	pr := "m_pre"
	for i := 0; i < 100000; i++ {
		mm[i] = pr + strconv.Itoa(i)
	}
	fmt.Println(len(mm))
	for i := 0; i < 100000; i++ {
		delete(mm, i)
	}
	fmt.Println("after delete, mm:", len(mm))

	var a interface{} = new(mySTT)
	if v, ok := a.(fmt.Stringer); ok {
		fmt.Printf("implements String(): %s\n", v.String())
	}
	classifier(a, 1, "jello")
	if runtime.GOOS == "darwin" {
		fileOp()
	}

	t := 1
	t += 010
	fmt.Println(t)
}

type mySTT struct {
}

func (receiver *mySTT) String() string {
	return "useless"
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func fileOp() {
	file, err := os.Open("/Users/wangzizhou/GolandProjects/learn-go/out/resources/a")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer file.Close()
	myCat(file)
	nP, errSeek := file.Seek(0, 0)
	if errSeek != nil {
		panic(errSeek)
	}
	fmt.Println("successfully seeked to ", nP)
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			fmt.Println("-----err:", err)
			return // error or EOF
		}
		fmt.Printf("The input was: %s", str)
	}

}

func myCat(f *os.File) {
	const NBUF = 256
	var buf [NBUF]byte
	for { // 非常经典的IO switch
		switch nr, rerr := f.Read(buf[:]); true {
		case nr < 0:
			_, _ = fmt.Fprintf(os.Stderr, "myCat error reading from %s:%s\n", f.Name(), rerr.Error())
			panic(rerr)
		case nr == 0:
			return
		case nr > 0:
			if nw, werr := os.Stdout.Write(buf[:nr]); nr != nw {
				_, _ = fmt.Fprintf(os.Stderr, "myCat error writing from %s :%s\n", f.Name(), werr.Error())
			}
		}

	}
}
