//多路复用
package io

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	channel1 := make(chan int, 1)
	channel2 := make(chan bool)
	c := 0

	go GeneratorNumber(channel1, channel2)

	for {
		select {
		case c = <-channel1:
			fmt.Println("Receive1", c)
		case c = <-channel1:
			fmt.Println("Receive2", c)
		case <-channel2:
			fmt.Println("interrupt")
			return
		}
	}
}

func TestSelectAdvance(t *testing.T) {
	//select 执行超时，after中断
	InterruptSelect(2000)
}

func TestSelectAdvance2(t *testing.T) {
	//select 正常执行
	InterruptSelect(500)
}

func GeneratorNumber(ch chan int, stopCh chan bool) {
	time.Sleep(1 * time.Second)
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(1 * time.Second)
	}
	stopCh <- true
}

func InterruptSelect(sleepTime int64) {
	channel := make(chan string)
	go func() {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		channel <- "ok"
	}()
	select {
	case str := <-channel:
		fmt.Println(str)
	case <-time.After(1 * time.Second):
		fmt.Println("ops timeout")
	}
}
