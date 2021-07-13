package structure

import (
	"math/rand"
	"testing"
	"time"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Value  int
}

func TestCreatTree(t *testing.T) {
	nodes := createNode(1 << 4)
	root := &Node{nil, nil, nil, 8}
	for i := 0; i < len(nodes); i++ {
		insert(root, nodes[i])
	}
	println(root)
}

func createNode(number int) []int {
	var nodes []int
	rand.Seed(time.Now().Unix())
	for i := 0; i < number; i++ {
		nodes = append(nodes, rand.Intn(number))
	}
	return nodes
}

func insert(t *Node, val int) *Node {
	if t == nil {
		return &Node{nil, nil, nil, val}
	}
	if t.Value >= val {
		t.Left = insert(t.Left, val)
	} else {
		t.Right = insert(t.Right, val)
	}
	return t
}
