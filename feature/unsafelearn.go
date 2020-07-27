package feature

import (
	"fmt"
	"runtime"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/26 21:03
// @description
// @version

func LearnUnsafe() {
	byteS := []byte{'h', 'e', 'l', 'l', 'o'}
	sss := string(byteS)
	sssUnsafe := *(*string)(unsafe.Pointer(&byteS))
	byteS[0] = '^'

	bytesFromStr := *(*[]byte)(unsafe.Pointer(&sss))
	fmt.Println("byteS:", byteS)
	fmt.Println("sss:", sss)
	fmt.Println("sssUnsafe:", sssUnsafe)
	fmt.Println("bytesFromStr:", bytesFromStr)
	bytesFromStr[0] = '%'
	fmt.Println("bytesFromStr:", bytesFromStr)

	sl := unsafeGetBytes()
	runtime.GC()
	fmt.Println(sl)
}

func unsafeGetBytes() []byte {
	s := "dkadlkdl;akw"
	return *(*[]byte)(unsafe.Pointer(&s))
}
