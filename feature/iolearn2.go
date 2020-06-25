package feature

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/24 13:58
// @description
// @version

func LearnIO2() {
	env := runtime.GOOS
	fmt.Println("env OS:", env)
	var resPth string
	if env == "linux" {
		resPth = "resources/"
	} else {
		resPth = "/Users/wangzizhou/GolandProjects/learn-go/out/resources/"
	}

	fName := "a"
	inputFile, errOpen := os.Open(filepath.Join(resPth, fName))
	if errOpen != nil {
		fmt.Println("open file failed, errOpen: ", errOpen)
		return
	}
	defer func() {
		_ = inputFile.Close()
	}()
	bufRdr := bufio.NewReader(inputFile)
	for {
		cont, err := bufRdr.ReadString('\n')
		// 注意,读进来的字符串自己带着结尾的换行符
		cont = strings.TrimRight(cont, "\n")
		fmt.Println("read string:", cont)
		if err == io.EOF {
			fmt.Println("EOF err:", err)
			break
		}
	}

	bytes, err := ioutil.ReadFile(filepath.Join(resPth, fName)) // 这种不用defer关闭,自己会关闭
	if err != nil {
		fmt.Println("whole read failed, err:", err)
		return
	}
	fmt.Println("whole file bytes len:", len(bytes))

	f2, _ := os.Open(filepath.Join(resPth, fName))
	defer f2.Close()

	x := bufio.NewReader(f2) // 默认大小 4096
	tarSl, err2 := x.ReadString('\n')
	fmt.Println("n : ", len(tarSl), " err2:", err2)
	for i, ll := 0, len(tarSl); i < ll; i++ {
		erB := x.UnreadByte() //真的可以回退! 但是只能回退1个字节或者字符
		if erB != nil {
			fmt.Println("unread error:", erB) //unread一次失败之后,后面再调用还是失败 invalid use
			break
		} else {
			fmt.Println("unread success")
		}
	}
	tarBSl := make([]byte, 128, 128)
	n, err3 := x.Read(tarBSl)
	// 给出字节小,就读满整个给定的区域;否则全读完,剩下空着;下一次再读的时候才会返回io.EOF的错误
	fmt.Println("n :", n, " err3:", err3)
	fmt.Printf("tarBSl :%v\n", tarBSl)

	// 从文件读进来的切片,在上面进行重组的时候,底层的数组是不变的,如果数组很大,内存可能占用较大
	//所以,如果想从读进来的大段内容中,获取有意义的信息,可以考虑将搜寻到的信息放到新的内存空间

	// Scanner 对token的大小有限制
	// Scanner is for safe, simple jobs.
	f3, _ := os.Open(filepath.Join(resPth, fName))
	defer f3.Close()
	mySc := bufio.NewScanner(f3)
	mySc.Split(bufio.ScanWords) // 这里设置按空格分隔
	fmt.Println("start scan loop:")
	for mySc.Scan() {
		fmt.Println(mySc.Text()) // 正常是按行扫
	}
	scErr := mySc.Err()
	fmt.Println("scan err:", scErr)

	s := strings.Repeat("Golang!", 20000)
	s += "\n"
	myStringRdr := strings.NewReader(s)
	bufStrRdr := bufio.NewReader(myStringRdr)
	bufStrRdr = bufio.NewReaderSize(bufStrRdr, 1e6)
	// 这玩意怎么搜的? 缓冲reader将数据读入buffer,并在其中寻找delimit,如果找到就返回,灌满buffer也没找到,就在ReadBytes中,
	// 把每个灌满的buffer做个copy,然后再次用新数据灌满buffer,如此往复,直到找到delimit.找到之后,返回,或者将得到的数据组拼接好,再返回
	// 这个问题启发我们, bufio.reader的缓冲区大小应该认真设定一下.在这个例子中,buffer被灌满了多次
	var sFromIO string
	var err4 error
	st := time.Now()
	for i := 0; i < 10000; i++ {
		sFromIO, err4 = bufStrRdr.ReadString('\n')
		myStringRdr.Reset(s)
	}
	tUsed := time.Since(st)
	// time used: 默认大小4096 844.424408ms 1e4 669.03669ms 1e6 506.319594ms
	fmt.Println("time used", tUsed)
	fmt.Println("len sFromIO", len(sFromIO), "err4 :", err4) // 没找到,返回EOF错误

	sPart := s[0:2000]
	f2Name := "b"
	// 本地有一个umask,会对权限做一下取反mask,在编写代码的环境中是0222 所以是组用户和去全部用户的权限是不能改的
	//old := syscall.Umask(0000)
	errWrite := ioutil.WriteFile(filepath.Join(resPth, f2Name), []byte(sPart), 0644) // 自己会关闭
	//syscall.Umask(old)
	fmt.Println("write err:", errWrite)

	wf1Name := "c"
	writeF1, errW := os.OpenFile(filepath.Join(resPth, wf1Name),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644)
	if errW != nil {
		fmt.Println("open write file error: ", errW)
		return
	}
	defer writeF1.Close() // 千万不要在某些时候手动关闭writeF1, 否则writeF1关闭了,writer还没flush

	wrt := bufio.NewWriter(writeF1)
	defer func() {
		errFlush := wrt.Flush()
		if errFlush != nil {
			panic(errFlush)
		}
	}()
	// 一定要Flush,不然写不进去;而且后注册的defer先执行, 也是为了这个
	// 要是用defer那就大家都defer,按先后次序来
	//要不就大家都不要defer,手动处理关闭顺序保证稳定
	//不要混用,非常容易出bug!

	wCont := "hello? hello! \n"
	_, errW2 := wrt.WriteString(wCont)
	if errW2 != nil {
		panic(errW2)
	}
	errW2 = wrt.Flush() // 这里写入了一次.被后面writeFile trunc了
	if errW2 != nil {
		panic(errW2)
	}

	err5 := ioutil.WriteFile(filepath.Join(resPth, wf1Name), []byte(wCont), 0644) // writeFile会trunc文件
	// 会自己close
	if err5 != nil {
		panic(err5)
	}
	_, err5 = fmt.Fprintf(writeF1, "this is written for learn time:%v\n", time.Now())
	// 注意使用场合,这个不需要负责文件或者writer的管理,只是用了个现成的
	if err5 != nil {
		panic(err5)
	}

	seFName := "d"
	seFile, seErr := os.Open(filepath.Join(resPth, seFName))
	if seErr != nil {
		panic(seErr) // 如果这里panic,下面的defer根本不会执行,因为都还没注册
	}
	defer func() { // 注意,在panic之后才注册的defer并不会被执行,但是之前注册的会被执行
		// (但是不能defer嵌套?)
		// panic -> 执行在此panic发生前函数的注册的defer -> 没有回复,触发fatalPanic
		// 总结:先判定error,然后defer 关闭
		seErr = seFile.Close()
		if seErr != nil {
			panic(seErr)
		}
	}()
	seRdr := bufio.NewReader(seFile)
	var rLine string
	for {
		//_, seErr := fmt.Fscanf(seFile, "%s\n", &rawLine) // scan

		rLine, seErr = seRdr.ReadString('\n')
		if seErr != nil && seErr != io.EOF {
			panic(seErr)
		}
		rLine = strings.TrimRight(rLine, "\n")
		var name string
		var price float64
		var cnt int
		parts := strings.Split(rLine, ";")
		//parts := strings.SplitAfter(rLine, ";") // 在分隔符之后切一刀
		//parts = strings.SplitN(rLine, ";", 2) // 最多只分成n个子串,没分的就统一放在最后一个子串
		name = strings.Trim(parts[0], "\"")
		price, _ = strconv.ParseFloat(parts[1], 64)
		cnt, _ = strconv.Atoi(parts[2])
		fmt.Printf("[name:%s  price:%.2f  cnt:%d]\n", name, price, cnt)
		if seErr == io.EOF {
			break
		}
	}

	offset, _ := seFile.Seek(0, 0)
	fmt.Println("after seek offset:", offset)
	// seek重设文件指针
	// whence 用来计算offset的初始位置 0 = 文件开始位置 1 = 当前位置 2 = 文件结尾处
	// offset 偏置字节数
	// 在 O_APPEND 打开的文件上,seek的行为是未知的
	csvRdr := csv.NewReader(seRdr)
	csvRdr.Comma = ';'
	csvAll, err8 := csvRdr.ReadAll()
	if err8 != nil {
		panic(err8)
	}
	fmt.Println("all:", csvAll)

}
