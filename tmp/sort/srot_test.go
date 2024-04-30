package sort

import (
	"fmt"
	"testing"
)

func TestQuictSort(t *testing.T) {
	res := []int{5, 2, 9, 1, 8, 4, 1}
	quickSort(res)
	fmt.Println(res)
}
