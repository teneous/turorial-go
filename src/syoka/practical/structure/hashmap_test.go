package structure

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"testing"
)

type Element struct {
	hashcode int64
	val      string
	next     *Element
}

type HashMap struct {
	table []*Element
	size  int64
}

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap()
	element := NewElement("hello world")
	element2 := NewElement("hello bilibili")

	hashMap.AddElement(element)
	hashMap.AddElement(element2)
}

//AddElement method add
func (hashMap *HashMap) AddElement(element *Element) bool {
	if element == nil {
		return false
	}
	if hashMap.table == nil {
		hashMap.table = initTable(hashMap.size)
	}
	table := hashMap.table
	index := calculateHash(element.hashcode, hashMap.size)
	root := table[index]
	if root != nil {
		deep := findNextElementRecursively(root)
		deep.next = element
	} else {
		table[index] = element
	}
	return true
}

//NewHashMap construct
func NewHashMap() *HashMap {
	return &HashMap{
		table: nil,
		size:  16,
	}
}

//NewElement construct
func NewElement(val string) *Element {
	m := md5.New()
	m.Write([]byte(val))
	hexStr := hex.EncodeToString(m.Sum(nil))
	//prevent out of range
	newHex := hexStr[:10]
	hashcode, err := strconv.ParseInt(newHex, 16, 64)
	if err == nil {
		return &Element{
			next:     nil,
			val:      val,
			hashcode: hashcode,
		}
	}
	return nil
}

//calculateHash calculate hash code
func calculateHash(val int64, size int64) int64 {
	if size != 0 {
		return val % size
	}
	return 0
}

func findNextElementRecursively(element *Element) *Element {
	if element.next == nil {
		return element
	} else {
		return findNextElementRecursively(element.next)
	}
}

func initTable(size int64) []*Element {
	elements := make([]*Element, size)
	return elements
}
