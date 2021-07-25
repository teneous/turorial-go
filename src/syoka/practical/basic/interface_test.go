package basic

import "testing"

func TestInterface(t *testing.T) {
	cow := Animal{name: "cow"}
	rosa := Plant{name: "rosa"}
	sayName(cow)
	sayName(rosa)
}

type Creature interface {
	//called name
	Name() string
	//is still living now
	Living() bool
}

type Animal struct {
	name string
}

type Plant struct {
	name string
}

func (a Animal) Name() string {
	return a.name
}

func (a Animal) Living() bool {
	return true
}

func (p Plant) Name() string {
	return p.name
}

func (p Plant) Living() bool {
	return true
}

func sayName(c Creature) {
	animal, ok := c.(Animal)
	if ok {
		println("c is animal ,named", animal.Name())
	}
	plant, ok := c.(Plant)
	if ok {
		println("c is plant ,named", plant.Name())
	}
}
