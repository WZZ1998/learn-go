package feature

// @author  wzz_714105382@icloud.com
// @date  2020/6/29 04:14
// @description
// @version

func producer(ch chan<- int) {
	i := 0
	for {
		ch <- i
		i++
	}
}

func consumer(ch <-chan int) {
	for {
		b := make([]byte, 512*1024)
		println(<-ch, b[0])
	}
}
func LearnSchedule() {
	ch := make(chan int, 100)
	go producer(ch)
	go consumer(ch)
	for {
	}
}
