package feature

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"runtime"
	"runtime/trace"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/17 19:10
// @description
// @version
func LearnGoroutine() {
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
	ch3 := make(chan int)
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			ch3 <- rand.Intn(30)
		}
	}()
	st := time.Now()
	for time.Since(st) <= 2*time.Second {
		fmt.Print(<-ch3, " ")
	}
	fmt.Println()
	//var ch3SendOnly chan<- int = ch3
	//ch3SendOnly <- 0xffff // 单向channel
	//var ch3ReceiveOnly <-chan int
	//ch3ReceiveOnly = ch3
	//x := <-ch3ReceiveOnly
	ch4 := make(chan int, 4)
	go func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("before block on receive")
		for x, ok := <-ch4; ok; x, ok = <-ch4 {
			fmt.Println("x, ok:", x, ok)
		}

	}()
	for i := 0; i < 20; i++ {
		go func() {
			defer func() {
				if errG := recover(); errG != nil {
					// close chan之后,所有挂起的
					fmt.Println("goroutine inside recover, err:", errG)
				}
			}()
			fmt.Println("before send")
			ch4 <- 0xffffff
			fmt.Println("finish from send to ch4")
		}()
	}
	time.Sleep(1 * time.Second)
	close(ch4)
	fmt.Println("finish close chan ch4")
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch4)
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
