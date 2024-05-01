package backtracking

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	res := permute([]int{1, 2, 3})

	fmt.Println(res)
}

func TestPartion(t *testing.T) {
	res := partition("aab")
	fmt.Println("res:", res)
}

func TestComSum(t *testing.T) {
	res := combinationSum2([]int{2, 3, 6, 7}, 7)
	fmt.Print(res)
}

func TestCopy(t *testing.T) {
	arr1 := []int{1, 2, 3}

	arr2 := make([]int, len(arr1))
	len := copy(arr2, arr1)

	fmt.Println(len, ",", arr2)
}

func TestBitOp(t *testing.T) {
	a := 'a'

	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", 'A')
	fmt.Printf("%b\n", 'Z')
	fmt.Printf("%b\n", 'z')

	a = a | 1<<7

	fmt.Printf("%b\n", a)

	a = a & (1<<7 - 1)

	fmt.Printf("%b\n", a)
}
