// Package concurrency
package concurrency

import (
	"fmt"
	"sync"
	"testing"
)

//TestGoThread 测试多协程
func TestGoThread(t *testing.T) {
	var group sync.WaitGroup
	//define state as 2
	group.Add(2)
	group.Add(2)

	go calculate(&group)
	go func() {
		defer group.Done()
		//state release
		for i := 0; i < 10; i++ {
			fmt.Printf("%d", i)
		}
		println("anonymous func")
	}()

	group.Wait()
}

func calculate(group *sync.WaitGroup) {
	defer group.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
}
