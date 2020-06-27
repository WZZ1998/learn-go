package feature

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

	uVc := new(VCard)
	uVc.Addresses = append(uVc.Addresses, nil)
	_ = json.Unmarshal(bs, uVc)
	// json列表元素,会对切片直接做重新分配,原来的切片变量直接被覆盖,指针和map元素也是类似的
	fmt.Println("uVc:", uVc)
	// 通用json解码类型,类似于java中的jsonObject
	// map[string]interface{} 还有 []interface{}

	jsSRdr := strings.NewReader(strings.Repeat(string(bs), 2000))
	jDecoder := json.NewDecoder(jsSRdr)
	// decoder拥有自己的缓冲buffer,目前看到大小是512;通过buffered方法可以拿到这个基于缓冲的reader
	fVC := new(VCard)
	err2 := jDecoder.Decode(fVC) // json反序列化
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("fVC", fVC)

	var gobWriteInSb strings.Builder
	gEnc := gob.NewEncoder(&gobWriteInSb)
	//gob.NewDecoder() 返回的decoder是自带buffer的; 可以用来解码gob数据
	fmt.Println("a1:", a1)
	err2 = gEnc.Encode(a1) // a1 是生成的一个Address数据
	if err2 != nil {
		panic(err2)
	}
	result := gobWriteInSb.String()
	fmt.Printf("raw gob vc data: %q \n", result)
	// 打印出来的东西有不少二进制编码
	gDec := gob.NewDecoder(strings.NewReader(result))

	uAd := new(Address)
	_ = gDec.Decode(uAd)
	fmt.Println("after decode uAd :", uAd)
	var f interface{}
	_ = json.Unmarshal(bs, &f)
	// 注意,这里必须是指针,虽然f已经是接口类型,但是必须传递&f,因为人家要修改f的值,传f还是值拷贝
	fmt.Println("f:", f)
	m := f.(map[string]interface{})
	for k, v := range m {
		fmt.Print("\t")
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Print("\t\t")
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}
	}
	resPath := "resources/"
	if runtime.GOOS == "darwin" {
		resPath = filepath.Join("/Users/wangzizhou/GolandProjects/learn-go/out", resPath)
	}
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
	JOutEncoder := json.NewEncoder(mtFW) // encoder没有自己的buffer!
	_ = JOutEncoder.Encode(vc)           // 注意写入规则,这里写的内容,在缓冲区剩下了一部分
	fmt.Println("# finish encoding call #")
}
