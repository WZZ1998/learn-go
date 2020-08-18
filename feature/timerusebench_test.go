package feature_test

import (
	"math/rand"
	"testing"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/18 15:06
// @description
// @version

func BenchmarkUseNewTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rCh := make(chan int, 1)
		go simulateCallTimeout(rCh)
		callTimer := time.NewTimer(10 * time.Millisecond)
		select {
		case res := <-rCh:
			_ = res
		case <-callTimer.C:
			b.Log("call failed. timeout!")
		}
		callTimer.Stop()
	}
}
func BenchmarkReuseTimer(b *testing.B) {
	callTimer := time.NewTimer(0)
	if !callTimer.Stop() {
		<-callTimer.C // 卡在这等,没问题,但是这里用select加default也不行
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rCh := make(chan int, 1)
		go simulateCallTimeout(rCh)
		callTimer.Reset(10 * time.Millisecond)
		select {
		case res := <-rCh:
			_ = res
			if !callTimer.Stop() {
				<-callTimer.C
			}
		case <-callTimer.C:
			b.Log("call failed. timeout!")
		}
	}
}
func timeCostCall(string) int {
	time.Sleep(20 * time.Microsecond)
	return rand.Intn(200)
}

func simulateCallTimeout(resCh chan<- int) {
	resCh <- timeCostCall("call!")
	close(resCh)
}
