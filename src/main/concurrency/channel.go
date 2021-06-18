package main

import (
	"fmt"
)

func main() {
	testConsumerAndProvider()
}

func testConsumerAndProvider() {
	//通道接受字符串类型
	channel := make(chan string)
	go producer(channel)
	consumer(channel)
}

func testReadWrite() {
	//测试读写
	channel := make(chan string)
	onlyWriter(channel)

}

func consumer(channel chan string) {
	msg := <-channel
	fmt.Printf("accept data from channel:%+v", msg)
}

//数据上有
func producer(channel chan string) {
	channel <- "hello world"
}

//只写
func onlyReader(channel <-chan string) {
	var _ = <-channel
	//channel <- "only read"//编译报错
}

//只写
func onlyWriter(channel chan<- string) {
	channel <- "only write"
	//msg := <-channel //编译报错
}
