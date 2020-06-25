package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/25 22:22
// @description
// @version
func LearnException() {
	defer func() {
		fmt.Println("defer call in LearnException")
	}()
	ttf1()
}

func ttf1() {
	defer func() {
		fmt.Println("defer call in ff1")
	}()
	ttf2()
}

func ttf2() {
	defer func() {
		fmt.Println("defer call in ff2")
	}()

	panic("deliberate panic!")
}
