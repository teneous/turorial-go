package structure

import (
	"container/list"
	"fmt"
	"testing"
)

// 测试连标打印的正向和逆向顺序
func TestContainerList(t *testing.T) {
	data := prepareData()

	for t := data.Back(); t != nil; t = t.Prev() {
		fmt.Println(t.Value, " ")
	}

}

//prepareData 准备数据
func prepareData() *list.List {
	l := list.New()
	l.PushFront(3)
	l.PushFront(6)
	l.PushFront(2)
	l.PushFront(1)
	l.PushFront(4)

	fmt.Printf("front next%v \n", *l.Front())
	fmt.Printf("back pre%v \n", *l.Back())
	return l
}
