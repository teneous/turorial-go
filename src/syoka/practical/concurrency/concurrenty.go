// Package concurrency
// 并发 测试
package concurrency

import (
	"time"
)

//TestGoThread 测试多线程窒息
func TestGoThread() {
	go doSing()
	go doDancing()
	println("working")
	//main thread block
	time.Sleep(time.Second * 5)
}

// thread a sing
func doSing() {
	time.Sleep(time.Second * 1)
	println("sings")
}

// thread b dancing
func doDancing() {
	time.Sleep(time.Second * 1)
	println("dancing")
}
