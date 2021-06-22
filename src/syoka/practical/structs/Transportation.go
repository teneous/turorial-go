package structs

import "fmt"

func main() {
	applewatch := AppleWatch{
		name:   "Apple Watch 6",
		weight: 12,
	}
	airplane := Airplane{
		name: "ANA",
	}
	//注意这里得传指针类型
	err := airplane.transport(&applewatch)
	if err != nil {
		return
	}
}

//传输
type Transportation interface {
	// 运输货物
	transport(cargo *Cargo) error
}

//货物
type Cargo interface {
	//货物名
	cargoName() string

	//重量
	cargoWeight() int
}

type Airplane struct {
	name string
}

func (airplane *Airplane) transport(cargo Cargo) error {
	fmt.Printf("运输公司:%s ，货物传输,货物名:%s,货物重量：%dt", airplane.name, cargo.cargoName(), cargo.cargoWeight())
	return nil
}

// Cargo的唯一实现
type AppleWatch struct {
	name   string
	weight int
}

func (watch *AppleWatch) cargoName() string {
	return watch.name
}

func (watch *AppleWatch) cargoWeight() int {
	return watch.weight
}
