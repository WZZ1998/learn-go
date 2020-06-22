package feature

import (
	"fmt"
	"reflect"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/22 05:14
// @description
// @version

type myStruct5 struct {
	Name  string
	id    int
	ratio float64
}

func (p *myStruct5) setId(id int) {
	p.id = id
}
func (p *myStruct5) String() string {
	return fmt.Sprintf("myStruct5:%s - %d - %.4f", p.Name, p.id, p.ratio)
}
func LearnReflect() {
	var x = 3.1415
	ty := reflect.TypeOf(x)
	va := reflect.ValueOf(x)
	fmt.Printf("reflect type %v reflect value %v\n", ty, va)
	fmt.Println("reflect ty type:", reflect.TypeOf(ty)) // 这是个reflect包的私有结构,人家开了个接口出来
	fmt.Printf("ty constant kind: %d %v == reflect.Float64 %t\n",
		ty.Kind(), ty.Kind(), ty.Kind() == reflect.Float64)
	// va.Type() 方法返回va的类型,返回类型为Type
	// 这个方法返回空接口类型的va底层值
	fmt.Printf("va Interface call : %v convert to float64 %f \n", va.Interface(), va.Interface().(float64))
	// fmt.Printf("va get int value: %d", va.Int()) 直接panic,在value上调用错误的类型取值方法

	// 修改值
	// va.SetFloat(6.28) 不行,直接panic, 做反射的时候进行了值拷贝,无法改变原来的值
	fmt.Printf("va can set : %t \n", va.CanSet())

	pva := reflect.ValueOf(&x) // 这个地方必须用指针,不用指针无法避免值传递
	fmt.Printf("pva type: %v can set %t\n", pva.Type(), pva.CanSet())
	pva = pva.Elem() // 发生了什么? 这是迂回地获取了原来值的反射对象,因为golang只有值传递,所以必须先绕道到指针,再转回来
	fmt.Printf("pva type: %v can set %t\n", pva.Type(), pva.CanSet())
	pva.SetFloat(6.28)
	fmt.Println("x : ", x)

	ms5 := &myStruct5{"Wa", 1, 3.00}
	sec := interface{}(ms5)
	ty2 := reflect.TypeOf(sec)
	va2 := reflect.ValueOf(sec)

	fmt.Printf("ty2 :%v kind: %v\n", ty2, ty2.Kind())
	fmt.Println("get ele for va2")
	va2 = va2.Elem()
	fmt.Printf("va2 %v num of fields %d .Type().NumField() %d\n",
		va2, va2.NumField(), va2.Type().NumField())
	for i, fn := 0, va2.NumField(); i < fn; i++ {
		fn := va2.Type().Field(i).Name //要想看到字段的名字,必须反射出类型
		vf := va2.Field(i)
		fmt.Printf("field %d :%s %v\n", i, fn, vf)
	}
	vf := va2.Field(0) // 只有导出的字段才能被反射修改
	vf.SetString("changed! Ha!")
	fmt.Println("va2 :", va2)

	// 再次注意,绑定到指针和绑定到值差很多呀,反射也不能接受
	va2 = reflect.ValueOf(sec) // va2这样是指针的反射类型
	res := va2.Method(0).Call(nil)
	fmt.Println("res:", res)

	var im IM = &myStruct6{"myStruct6 topped:", &myStruct5{}}
	fmt.Println("call m1 on m6:", im.m1())

}

type IM interface {
	m1() string
}

func (p *myStruct5) m1() string {
	return "yes, myStruct5 implements interface IM"
}

type myStruct6 struct {
	ap string
	IM
}

func (p *myStruct6) m1() string {
	return p.ap + p.IM.m1()
}
