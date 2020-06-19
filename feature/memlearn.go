package feature

import (
	"fmt"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/19 00:09
// @description
// @version

type testS struct {
	fie [5]int
}

func LearnMem() {
	//literalParamResearch()
	//retAsParamResearch()
	retStructAsParamResearch()
	//literalStructParamResearch()
}

func literalParamResearch() {
	testLP(0xeeeeeee, [5]int{1000, 1000, 1000, 1000, 1000})
}
func retAsParamResearch() {
	testRP2(testRP1())
}

func retStructAsParamResearch() {
	//testRetSP2(testRetSP1())
	testRetSP2(testRetSP1NoRetName())

}
func literalStructParamResearch() {
	testRetSP2(testS{})
}
func testLP(qqq int, arr [5]int) {
	qqq = 0xfffffff
	arr[2] = -999
	qqqAddr := getAddrValueFromPointer(&qqq)
	fmt.Printf("in testLP: para qqq addr: %x\n", qqqAddr)
	arrParamPtrV := uintptr(unsafe.Pointer(&arr))
	var bias uintptr = 5 * 8
	fmt.Printf("in testLP: arg array addr :%x\n", arrParamPtrV)
	fmt.Println("dump testLP mem:")
	dumpMemoryWithUnsafe(arrParamPtrV-bias, arrParamPtrV+bias, 8)
}

func testRP2(ret int) {
	pt := &ret
	ptv := getAddrValueFromPointer(pt) //防止逃逸
	fmt.Printf("in testRP2 arg ret addr %x\n", ptv)
	var bias uintptr = 16 * 8
	dumpMemoryWithUnsafe(ptv-bias, ptv+bias, 8)
}

func testRP1() (ret int) {
	a := 15
	b := a
	ret = 0xccccc
	println("in testRP1 ret addr:", &ret)
	println("in test RP1 a addr", &a, "b addr", &b)
	return
}

func testRetSP1() (rs testS) {
	rs = testS{[5]int{0xdd, 0xdd, 0xdd, 0xdd, 0xdd}}
	println("in testRetSP1 return rs addr :", &rs)
	return
}

func testRetSP1NoRetName() testS {
	return testS{[5]int{0xdd, 0xdd, 0xdd, 0xdd, 0xdd}}
}

func testRetSP2(rs testS) {
	println("in testRetSP2 arg rs addr: ", &rs)
	sPtrV := uintptr(unsafe.Pointer(&rs))
	var bias uintptr = 10 * 8
	dumpMemoryWithUnsafe(sPtrV-bias, sPtrV+bias, 8)
}

func dumpMemoryWithUnsafe(start, end, step uintptr) {
	for addrV := end; addrV >= start; addrV -= step {
		xv := *(getIntPointerFromAddrValue(addrV))
		fmt.Printf("addr v [%x] deref int %15d hexi %x\n", addrV, xv, xv)
	}
}
func getAddrValueFromPointer(p *int) uintptr {
	return uintptr(unsafe.Pointer(p))
}
func getIntPointerFromAddrValue(av uintptr) *int {
	return (*int)(unsafe.Pointer(av))
}
