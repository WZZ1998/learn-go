package feature

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/28 10:32
// @description
// @version

func LearnContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	// 这里应该这么写, 在操作提前返回时,释放context的资源
	// 对于一个已经timeout的context,cancel似乎没有什么作用
	rc := make(chan string)
	st := time.Now()
	go doSomethingWithContext(ctx, rc)
	fmt.Println(<-rc)
	fmt.Println(time.Since(st))

	//overWithSignal()

}

func doSomethingWithContext(ctx context.Context, rc chan<- string) {
	execRes := ""
	tmr := time.NewTimer(1 * time.Second)
	select {
	case <-tmr.C:
		execRes = "normally exec"
	case <-ctx.Done():
		execRes = "context timeout!"
	}
	tmr.Stop()
	rc <- execRes
	// 无论如何,也需要阻塞在通道上跳出执行流程
	//goroutine貌似难以褫夺执行然后返回
	//context适用于取消一组具有执行关联的goroutine
}

func overWithSignal() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch := make(chan os.Signal)
		// 将信号注册到对应的channel来进行接收和处理
		// 接收SIGINT会导致原来的默认行为失效, 可以拦截信号
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(ch)
		// 停止之后,再收到的signal不会再发送到ch,这样就取消了signal和channel的绑定
		//对signal的处理也会恢复到默认行为
		ss := <-ch
		fmt.Println("received:", ss)
		time.Sleep(2 * time.Second)

		return
	}()
	wg.Wait()
	fmt.Println("main return")

}
