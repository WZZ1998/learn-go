package feature

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/trace"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 04:14
// @description
// @version

func LearnSchedule() {
	if runtime.GOOS != "darwin" {
		fmt.Println("not on mac darwin, return!")
		return
	}
	oldMaxProcs := runtime.GOMAXPROCS(1)
	fmt.Println("old Max Proc:", oldMaxProcs)
	defer runtime.GOMAXPROCS(oldMaxProcs)
	const profileDirPth = "/Users/wangzizhou/Downloads/learn-go-profiles"
	tf, _ := os.Create(path.Join(profileDirPth, "scheTrace"))
	defer tf.Close()
	cc := make(chan int, 100)
	if errT := trace.Start(tf); errT != nil {
		fmt.Println("trace failed:", errT)
	}
	go producer(cc)
	go consumer(cc)
	go infiniteLoop() // hang住一个线程
	time.Sleep(100 * time.Millisecond)
	trace.Stop()
}
func infiniteLoop() {
	for {
	}
}
func consumer(cc chan int) {
	for {
		sp := make([]int, 1024*1000)
		sp[127] = <-cc
	}
}

func producer(cc chan int) {
	c := 0
	for {
		cc <- c
		c++
	}
}
