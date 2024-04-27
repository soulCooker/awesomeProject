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
