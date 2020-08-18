package feature

import (
	"fmt"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/18 18:17
// @description
// @version

func LearnTimer() {
	defer func() {
		if errTimer := recover(); errTimer != nil {
			fmt.Println("recover runtime panic:", errTimer)
		}
	}()
	for {
		runTimer()
	}
}
func reTimer(t *time.Timer, d time.Duration) {
	if !t.Stop() {
		select {
		case <-t.C: // 问题就出在这里
		// 为什么出在这里? 这里的select就是,想取值又不想等,所以导致竞态
		// 解决方法就是,放弃Reset,或者搞清楚到底t.C里面有没有值,
		//确定会有值,没别人抢, 然后直接挂起在channel上,把其抽干

		default:
			fmt.Println("t expired.But the channel is not OK.")
		}
	}
	t.Reset(d) // Reset是由竞态的,只应该在抽干channel且stopped或者expired的情况下调用
}

func runTimer() {
	tmr := time.NewTimer(0)
	reTimer(tmr, time.Minute)
	select {
	case <-tmr.C:
		panic("unexpected firing of Timer") // 这玩意糟糕透实锤了
	default:
	}
}
