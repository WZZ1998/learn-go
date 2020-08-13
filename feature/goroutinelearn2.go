package feature

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/10 21:51
// @description
// @version

type synDataS struct {
	name string
	age  int
}

func LearnGoroutine2() {
	myCh1 := make(chan int, 20)
	//go checkGetThenCheck(myCh1)
	go checkGetThenCheck(myCh1, "receive goroutine 1")
	go checkValueWithRange(myCh1, "receive goroutine 2")
	// 两个goroutine在分配数量较大的时候,调度还是比较均等的
	time.Sleep(1 * time.Millisecond)
	for i := 0; i < 200; i++ {
		myCh1 <- 50 + rand.Intn(500000)
		time.Sleep(1 * time.Millisecond)
	}
	time.Sleep(100 * time.Millisecond)
	close(myCh1)
	time.Sleep(3 * time.Millisecond)

	mD := synDataS{"foo", 100}
	fmt.Printf("in learn, mD %v addr %p\n", mD, &mD)
	dataCh := make(chan synDataS)
	mCh := make(chan bool)
	go checkChannelSendData(dataCh, mCh)
	dataCh <- mD
	<-mCh
	fmt.Printf("after change, in learn, mD %v addr %p\n\n", mD, &mD)

	synCh := make(chan synDataS)
	mD2 := &synDataS{
		name: "unchanged!",
		age:  20,
	}
	go func() {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("in receive lambda, get:", <-synCh)
	}()
	go func() {
		time.Sleep(5 * time.Millisecond)
		mD2.name = "name changed in lambda"
	}()
	synCh <- *mD2
	time.Sleep(50 * time.Millisecond)
	fmt.Printf("md2 now: %p %v\n", mD2, mD2)

	// channel 变量的零值是nil

	pt := "^((?!Std).)*$"
	fmt.Println(regexp.MatchString(pt, "StdTest"))
}

func checkGetThenCheck(ch <-chan int, name string) {
	lCnt := 0
	for {
		_, b := <-ch
		if b == false {
			break
		}
		lCnt++
	}
	// 在通道关闭之后,接收方从通道只能拿到对应类型的零值,并且会立即返回
	ov, ob := <-ch
	fmt.Println(name, " after break from loop, get :", ov, ob, "total cnt:", lCnt)
	// 注意,向已经关闭的channel发送,或者关闭一个已经关闭的channel,都会导致panic
	//可以包装一层,在内部recover一下
	fmt.Println(name, "func return.")
}

func checkValueWithRange(ich <-chan int, name string) {
	lCnt := 0
	for range ich {
		lCnt++
		//fmt.Println(name, " received:", v)
	}
	// range会持续等待,直到通道被关闭时,退出循环
	ov, ob := <-ich
	fmt.Print(name, " out of range, rec ich:", ov, ob)
	fmt.Printf(" totally %d int\n", lCnt)
	fmt.Println(name, "func return.")
}

func checkChannelSendData(dc <-chan synDataS, cWaitForMo chan<- bool) {
	d := <-dc
	fmt.Printf("addr %p d %v\n", &d, d)
	d.name = "mo"
	d.age = 20
	fmt.Printf("after change addr %p d %v\n", &d, d)
	cWaitForMo <- true
}
