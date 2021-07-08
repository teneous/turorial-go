//slice demo
//in my opinion ,slice equals to arrayList
package basic

import (
	"fmt"
	"reflect"
	"testing"
)

//slice
func TestAnnounce(t *testing.T) {
	//slice
	//no length size announced
	sliceA := []int{1, 2, 3, 4, 5}
	fmt.Printf("sliceA type is :%v \n", reflect.TypeOf(sliceA))

	//array
	//announce size
	sliceB := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("sliceB type is :%v \n", reflect.TypeOf(sliceB))

	//slice
	//make create
	sliceC := make([]int, 23)
	fmt.Printf("sliceC type is :%v \n", reflect.TypeOf(sliceC))
}

//slice expand with double
func TestSliceExpand(t *testing.T) {
	slices := make([]int, 5)
	//len:0 cap:5
	fmt.Printf("origin len:%d,cap:%d \n", len(slices), cap(slices))

	//len:5 cap:5 * 2 = 10
	slices = append(slices, 2)
	fmt.Printf("origin len:%d,cap:%d \n", len(slices), cap(slices))
}

//array expand with double
func TestSliceSubList(t *testing.T) {
	slices := make([]int, 5)
	//0 0 0 0 0
	fmt.Printf("origin %+v \n", slices)

	//0
	sub := slices[2:3]
	fmt.Printf("sub %+v \n", sub)
	//memory address are the same,which mean when you change sub[0]'s value,slices will change too
	fmt.Printf("slices[2] addr:%v sub[0] addr:%v \n", &slices[2], &sub[0])

	sub[0] = -1
	fmt.Printf("origin %+v , sub %+v \n", slices, sub)
}
