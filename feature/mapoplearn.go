package feature

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/26 10:44
// @description
// @version
func LearnMapOP() {
	//runtime.GC() // 手动激发一次GC,防止被干扰
	time.Sleep(5 * time.Millisecond)

	doProfile := false
	if runtime.GOOS == "darwin" && doProfile == true {
		oPth := "/Users/wangzizhou/Downloads/learn-go-profiles"
		pf, errCreatePF := os.Create(path.Join(oPth, "mapCPU.prof"))
		if errCreatePF != nil {
			fmt.Println("create prof file failed, error:", errCreatePF)
			return
		}
		_ = pprof.StartCPUProfile(pf)
		defer pprof.StopCPUProfile()

		//tf, errCreateTrace := os.Create(path.Join(oPth, "map1000000.trace"))
		//if errCreateTrace != nil {
		//	fmt.Println("create trace file error:", errCreateTrace)
		//}
		//trace.Start(tf)
		//defer trace.Stop()
	}

	st := time.Now().UnixNano()
	kBuf := make([]byte, 0, 32)
	mm := make(map[string]int64, 1000000)
	var i int64
	for i = 0; i < 1000000; i++ {
		v := time.Now().Unix()
		kBuf = strconv.AppendInt(kBuf, i, 10)
		kBuf = append(kBuf, '_')
		kBuf = strconv.AppendInt(kBuf, v, 10)
		mm[string(kBuf)] = v
		kBuf = kBuf[:0]
	}
	en := time.Now().UnixNano()
	fmt.Println("time ms elapsed:", (en-st)/1e6)
}
