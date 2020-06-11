package feature

import (
	"fmt"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/11 18:38
// @description learn time and date of golang
// @version

// LearnTime contains some time & date operations for golang study.
func LearnTime() {
	now := time.Now()
	fmt.Printf("当前时间:%v \n", now)
	fmt.Printf("UTC 当前时间:%v \n", now.UTC())
	week := 7 * 24 * 60 * 60 * 1e9
	weekDua := time.Duration(week)
	timeOneWeekLater := now.Add(weekDua)
	fmt.Printf("one week later : %v \n", timeOneWeekLater)
	fmt.Println("1 week :", now, "=>", timeOneWeekLater)
	fmt.Printf("日月年:%02d.%02d.%4d\n", now.Day(), now.Month(), now.Year())
	fmt.Printf("ANSIC格式化时间: %v\n", now.Format(time.ANSIC))
	fmt.Printf("自定义原型格式化时间: %v\n", now.Format("02 ~~ Jan ~~ 2006 15:04 MyTime"))
	//seconds := time.Duration(3 * 1e9)
	//time.Sleep(seconds) 睡眠三秒
}
