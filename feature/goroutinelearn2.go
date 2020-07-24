package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/7/21 00:02
// @description
// @version
func LearnGoroutine2() {
	fmt.Println(len(lcLookAndSay(20)))
}

func lcLookAndSay(c int) string {
	a := make([]byte, 0)
	b := make([]byte, 0)
	a = append(a, '1')
	for i := 0; i < c-1; i++ {
		ll := len(a)
		var currentB = a[0]
		var currentCounter byte = 1
		for j := 1; j < ll; j++ {
			if a[j] == currentB {
				currentCounter++
			} else {
				b = append(b, currentCounter+'0')
				b = append(b, currentB)
				currentCounter = 1
				currentB = a[j]
			}
		}
		b = append(b, currentCounter+'0')
		b = append(b, currentB)
		a = a[:0]
		a = append(a, b...)
		b = b[:0]
	}
	return string(a)
}
