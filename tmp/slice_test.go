package tmp

import (
	"fmt"
	"testing"
	"unsafe"
)

//type SliceItem struct {
//	val     int32
//	trigger bool
//	text    string
//}

func TestSlice(t *testing.T) {
	arr := [16][]int{}

	//a := arr[6:]
	//a = append(a, make([]int, 0), make([]int, 0), make([]int, 0))
	//a = append(a, make([]int, 0), make([]int, 0))
	//
	//fmt.Printf("%+v\n", unsafe.Pointer(&a))
	fmt.Printf("%+v", unsafe.Pointer(&arr))
}
