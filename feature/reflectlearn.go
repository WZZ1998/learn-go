package feature

import (
	"fmt"
	"reflect"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/22 05:14
// @description
// @version
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

}
