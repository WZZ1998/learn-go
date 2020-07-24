package feature

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/17 19:10
// @description
// @version
func LearnGoroutine() {
	if runtime.GOOS == "darwin" {
		doTrace := true
		fmt.Print("go max P:", runtime.GOMAXPROCS(-1))
		fmt.Println(" CPU cnt:", runtime.NumCPU())
		if doTrace {
			trF, errT := os.Create(path.Join(profileDirPth, "goroutineCPU.trace"))
			if errT != nil {
				fmt.Println("trace file open failed:", errT)
				return
			}
			if errStartTrace := trace.Start(trF); errStartTrace != nil {
				fmt.Println("trace start error:", errStartTrace)
				return
			}
			defer trace.Stop()
		}

		for i := 0; i < 2; i++ {
			st := time.Now()
			runtime.Gosched() //yield CPU
			fmt.Println("after schedule:", time.Since(st))
		}
	}

	ch := make(chan int)
	fmt.Println("in main goroutine, ch addr:", &ch)
	go getIntFromChan(ch)
	for i := 0; i < 50; i++ {
		ch <- rand.Intn(100)
	}
	time.Sleep(5 * time.Millisecond)
	for i := 0; i < 3; i++ {
		giveGoroutineAndSendBackWithCh()
	}

	//ch3 := make(chan int)
	//go func() {
	//	for {
	//		time.Sleep(60 * time.Millisecond)
	//		ch3 <- rand.Intn(30)
	//	}
	//}()
	//st := time.Now()
	//for time.Since(st) <= 1*time.Second {
	//	fmt.Print(<-ch3, " ")
	//}
	//fmt.Println()

	//var ch3SendOnly chan<- int = ch3
	//ch3SendOnly <- 0xffff // 单向channel
	//var ch3ReceiveOnly <-chan int
	//ch3ReceiveOnly = ch3
	//x := <-ch3ReceiveOnly
	ch4 := make(chan int, 4)
	chSyWithReceiver := make(chan bool)
	go func() {
		<-chSyWithReceiver
		for x, ok := <-ch4; ok; x, ok = <-ch4 {
			fmt.Println("x, ok:", x, ok)
		} // 拿到的并不是按顺序,因为放的时候是随机调度
		chSyWithReceiver <- true

	}()

	senderWg := new(sync.WaitGroup)
	senderWg.Add(6)
	for i := 0; i < 6; i++ {
		go func(sendI int) {
			defer func() {
				if errG := recover(); errG != nil {
					// close chan之后,所有挂起的
					fmt.Println("goroutine inside recover, err:", errG)
				}
				senderWg.Done()
			}()
			ch4 <- sendI //这里如果直接写i,问题就很大,因为直接闭包捕获循环变量i
			// 无法保证每次循环恰好goroutine都能得到调度,因此往往最后拿到的都是最后的值
			// 循环变量仅仅初始化一次,所以每个goroutine捕获的是同一个变量.
			// 两种解决方法: 改为lambda的参数输入,或者在循环体内开个新变量做中间值
			fmt.Println("finish from send to ch4")
		}(i) // 这里发生值拷贝
	}
	// 等一会让sender发出去,这里的close和上面循环的go形成了一个可以检测到的race
	// 不过这里等一段时间协程写channel基本写完了
	time.Sleep(60 * time.Millisecond)
	close(ch4)
	senderWg.Wait()
	chSyWithReceiver <- true
	fmt.Println("finish close chan ch4")
	<-chSyWithReceiver

	// 注意,for range循环中,每次循环体的循环变量地址都一样.也就是初始化一次,赋值多次
	//sl := []int{0xf, 0xf, 0xf, 0xf}
	//for ix, el := range sl {
	//	println("ix el addr:", &ix, &el)
	//}
	//mp := map[string]int{"hello": 1, "你好": 2, "Good Morning": 3}
	//for k, v := range mp {
	//	println("k v addr:", &k, &v)
	//}

}
func sendIntToCh(ch chan int, num int) {
	ch <- num
}
func getIntFromChan(ch chan int) {
	fmt.Println("in getIntFromChan ch addr:", &ch)
	for {
		if <-ch%17 == 0 { // 可以这样直接用来做判断
			fmt.Println("divided by 17!")
		}
	}
}

func giveGoroutineAndSendBackWithCh() {
	ch2 := make(chan int)
	for i := 0; i < 10; i++ {
		go sendIntToCh(ch2, i) // 开协程的时候按顺序分发
	}
	fmt.Print("Get int back: ")
	for i := 0; i < 10; i++ {
		fmt.Print(<-ch2, " ") // 回收到的是乱序的,展现了goroutine调度的随机性
	}
	fmt.Println()
}
