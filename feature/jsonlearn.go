package feature

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/27 02:06
// @description
// @version

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func LearnJson() {
	sr := strings.NewReader("123456789abcdefghijklmnopqrstuvwxyz")
	sr2 := strings.NewReader("ABCDEFGHIJKLMNOPQRST")
	bsr := bufio.NewReaderSize(sr, 16)
	da := make([]byte, 15)
	n, _ := bsr.Read(da)
	fmt.Println("read 1 :", string(da[:n]))
	n, _ = bsr.Read(da)
	fmt.Println("read 2 :", string(da[:n]))
	n, _ = bsr.Read(da)
	fmt.Println("read 3 :", string(da[:n]))
	// 在这之后,缓冲区里面还有个w
	bsr.Reset(sr2) // 重用之后,缓存中还没读的数据被干掉了
	n, _ = bsr.Read(da)
	fmt.Println("read after reset :", string(da[:n]))

	a1 := &Address{"private", "Beijing", "China"}
	a2 := &Address{"work", "Shanghai", "China"}
	vc := &VCard{"Beck", "Sartre", []*Address{a1, a2}, "Great"}
	bs, err := json.Marshal(vc)
	if err != nil {
		panic(err)
	}
	fmt.Println("json format vc:", string(bs))

	resPath := "resources/"

	f1Name := "vcard.json"
	f1, err2 := os.OpenFile(filepath.Join(resPath, f1Name),
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err2 != nil {
		panic(err2)
	}
	defer f1.Close()

	jEncoder := json.NewEncoder(f1)
	err3 := jEncoder.Encode(vc)
	if err3 != nil {
		panic(err3)
	}

	mtFW := bufio.NewWriterSize(os.Stdout, 94)
	defer mtFW.Flush()
	mtFW.WriteString(strings.Repeat("a", 10))
	JOutEncoder := json.NewEncoder(mtFW)
	_ = JOutEncoder.Encode(vc) // 注意写入规则,这里写的内容,在缓冲区剩下了一部分
	fmt.Println("# finish encoding call #")

}
