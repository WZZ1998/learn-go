package feature

import (
	"fmt"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/22 17:23
// @description 用golang的方法实现一下装饰者模式,远远没有java那么麻烦了
// @version
func LearnReflect2() {
	myGood :=
		NewStrawberryTopped(
			NewCheeseTopped(
				NewStrawberryTopped(
					NewCake())))
	fmt.Println("my good OK!", myGood.GetDes(), myGood.GetCost())

}

// 共同接口
type Sweet interface {
	GetCost() float32
	GetDes() string
}

// 被增强类
type cake struct{}

func (p *cake) GetCost() float32 {
	return 45.5
}
func (p *cake) GetDes() string {
	return "one cake"
}

func NewCake() Sweet {
	return &cake{}
}

// 点睛之笔,内嵌Sweet后,strawberry直接实现了Sweet接口,所以不需要继承那一套了
// 内嵌太强了
type strawberryTopped struct {
	Sweet
}

func NewStrawberryTopped(s Sweet) Sweet {
	return &strawberryTopped{s}
}
func (p *strawberryTopped) GetCost() float32 {
	return 5.6 + p.Sweet.GetCost()
}
func (p *strawberryTopped) GetDes() string {
	return p.Sweet.GetDes() + ", add strawberry"
}

type cheeseTopped struct {
	Sweet
}

func NewCheeseTopped(s Sweet) Sweet {
	return &cheeseTopped{s}
}
func (p *cheeseTopped) GetCost() float32 {
	return 7.9 + p.Sweet.GetCost()
}
func (p *cheeseTopped) GetDes() string {
	return p.Sweet.GetDes() + ", add cheese"
}
