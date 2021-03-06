package feature

// @author  wzz_714105382@icloud.com
// @date  2020/6/10 16:04
// @description 学习随机数的使用
// @version
import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func LearnRandom() {
	x := time.Now()
	fmt.Println("Nanosecond() means nano part of the Time:", x.Nanosecond())
	fmt.Println("UnixNano() means the unix nano elapsed since 1970.XXX", x.UnixNano())
	fmt.Println("random int and float32")

	// 如果不设置随机种子, 默认的全局随机Rand的种子是1,也就是说如果不设置,那么每次的序列都一样
	// 第一个5577006791947779410 第二个 8674665223082153551
	rand.Seed(time.Now().UnixNano()) //使用指定的随机种子
	for i := 0; i < 10; i++ {
		r := rand.Intn(5000) //这种用法会拿一个全局lock,有可能影响性能
		fmt.Printf("%v  ", r)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f / ", 100*rand.Float32())
	}
	fmt.Println()
	fmt.Println("ABout crypto rand:")
	rbySlice := make([]byte, 48)
	for i := 0; i < 3; i++ {
		n, errReadR := crand.Read(rbySlice)
		if errReadR != nil {
			fmt.Println("Read from crypto rand failed:", errReadR)
			break
		}
		fmt.Printf("read %d bytes from crypto rand.\n", n)
		b64RandStr := base64.StdEncoding.EncodeToString(rbySlice)
		fmt.Println("get random base64 str:", b64RandStr)
	}
	for i := 100; i < 10000; i += 2000 {
		bi := new(big.Int)
		bi.SetInt64(int64(i))
		fmt.Print("bound ", bi)
		n, errRI := crand.Int(crand.Reader, bi) //注意,bound <= 0 会直接panic!
		if errRI != nil {
			fmt.Println(" crand int failed, error:", errRI)
			break
		}
		fmt.Println(" cr big int:", n)

	}
}
