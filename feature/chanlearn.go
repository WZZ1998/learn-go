package feature

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/14 18:40
// @description
// @version

func LearnChan() {
	//ff, _ := os.Create(path.Join(profileDirPth, "hash256cpu.pprof"))
	//pprof.StartCPUProfile(ff)
	//defer pprof.StopCPUProfile()
	var Emp struct{}
	senderN := 4
	receiverN := 8
	wg := new(sync.WaitGroup)
	wg.Add(senderN)
	wg.Add(receiverN)

	dataCh := make(chan int64, 100)
	//moCh := make(chan struct{})
	//defer close(moCh)
	//go func() {
	//	for {
	//		select {
	//		case <-moCh:
	//			fmt.Println("mo len goroutine return")
	//			return
	//		default:
	//			fmt.Println(len(dataCh))
	//			time.Sleep(1 * time.Second)
	//		}
	//	}
	//}()

	stopCh := make(chan struct{})
	toStop := make(chan struct{}, 1)
	tCh := make(chan time.Time)
	// moderator
	go func() {
		ttt := time.Now()
		tmr := time.NewTimer(500 * time.Millisecond)
		defer tmr.Stop()
		select {
		case <-toStop:
			break
		case v := <-tmr.C:
			fmt.Println("timeout:", v.Sub(ttt))
		}
		close(stopCh)
		tCh <- time.Now()
	}()
	//cutP := 21922719 + 1
	cutP := 0
	for i := cutP; i < cutP+senderN; i++ {
		var gPoint = int64(i)
		go func() {
			defer wg.Done()
			//gran := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				rv := gPoint
				gPoint += int64(senderN)
				select {
				case <-stopCh:
					return
				case dataCh <- rv:
				}
			}
		}()
	}
	for i := 0; i < receiverN; i++ {
		go func() {
			defer wg.Done()
			rB := make([]byte, 0, 96)
			var nb = []byte("wangzizhou")
			rB = append(rB, nb...)
			stdB := []byte{102, 102, 102, 102}
			h256 := sha256.New()
			sha256Buf := make([]byte, 0, 32)
			for {
				select {
				case <-stopCh:
					return
				case v := <-dataCh:
					rB = strconv.AppendInt(rB, v, 10)
					//md5b := md5.Sum(rB)
					h256.Reset()
					h256.Write(rB)
					h256ResB := h256.Sum(sha256Buf)
					if bytes.Equal(stdB, h256ResB[:len(stdB)]) {
						fmt.Printf("rB %s sha256 B %x\n", string(rB), h256ResB)
						select {
						case toStop <- Emp:
						default:
						}
						return
					}
					rB = rB[:len(nb)]
				}
			}
		}()
	}
	wg.Wait()
	cTime := <-tCh
	fmt.Println("all clear after close, time used:", time.Since(cTime))
}
