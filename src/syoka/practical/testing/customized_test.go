package testing

import (
	"fmt"
	"log"
	"testing"
)

func TestStringNotEmpty(t *testing.T) {
	a := 1
	b := 1
	c := 2
	defer log.Println("a:=", a)
	defer log.Println("b:=", b)
	defer log.Println("c:=", c)
	if a+b != c {
		log.Fatal("function execute equal fails")
		t.Fatal(fmt.Printf("a+b!=c.a:%d,b:%d,c:%d", a, b, c))
	} else {
		log.Println("function execute equal success")
	}
}
