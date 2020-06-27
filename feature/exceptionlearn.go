package feature

import (
	"errors"
	"fmt"
	"os"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/25 22:22
// @description
// @version
func LearnException() {
	defer func() {
		fmt.Println("defer call in LearnException")
	}()

	err := errors.New("my strange error")
	fmt.Println("err:", err)
	//panic("my strange panic")
	fmt.Println("user:", os.Getenv("USER"))
	//err2 := fmt.Errorf("this is error created by %s", "MLRX")
	//panic("Error occurred:" + err2.Error())
	callWithRecover()

	fmt.Println("LearnException body ended.")
}
func callWithRecover() {
	defer func() {
		fmt.Println("done")

	}()
	fmt.Println("start")
	ttf1()
}

func ttf1() {
	defer func() {
		if err := recover(); err != nil { // 给捕获了
			fmt.Println("recover runtime panic:", err)
			// recover 之后,如果有新的panic,那么这个被recover的panic结构体会被保留下来,
			// 否则就要连着后面aborted的那个panic结构体,一起被清除
		}
		//panic("another panic right after recover")
		// 比如,将这个panic挪到callWithRecover的defer中,panic报错只显示这一个panic
		// 因为之前的两个都被清除了

	}()
	defer func() {
		fmt.Println("defer call in ff1")
		panic("panic in ff1 defer!") // 这个panic使得前面的panic被设置为aborted
	}()
	ttf2()
	fmt.Println("in ttf1 body...") // 这句永远调用不到的,panic会直接走defer,然后返回
}

func ttf2() {
	defer func() {
		fmt.Println("defer call in ff2")
	}()
	//fmt.Println("recover:", recover()) 直接调用recover只返回nil,没有其他效果了
	panic("deliberate panic!")

}
