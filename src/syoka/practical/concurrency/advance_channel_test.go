package concurrency

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var flag = true

func TestAdvanceChannel(t *testing.T) {
	cookChannel := make(chan string)
	transChannel := make(chan string)

	go cookerCookFoods(cookChannel)

	go waiter(cookChannel, transChannel)

	go customer(transChannel)

	time.Sleep(10 * time.Second)
	flag = false

}

//厨师创建食物
func cookerCookFoods(foods chan<- string) {
	for flag {
		n := rand.Int63n(32)
		fmt.Printf("厨师制作菜品. weight %d kg\n", n)
		foods <- "food" + strconv.FormatInt(n, 10)
		time.Sleep(1 * time.Second)
	}
}

//服务员将食物发放给顾客
func waiter(toWaiter <-chan string, toCus chan<- string) {
	for t := range toWaiter {
		if flag {
			fmt.Printf("服务员收到菜品，%s \n", t)
			toCus <- t
		}
	}
}

//顾客消费
func customer(foods <-chan string) {
	for f := range foods {
		if flag {
			fmt.Printf("客户消费菜品，%s \n", f)
		}
	}
}
