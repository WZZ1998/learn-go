package feature

// @author  wzz_714105382@icloud.com
// @date  2020/6/10 16:04
// @description 学习随机数的使用
// @version
import (
	"fmt"
	"math/rand"
	"time"
)

func LearnRandom() {
	fmt.Println("random int and float32")
	timeNS := int64(time.Now().Nanosecond())
	rand.Seed(timeNS) //使用指定的随机种子
	for i := 0; i < 10; i++ {
		r := rand.Intn(5000)
		fmt.Printf("%v  ", r)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
}
