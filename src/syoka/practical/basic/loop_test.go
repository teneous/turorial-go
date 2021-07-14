//for , foreach test case
package basic

import "testing"

//test normal
func TestFor(t *testing.T) {
	var total = 0
	for i := 0; i < 10_000; i++ {
		total += i
	}
	println("total val :", total)
}

//test while
func TestInfiniteLoop(t *testing.T) {
	for {
		println("infinite basic")
	}
}

//test range
func TestRangeLoop(t *testing.T) {
	sliceArray := []int{1, 2, 3, 4, 5}
	for idx, val := range sliceArray {
		println("idx:", idx, "val:", val)
	}
}
