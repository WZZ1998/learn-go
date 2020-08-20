package feature

import (
	"fmt"
	"sync"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/19 16:09
// @description
// @version

type triOp func(int, int, int) int

type myRequest struct {
	a, b, c int
	replyC  chan int
	reqOK   bool
}

func (r *myRequest) String() string {
	return fmt.Sprintf("myRequest: a %d b %d c = %d ok %t replyC %p", r.a, r.b, r.c, r.reqOK, r.replyC)
}

func simulateRunTripOp(op triOp, req *myRequest) {
	//grand := rand.New(rand.NewSource(time.Now().UnixNano()))
	//tn := grand.Intn(20) + 1
	//time.Sleep(time.Duration(tn) * time.Millisecond)
	time.Sleep(20 * time.Millisecond)
	req.replyC <- op(req.a, req.b, req.c)
}

func serverWithQC(op triOp, reqService <-chan *myRequest, quitC <-chan bool) {
	defer fmt.Println("server with quit terminated.")
	for {
		select {
		case rq := <-reqService:
			go simulateRunTripOp(op, rq)
		case <-quitC:
			return
		}
	}

}
func server(op triOp, reqService <-chan *myRequest) {
	defer fmt.Println("server terminated.")
	const MAXREQ = 900 // 带缓冲的channel可以直接作为信号量使用,P对应发送元素,
	// V对应取出元素;有N的缓冲位置的channel将在连续写入N个值后阻塞继续试图写入的goroutine,
	//而receive会使得空出一个位置,其他goroutine可以继续插入;如上所述就实现了对并发的控制
	sem := make(chan int, MAXREQ) // 限制同时处理的最大并发数量,模拟信号量
	for req := range reqService {
		go func(gReq *myRequest) {
			sem <- 1 // 如果还有空位就写入执行,否则挂起等待
			simulateRunTripOp(op, gReq)
			<-sem
		}(req)
	}
}

func sendRequest(reqService chan<- *myRequest) *myRequest {
	rqErrMess := ""
	q := &myRequest{
		a:      25,
		b:      25,
		c:      25,
		replyC: make(chan int, 1),
		reqOK:  false,
	}
	//defer func() {
	//	if !q.reqOK {
	//		fmt.Println(q, rqErrMess)
	//	}
	//}()

	defer func() {
		if errP := recover(); errP != nil {
			rqErrMess = fmt.Sprint("runtime panic:", errP)
		}
	}()
	sTimer := time.NewTimer(1 * time.Second)
	select {
	case reqService <- q:
	case <-sTimer.C:
		rqErrMess = "send request timeout."
		return q
	}
	sTimer.Stop()
	rTimer := time.NewTimer(3 * time.Second)
	select {
	case res := <-q.replyC:
		if res != q.a*q.b*q.c {
			rqErrMess = "wrong result."
			return q
		}
	case <-rTimer.C:
		rqErrMess = "wait receive timeout."
		return q
	}
	rTimer.Stop()
	q.reqOK = true
	return q
}

func startServer(op triOp) chan *myRequest {
	reqService := make(chan *myRequest, 1e3*20)
	go server(op, reqService)
	return reqService
}

func concurrentSendRq(rqSrv chan *myRequest) int {
	const cn = 10000
	bch := make(chan bool, 1024)
	wg := new(sync.WaitGroup)
	wg.Add(cn)
	for i := 0; i < cn; i++ {
		go func() {
			q := sendRequest(rqSrv)
			bch <- q.reqOK
			wg.Done()
		}()
	}
	failedResC := make(chan int)
	go func() {
		failed := 0
		for success := range bch {
			if !success {
				failed++
			}
		}
		failedResC <- failed
	}()
	wg.Wait()
	close(bch)
	return <-failedResC
}
func LearnServer() {
	triMultiply := func(a, b, c int) int { return a * b * c }
	reqService := startServer(triMultiply)
	f := concurrentSendRq(reqService)
	fmt.Println("failed:", f)
	close(reqService)
	//waitPauseOnKeyboardInputToContinue()
	time.Sleep(100 * time.Millisecond)
}

// 先用panic写,然后再用quit

func waitPauseOnKeyboardInputToContinue() {
	fmt.Print("press enter to continue:")
	_, _ = fmt.Scanln()
}
