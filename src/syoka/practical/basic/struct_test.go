package basic

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	//return pointer address
	metadata := new(ClassMetadata)
	fmt.Printf("memory address->%p,val->%v", metadata, *metadata)
}

type ClassMetadata struct {
	className string
	modifier  int
	namespace string
}
