package lib

import (
	"fmt"
	"sort"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/26 00:32
// @description
// @version

type BlaDO struct {
	X, Y int
}

func (s *BlaDO) String() string {
	return fmt.Sprintf("BlaDO: X %v Y %v", s.X, s.Y)
}

type sortableBlaDos []*BlaDO

func (s sortableBlaDos) Len() int {
	return len(s)
}

func (s sortableBlaDos) Less(i, j int) bool {
	return s[i].X < s[j].X
}

func (s sortableBlaDos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortBlas(a []*BlaDO) {
	var ss sortableBlaDos
	ss = a
	sort.Sort(ss)
}
