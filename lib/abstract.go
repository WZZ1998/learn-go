package lib

// @author  wzz_714105382@icloud.com
// @date  2020/10/27 22:27
// @description
// @version

type FullABAble interface {
	A() int
	B() int
	Sum() int
}

type simpleAB interface {
	A() int
	B() int
}

type absSumAbleC struct {
	simpleAB
	//FullABAble 千万不要这样写,如果忘了写Sum()会导致循环引用 cyclic reference,导致SO
	internalA, internalB int
}

func (p *absSumAbleC) Sum() int {
	return p.A() + p.B()
}

type concreteAbSum1 struct {
	absSumAbleC
}

func (s *concreteAbSum1) A() int {
	return s.internalA
}

func (s *concreteAbSum1) B() int {
	return s.internalB
}

type concreteAbSum2 struct {
	absSumAbleC
}

func NewConcreteAbSum1(insideA, insideB int) *concreteAbSum1 {
	//concrete := &concreteAbSum1{}
	//ac := absSumAbleC{
	//	simpleAB: concrete,
	//	internalA:  insideA,
	//	internalB:  insideB,
	//}
	//concrete.absSumAbleC = ac // 做的是值传递!, 这样写不好

	// 先初始化抽象类,再搞子类
	con := &concreteAbSum1{
		absSumAbleC{nil, insideA, insideB},
	}
	con.simpleAB = con
	return con
}
