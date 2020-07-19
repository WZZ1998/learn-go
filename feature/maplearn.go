package feature

import (
	"fmt"
	"math/rand"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/16 02:35
// @description
// @version

func LearnMap() {
	m1 := map[string]string{}
	m1["Jan"] = "一月"
	m2 := m1
	fmt.Println("map m2:", m2)

	mf := map[int]func() int{
		1: func() int { return 100 },
		2: func() int { return 200 },
		3: func() int { return 300 },
	}
	fmt.Println("mf: ", mf) //将整数映射到函数地址

	mBig := make(map[string]int, 1024) // 指定容量,推荐
	fmt.Println("mBig", mBig)

	msl := map[string][]int{"A": {1, 2, 3}, "B": {100, 200, 300}, "C": {1000, 2000, 3000}}
	slA := msl["A"]
	slA = append(slA, 4)
	fmt.Printf("slA value %v addr %p msl[\"A\"] value %v\n", slA, &slA, msl["A"]) // 注意,这是个天大的陷阱!
	msl["A"] = append(msl["A"], 5)
	fmt.Printf("after direct append msl[\"A\"] value %v\n", msl["A"])

	msp := map[string]*[]int{"A": {1, 2, 3}, "B": {100, 200, 300}, "C": {1000, 2000, 3000}}
	slpA := msp["A"]
	*slpA = append(*slpA, 4)
	fmt.Printf("slpA %v content %v msp[\"A\"] %v content %v\n", slpA, *slpA, msp["A"], *msp["A"])

	si := 2048
	inm := make(map[int]int, si)
	for i := 1; i <= si; i++ {
		inm[i] = rand.Intn(8192)
	}
	a1, ok1 := inm[3000]
	a2, ok2 := inm[2000]
	fmt.Printf("inm: 3000 in map: %t get value %d    2000 in map :%t value %d\n", ok1, a1, ok2, a2)
	delete(inm, 2000)
	a1, ok1 = inm[3000]
	a2, ok2 = inm[2000]
	// ppp := &inm[2000] 无法获得map元素的地址
	fmt.Printf("after delete inm: 3000 in map: %t get value %d    2000 in map :%t value %d\n", ok1, a1, ok2, a2)
	for k, v := range inm {
		if k%500 == 0 {
			fmt.Printf("key = %d value = %d\n", k, v)
		}
	}

}
