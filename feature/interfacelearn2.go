package feature

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/21 23:43
// @description
// @version
func LearnInterface2() {
	rfp, err := newRandomFloat64Array(6)
	if err != nil {
		fmt.Println("new r float array error :", err)
		return
	}
	fmt.Println("before sort rfp:", rfp)
	sort.Sort(rfp)
	fmt.Println("after sort rfp:", rfp)
	fmt.Println("min in rfp:", min(rfp))

	// sort
	cnt := 100_0000
	ints := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		ints[i] = rand.Intn(1e12)
	}
	sort.Ints(ints)
	fmt.Println("sort ints[0]:", ints[0])
}

type float64RArray struct {
	data []float64
}

func (fap *float64RArray) Len() int {
	return len(fap.data)
}
func (fap *float64RArray) Less(i, j int) bool {
	return fap.data[i] < fap.data[j]
}
func (fap *float64RArray) Swap(i, j int) {
	d := fap.data
	d[i], d[j] = d[j], d[i]
}

func newRandomFloat64Array(cnt int) (*float64RArray, error) {
	if cnt >= 50 {
		return nil, fmt.Errorf("too big cnt:%d", cnt)
	}
	tn := time.Now().Nanosecond()
	rand.Seed(int64(tn))
	ar := &float64RArray{make([]float64, cnt)}
	for i := 0; i < cnt; i++ {
		ar.data[i] = round(rand.Float64(), 4)
	}
	return ar, nil
}
func (fap *float64RArray) list() string {
	return fmt.Sprintf("~float64RArray  : %v", fap.data)
}
func (fap *float64RArray) String() string {
	return fap.list()
}
func (fap *float64RArray) getEle(i int) interface{} {
	return fap.data[i]
}

// 保留小数位数的函数
func round(f float64, n int) float64 {
	p10n := math.Pow10(n)
	return math.Trunc(f*p10n+0.5) / p10n
}

type miner interface {
	sort.Interface
	getEle(i int) interface{}
}

func min(m miner) interface{} {
	if m.Len() <= 0 {
		return nil
	}
	mIx := 0
	l := m.Len()
	for i := 1; i < l; i++ {
		if m.Less(i, mIx) {
			mIx = i
		}
	}
	return m.getEle(mIx)
}
