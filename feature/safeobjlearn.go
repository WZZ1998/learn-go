package feature

import (
	"fmt"
	"sync"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/8/20 21:07
// @description
// @version

type gSafePerson struct {
	Name   string
	salary float64
	opCh   chan func()
}

func (p *gSafePerson) SetSalary(salary float64) (setErr error) {
	defer func() {
		setErr = formatRecoverRetAsErr(recover(), "obj already deleted")
	}()
	p.opCh <- func() {
		p.salary = salary
	}
	return
}

func (p *gSafePerson) SetSalaryAdd(delta float64) (addErr error) {
	defer func() {
		addErr = formatRecoverRetAsErr(recover(), "obj already deleted")
	}()
	p.opCh <- func() {
		p.salary = p.salary + delta
	}
	return
}
func (p *gSafePerson) Salary() (ret float64, salaryErr error) {
	defer func() {
		salaryErr = formatRecoverRetAsErr(recover(), "obj already deleted")
	}()
	salaryCh := make(chan float64)
	p.opCh <- func() {
		salaryCh <- p.salary
	}
	ret = <-salaryCh
	return
}
func NewGSafePerson(name string, salary float64) *gSafePerson {
	ob := &gSafePerson{
		Name:   name,
		salary: salary,
		opCh:   make(chan func()),
	}
	go ob.background()
	return ob
}

func (p *gSafePerson) background() {
	for opF := range p.opCh {
		opF()
	}
}

func (p *gSafePerson) Delete() (deleteErr error) {
	defer func() {
		deleteErr = formatRecoverRetAsErr(recover(), "obj already deleted")
	}()
	close(p.opCh)
	return
}

type AA struct {
	name string
	age  int
}

func LearnSafeObj() {
	//ma := AA{"ori", 10}
	//var x interface{} = ma
	//ma.name = "mmm"
	//ma.age = 999
	//fmt.Printf("ma %v x %v", &ma, x) // 值类型赋值给接口时,确实是会做值拷贝,所以要慎重
	for i := 0; i < 200; i++ {
		time.Sleep(2 * time.Millisecond)
		tryCreate()
	}
	gPerson := NewGSafePerson("wang", 0.0)
	defer gPerson.Delete()

	gCh := make(chan bool)
	gCnt := 3721
	wg := new(sync.WaitGroup)
	wg.Add(gCnt)

	for i := 0; i < gCnt; i++ {
		go func() {
			<-gCh
			_ = gPerson.SetSalaryAdd(1.0)
			wg.Done()
		}()
	}
	for i := 0; i < gCnt; i++ {
		gCh <- true
	}
	wg.Wait()
	sal, _ := gPerson.Salary()
	fmt.Println("after concurrent modify, salary:", sal)
}

func formatRecoverRetAsErr(re interface{}, desc string) error {
	if re != nil {
		return fmt.Errorf("err:%v | desc:%s", re, desc)
	}
	return nil
}

func tryCreate() {

	const pCnt = 1e3
	ps := make([]*gSafePerson, 0, pCnt)
	for i := 0; i < pCnt; i++ {
		pp := NewGSafePerson("Jennifer", 1000)
		ps = append(ps, pp)
	}
	time.Sleep(1 * time.Millisecond)
	for _, pp := range ps {
		_ = pp.Delete()
	}
}
