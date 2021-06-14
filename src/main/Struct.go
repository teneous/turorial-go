package main

import (
	"fmt"
)

func main() {
	car := newCar("benz", 400_000, true)
	fmt.Printf("%#v", car)
}

type Car struct {
	brand  string
	price  int
	luxury bool
}

func newCar(brand string, price int, luxury bool) Car {
	car := Car{
		brand:  "green bow",
		price:  22_000,
		luxury: false,
	}
	car.brand = brand
	car.price = price
	if luxury {
		car.luxury = luxury
	}
	return car
}
