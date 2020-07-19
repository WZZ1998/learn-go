package utils

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/19 23:09
// @description
// @version

func init() {
	fmt.Println("init utils/randomdata: set global rand seed.")
	rand.Seed(time.Now().UnixNano())
}

const (
	maxIntRSliceLen          int = 4e9
	rSliceLenConcurrentBound int = 2e5
)

func GetRandIntSliceOfLength(targetLen int) ([]int, error) {
	if targetLen > maxIntRSliceLen {
		return nil, fmt.Errorf("too big target length %d (max %d)", targetLen, maxIntRSliceLen)
	}
	if targetLen >= rSliceLenConcurrentBound {
		return concurrentGenerateRIntSliceOfLen(targetLen), nil
	} else {
		return generateRIntSliceOfLen(targetLen), nil
	}

}

func generateRIntSliceOfLen(ll int) []int {
	lra := getRandWithUnixNanoSrc() // only in this func
	res := make([]int, ll)
	for i := 0; i < ll; i++ {
		res[i] = lra.Int()
	}
	return res
}

func concurrentGenerateRIntSliceOfLen(ll int) []int {
	nCpu := runtime.NumCPU()
	nNewRoutine := nCpu
	di := ll / nNewRoutine
	res := make([]int, ll)
	wg := new(sync.WaitGroup)
	for i := 0; i < nNewRoutine; i++ {
		wg.Add(1)
		go func(startIx, endIx int) {
			gRa := getRandWithUnixNanoSrc()
			for ix := startIx; ix < endIx; ix++ {
				res[ix] = gRa.Int()
			}
			wg.Done()
		}(i*di, (i+1)*di)
	}
	mRa := getRandWithUnixNanoSrc()
	for i := nNewRoutine * di; i < ll; i++ {
		res[i] = mRa.Int()
	}
	wg.Wait()
	return res
}

func getRandWithUnixNanoSrc() *rand.Rand {
	rSrc := rand.NewSource(time.Now().UnixNano())
	return rand.New(rSrc)
	// 全局rand会上锁,因此每个goroutine都用自己的rand
}

func SimpleIRSliceOfLForBench(targetLen int) ([]int, error) {
	if targetLen > maxIntRSliceLen {
		return nil, fmt.Errorf("too big target length %d (max %d)", targetLen, maxIntRSliceLen)
	}
	return generateRIntSliceOfLen(targetLen), nil
}
