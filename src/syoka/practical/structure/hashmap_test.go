//customized hashmap ref by java
//hash function is not a properly implementation
package structure

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

type Element struct {
	key      string
	val      string
	hashcode int64
	next     *Element
}

type HashMap struct {
	table []*Element
	size  int64
}

func TestHashMap(t *testing.T) {
	hashMap := NewHashMap(1 << 3)

	hello := NewElement("hello", "world")
	nihao := NewElement("nihao", " syoka")
	mergeHello := NewElement("hello", "bilibili")
	bonjour := NewElement("bonjour", " friends")

	hashMap.AddElement(hello)
	hashMap.AddElement(nihao)
	hashMap.AddElement(mergeHello)
	hashMap.AddElement(bonjour)

	value := hashMap.GetElement("hello")
	assert.True(t, strings.Compare(value, "bilibili") == 0, "hello should return bilibili")
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
	index := calculateIndex(element.hashcode, hashMap.size)
	tableRoot := table[index]
	if tableRoot != nil {
		findElementRecursively(tableRoot, element)
	} else {
		table[index] = element
	}
	return true
}

//GetElement get element by key
func (hashMap *HashMap) GetElement(key string) string {
	if hashMap.table == nil {
		panic("hashmap not initial")
	}
	hashCode := hash(key)
	index := calculateIndex(hashCode, hashMap.size)
	table := hashMap.table
	if table[index] == nil {
		//data not found
		return ""
	}
	return findElementByKey(table[index], key)
}

//NewHashMap construct
func NewHashMap(size int64) *HashMap {
	return &HashMap{
		//lazy initial
		table: nil,
		size:  size,
	}
}

//NewElement construct
func NewElement(key string, val string) *Element {
	hashcode := hash(key)
	return &Element{
		key:      key,
		val:      val,
		next:     nil,
		hashcode: hashcode,
	}
}

func findElementByKey(element *Element, key string) string {
	if element == nil {
		return ""
	} else {
		if element.key == key {
			return element.val
		} else {
			return findElementByKey(element.next, key)
		}
	}
}

func hash(key string) int64 {
	m := md5.New()
	m.Write([]byte(key))
	hexStr := hex.EncodeToString(m.Sum(nil))
	//prevent out of range
	newHex := hexStr[:10]
	hashcode, err := strconv.ParseInt(newHex, 16, 64)
	if err == nil {
		return hashcode
	}
	return -1
}

//calculateIndex calculate hash code
func calculateIndex(val int64, size int64) int64 {
	if size != 0 {
		return val % size
	}
	return 0
}

//findElementRecursively compare key then replace it when match
//when hashcode are the same,then add it at tail
func findElementRecursively(element *Element, target *Element) {
	if element.key == target.key {
		element.val = target.val
	} else {
		if element.next == nil {
			element.next = target
		} else {
			findElementRecursively(element.next, target)
		}
	}
}

func initTable(size int64) []*Element {
	elements := make([]*Element, size)
	return elements
}
