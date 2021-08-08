// Package concurrency
package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

//TestGoThread 测试多协程
func TestGoThread(t *testing.T) {
	var group sync.WaitGroup
	//define state as 2
	group.Add(2)

	go calculate(&group)
	go func() {
		defer group.Done()
		//state release
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		println("anonymous func")
	}()

	group.Wait()
}

func TestGoWaitGroup(t *testing.T) {
	var group sync.WaitGroup

	for i := 0; i < 10; i++ {
		group.Add(1)
		time.Sleep(time.Second)
		go func() {
			defer group.Done()
			fmt.Println("go done")
		}()
	}
	group.Wait()
	fmt.Println("main finished")
}

func TestGoProc(t *testing.T) {
	print(runtime.GOMAXPROCS(1))
}

func calculate(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}
