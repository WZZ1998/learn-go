package feature

import "fmt"

// @author  wzz_714105382@icloud.com
// @date  2020/6/4 02:01
// @description
// @version

func LearnSlice() {
	fmt.Println("modify with value & modify with pointer:")
	var a []int
	fmt.Println(a)
	modifyWithV(a)
	fmt.Println(a)
	modifyWithP(&a)
	fmt.Println(a)

}

func modifyWithV(s []int) []int {
	fmt.Println("modify with V")
	s = append(s, 1)
	return s
}
func modifyWithP(s *[]int) {
	fmt.Println("modify with P")
	*s = append(*s, 1)
}
