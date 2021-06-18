package main

import (
	"time"
)

func main() {
	doSomeThing()
}

func doSomeThing() {
	go doSing()
	go doDancing()
	println("working")
	//main thread block
	time.Sleep(time.Second * 5)
}

//run new thread
func doSing() {
	time.Sleep(time.Second * 1)
	println("sings")
}

//run new thread
func doDancing() {
	time.Sleep(time.Second * 1)

	println("dancing")
}
