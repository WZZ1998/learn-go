package feature

import (
	"fmt"
	"reflect"
	"unsafe"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/19 14:21
// @description
// @version
type myStruct struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type myStruct2 struct {
	ImportantName string
	int
	float64
	myStruct
	myStruct3
}
type myStruct3 struct {
	myStruct4
	//Name string 这里定义一个Name,就会和myStruct的Name形成同级字段
}
type myStruct4 struct {
	Name string
}

func NewMyStruct(name string, id int) *myStruct { // 这种遮蔽方法可以保证拿到实例必须走工厂
	if id < 0 {
		return nil
	}
	return &myStruct{name, id}
}
func LearnStruct2() {
	fmt.Println(unsafe.Sizeof(myStruct{})) // 24个字节,string16个字节,int 8个字节
	ms1 := NewMyStruct("Jennifer", 22)
	rType := reflect.TypeOf(*ms1)
	fmt.Println("rType: ", rType)
	for i := 0; i < 2; i++ {
		indexField := rType.Field(i)
		fmt.Printf("%v \n", indexField.Tag) // 打印了相应的
	}

	ms2Ins := &myStruct2{
		"Richard",
		0,
		3.3,
		myStruct{"Frank", 3},
		myStruct3{myStruct4{"Park"}},
	}
	ms2Ins.ImportantName = ms2Ins.ImportantName + " in myStruct2!"
	ms2Ins.myStruct3.Name = ms2Ins.myStruct3.Name + " in myStruct3"
	// 外层结构变量会遮蔽内层结构的变量,假如myStruct2 有一个name,那么ms2Ins.name访问的是外层的字段
	ms2Ins.Name = ms2Ins.Name + " which field to modify?"
	// 同级同名字段,不指名结构体,访问是无法通过编译的, 报ambiguous reference;
	// 多级同名变量,不指名结构体,优先访问靠外层的字段
	fmt.Printf("ms2Ins: %+v\n", ms2Ins)

}
