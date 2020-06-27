package feature

import (
	"fmt"
	"math/rand"
	"sort"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/22 23:23
// @description
// @version
var brands []string
var brandToModels map[string][]string
var idr int

func LearnSimpleFP() {
	prepareBrandsAndModels()
	var carSl []*Car
	cCnt := 40
	for i := 0; i < cCnt; i++ {
		carSl = append(carSl, NewCar(getRandomBrandModelDateID()))
	}
	myCs := Cars(carSl)
	oldAudiCars := myCs.FindAll(func(c *Car) bool {
		return c.brand == "Audi" && c.productionDate < 2007
	})
	fmt.Println("old Audi:", oldAudiCars)
	apd, groupedCars := MakeGroupedAppender()
	myCs.Process(apd)
	teslaCars := groupedCars["Tesla"]
	nn := len(teslaCars)
	// 做一下内存拷贝,防止把原来的东西弄乱了
	teslaCars = append(make([]*Car, 0, nn), teslaCars...)
	fmt.Printf("we have %d Tesla cars\n", len(teslaCars))
	sort.Slice(teslaCars, func(i, j int) bool {
		// 从新到旧排
		return !(teslaCars[i].productionDate < teslaCars[j].productionDate)
	})
	if len(teslaCars) > 0 {
		fmt.Println("the new one:", teslaCars[0])
	}

}

func prepareBrandsAndModels() {
	brands = []string{"Benz", "BMW", "Audi", "Tesla", "BYD", "HongQi", "Toyota"}
	brandToModels = make(map[string][]string)
	brandToModels["Benz"] = []string{"S100", "S200", "S300", "S500", "S600"}
	brandToModels["BMW"] = []string{"X3", "X5", "X7", "Series-3", "Series-5", "Series-7", "Series-9"}
	brandToModels["Audi"] = []string{"Q5", "Q7", "A3", "A4", "A4L", "A6", "A6L", "A8"}
	brandToModels["Tesla"] = []string{"model-A", "model-S", "RoadStar"}
	brandToModels["BYD"] = []string{"B1", "B2", "B45", "B60", "B95", "Ele-Song"}
	brandToModels["HongQi"] = []string{"Hong1", "Hong2", "Hong3", "Hong7", "Classical-HongQi"}
	brandToModels["Toyota"] = []string{"Camry", "Carola", "SUV", "Super-SUV"}
}

func getRandomBrandModelDateID() (model, brand string, prodDate, id int) {
	brandsL := len(brands)
	brandIx := rand.Intn(brandsL)
	brand = brands[brandIx]
	ms := brandToModels[brand]
	modelsL := len(ms)
	modelIx := rand.Intn(modelsL)
	model = ms[modelIx]
	prodDate = 1999 + rand.Intn(22)
	id = getIdNow()
	return
}

type Car struct {
	id             int
	model          string
	brand          string
	productionDate int
}

func (cp *Car) String() string {
	return fmt.Sprintf("[Car: %d | %s | %s | %d]", cp.id, cp.brand, cp.model, cp.productionDate)
}
func NewCar(model, brand string, prodDate, id int) *Car {
	return &Car{model: model, brand: brand, productionDate: prodDate, id: id}
}

type Cars []*Car

func (cs Cars) Process(f func(c *Car)) {
	for _, c := range cs {
		f(c)
	}
}

func (cs Cars) FindAll(f func(c *Car) bool) Cars {
	cars := Cars{}
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

func MakeGroupedAppender() (func(c *Car), map[string]Cars) {
	groupedCars := make(map[string]Cars)
	for _, br := range brands {
		groupedCars[br] = Cars{}
	}
	groupedCars["unknown"] = Cars{}
	appender := func(c *Car) {
		if _, in := groupedCars[c.brand]; in {
			groupedCars[c.brand] = append(groupedCars[c.brand], c)
		} else {
			groupedCars["unknown"] = append(groupedCars["unknown"], c)
		}
	}
	return appender, groupedCars
}
func getIdNow() int {
	old := idr
	idr++
	return old
}
