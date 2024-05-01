package sort

import (
	"fmt"
	"testing"
)

func TestQuictSort(t *testing.T) {
	res := []int{1, 2}
	quickSort(res)
	fmt.Println(res)
	res = []int{3, 2, 1}
	quickSort(res)
	fmt.Println(res)
}
