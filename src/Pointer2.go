package main

import "fmt"

//函数中传递对象只要是传递的指针则内部修改会对外部有影响
func main() {
	testcase1()
	println("-----------------")
	testcase2()
	println("-----------------")
	testcase3()
}

//相当于传入hobby的副本，interestingClub函数执行完毕不会对外部hobby有任何影响
func testcase1() {
	hobby := Hobby{"dancing"}
	println(">>> origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
	changeInteresting(hobby)
	//函数执行并不会影响外部对象
	println("<<< origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
}

//传入指针地址 ->&，内部修改对外部有影响
func testcase2() {
	hobby := Hobby{"dancing"}
	println(">>> origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
	changeInteresting2(&hobby)
	//函数执行并不会影响外部对象
	println("<<< origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
}

//传入指针地址 打印地址所指向的值，内部修改对外部有影响
func testcase3() {
	hobby := Hobby{"dancing"}
	println(">>> origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
	changeInteresting3(&hobby)
	//函数执行并不会影响外部对象
	println("<<< origin memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
}

func changeInteresting(hobby Hobby) {
	println(">>> function memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
	hobby.name = "singing"
	println("<<< function memory address:", &hobby, ",filed address:", &hobby.name, ",field val:", hobby.name)
}

func changeInteresting2(hobby *Hobby) {
	//对象地址都变了
	println(">>> function memory address:", &hobby, ",filed address:", &(hobby.name), ",field val:", hobby.name)
	hobby.name = "singing"
	println("<<< function memory address:", &hobby, ",filed address:", &(hobby.name), ",field val:", hobby.name)
}

func changeInteresting3(hobby *Hobby) {
	//打印对象值
	fmt.Println(">>> function memory address:", *hobby, ",filed address:", (*hobby).name, ",field val:", hobby.name)
	hobby.name = "singing"
	fmt.Println("<<< function memory address:", *hobby, ",filed address:", (*hobby).name, ",field val:", hobby.name)
}

type Hobby struct {
	name string
}
