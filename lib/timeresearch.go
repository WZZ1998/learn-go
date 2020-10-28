package lib

import (
	"fmt"
	"strings"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/26 20:47
// @description
// @version
func playWithTime() string {
	var sb strings.Builder
	wd := func(a ...interface{}) {
		sb.WriteString(fmt.Sprintln(a...))
	}
	wd(time.Now().UnixNano()) // 上限2262年
	//wd(time.Now())
	loc, _ := time.LoadLocation("America/Los_Angeles")
	wd(time.Unix(0, time.Now().UnixNano()).In(loc))
	// 对于Time的String()方法,有如下说明:
	// The returned string is meant for debugging; for a stable serialized
	// representation, use t.MarshalText, t.MarshalBinary, or t.Format
	// with an explicit format string.
	myTime := time.Date(2015, 6, 30, 23, 59, 56, 0, time.UTC)
	myTime = time.Now()
	for i := 0; i < 2; i++ {
		myTime = myTime.Add(-1 * time.Second)
		wd(myTime)
	}
	t, err := time.Parse("2006-01-02 15:04:05 -0700", "2015-06-30 14:20:20.123123 +0800")
	// 2006-01-02 15:04:05 -0700
	if err != nil {
		panic(err)
	}
	wd(t)
	return sb.String()
}
func PlayWithTime() string {
	return playWithTime()
}
