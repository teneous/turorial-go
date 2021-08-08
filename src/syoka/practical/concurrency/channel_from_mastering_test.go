package concurrency

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	channel := make(chan int, 5)
	go sumData(channel)
	go sendData(channel)
	time.Sleep(time.Duration(3) * time.Second)
}

func sendData(channel chan<- int) {
	//dead iterate
	for {
		channel <- rand.Intn(5)
	}
}

func sumData(channel <-chan int) {
	//2 secdone timer
	timer := time.NewTimer(time.Duration(2) * time.Second)
	sum := 0
	for {
		select {
		case t := <-channel:
			sum += t
		case <-timer.C:
			channel = nil
			fmt.Println("sum val ", sum)
		}

	}
}
