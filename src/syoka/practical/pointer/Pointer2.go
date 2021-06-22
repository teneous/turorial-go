package pointer

//函数中传递对象只要是传递的指针则内部修改会对外部有影响
func main() {
	println("-------case1----------")
	testcase1()
	println("-------case2----------")
	testcase2()
}

//值传递-传递副本
//conclusion: 通过对比指针地址，我们可以看到传递的是对象副本信息
func testcase1() {
	hobby := Hobby{"dancing"}
	println("外部>>> 执行前-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
	changeInteresting(hobby)
	println("外部>>> 执行后-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
}

//值传递-传递的是指针副本
//conclusion: 对象地址指针出现了拷贝，但是对象内部属性指向的地址是一样的
func testcase2() {
	hobby := Hobby{"dancing"}
	println("外部>>> 执行前-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
	changeInteresting2(&hobby)
	println("外部>>> 执行后-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
}

//修改兴趣
func changeInteresting(hobby Hobby) {
	//对象内存地址发生了变化，说明是拷贝了一份对象
	println("内部>>> 执行前-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
	hobby.name = "singing"
	println("内部<<< 执行后-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
}

func changeInteresting2(hobby *Hobby) {
	println("内部>>> 执行前-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
	hobby.name = "singing"
	println("内部<<< 执行后-对象指针地址", &hobby, ",对象属性名指针地址", &hobby.name, ",属性值:", hobby.name)
}

//兴趣
type Hobby struct {
	name string
}
