package feature

import (
	"fmt"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/7/26 21:03
// @description
// @version

func LearnUnsafe() {
	byteS := []byte{'h', 'e', 'l', 'l', 'o'}
	sss := string(byteS)
	sssUnsafe := *(*string)(unsafe.Pointer(&byteS)) // 强制转换
	byteS[0] = '^'

	fmt.Println("byteS:", byteS)
	fmt.Println("sss:", sss)
	fmt.Println("sssUnsafe:", sssUnsafe)

	// sss 本来是个string
	bytesFromStr := *(*[]byte)(unsafe.Pointer(&sss))
	fmt.Println("bytesFromStr:", bytesFromStr)
	bytesFromStr[0] = '%'
	fmt.Println("After modify, bytesFromStr:", bytesFromStr)
	fmt.Println("now sss:", sss)
	// 这里发现,本来不可变的字符串被修改了,并没有panic或者报错
}
