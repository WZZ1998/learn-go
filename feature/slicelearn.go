package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/4 02:01
// @description
// @version

func LearnSlice() {
	var a []int
	modifyWithV(a)
	fmt.Println(a)
	modifyWithP(&a)
	fmt.Println(a)

}

func modifyWithV(s []int) []int {
	s = append(s, 1)
	return s
}
func modifyWithP(s *[]int) {
	*s = append(*s, 1)
}
