package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type EntryNode struct {
	val  string
	next *EntryNode
}

//NewEntryNode 创建entryNode
func NewEntryNode(val string) *EntryNode {
	return &EntryNode{val, nil}
}

type LinkedList struct {
	head *EntryNode
	tail *EntryNode
	size int
}

//NewLinkedList constructor
func NewLinkedList() *LinkedList {
	return &LinkedList{nil, nil, 0}
}

//AddEntryNode
func (list *LinkedList) AddEntryNode(node *EntryNode) bool {
	if node == nil {
		return false
	}
	if list.size == 0 && list.head == nil {
		list.head, list.tail = node, node
		list.size += 1
		return true
	}
	oldTail := list.tail

	oldTail.next = node
	list.tail = node
	list.size++
	return true
}

//GetEntryNode find node by index
func (list *LinkedList) GetEntryNode(index int) *EntryNode {
	if index < 0 || list.size-1 < index {
		return nil
	}
	loopNode := list.head
	for i := 0; i < index; i++ {
		loopNode = loopNode.next
	}
	return loopNode
}

//GetSize 获取size
func (list *LinkedList) GetSize() int {
	return list.size
}

//RemoveEntryNode 删除Entry节点 a->b->c a->c
func (list *LinkedList) RemoveEntryNode(index int) bool {
	if index > 0 {
		index = index - 1
	}
	//find pre one
	cNode := list.GetEntryNode(index)
	if cNode.next != nil {
		cNode.next = cNode.next.next
		list.size--
		return true
	}
	return false
}

//test add() + size()
func TestAddAndSize(t *testing.T) {
	list := mockLinkedList()
	assert.True(t, list.size == 5, "list size must be 5.")
}

//test get() + remove()
func TestGetAndRemove(t *testing.T) {
	list := mockLinkedList()
	node := list.GetEntryNode(2)
	assert.True(t, node.val == "l", "node val must be l")

	list.RemoveEntryNode(2)
	list.RemoveEntryNode(3)
	assert.True(t, list.size == 3, "list size must be 3.")
	assert.True(t, list.head.val == "h", "node val must be h")
	assert.True(t, list.tail.val == "o", "node val must be o")

}

func mockLinkedList() *LinkedList {
	nodes := []*EntryNode{
		NewEntryNode("h"),
		NewEntryNode("e"),
		NewEntryNode("l"),
		NewEntryNode("l"),
		NewEntryNode("o")}

	list := NewLinkedList()
	for _, ele := range nodes {
		list.AddEntryNode(ele)
	}
	return list
}
