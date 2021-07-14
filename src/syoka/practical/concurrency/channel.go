// Package concurrency
// channel 测试
package concurrency

import (
	"fmt"
)

//TestConsumerAndProvider 测试生产者,消费者通过channel传递消息
func TestConsumerAndProvider() {
	//create 字符串channel
	channel := make(chan string)
	go producer(channel)
	consumer(channel)
}

func consumer(channel chan string) {
	msg := <-channel
	fmt.Printf("consumer : accept data from channel:%+v", msg)
}

//数据上有
func producer(channel chan string) {
	channel <- "producer : produce hello world"
}

//读写
func TestReadWrite() {
	//测试读写
	channel := make(chan string)
	TestOnlyReader(channel)
}

//只写
func TestOnlyReader(channel <-chan string) {
	var _ = <-channel
	//channel <- "only read"//编译报错
}

//只写
func TestOnlyWriter(channel chan<- string) {
	channel <- "only write"
	//msg := <-channel //编译报错
}
