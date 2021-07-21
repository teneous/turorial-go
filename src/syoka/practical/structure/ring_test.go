package structure

import (
	"container/ring"
	"testing"
)

//测试ring
func TestRing(t *testing.T) {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	sum := 0
	r.Do(func(i interface{}) {
		t := i.(int)
		sum += t
	})

	println(r)
	println(sum)
}

func swap() {

}
