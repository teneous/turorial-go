package concurrency

import (
	"fmt"
	"testing"
)

//read from nil channel
//return fatal error: all goroutines are asleep - deadlock!
func TestReadFromNilChan(t *testing.T) {
	//define nil channel
	var dev chan string
	source := <-dev
	fmt.Printf("accept from nil channel %v", source)
}

//read from close chan
//return ""
func TestReadFromCloseChan(t *testing.T) {
	strChan := make(chan string, 10)
	close(strChan)
	source := <-strChan
	fmt.Printf("accept from close channel %v", source)
}

//close nil channel
//return panic: close of nil channel [recovered]
func TestCloseNilChan(t *testing.T) {
	//define nil chan
	var dev chan string
	close(dev)
}
