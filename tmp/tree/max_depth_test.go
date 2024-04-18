package tree

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	maxDepth := MaxDepth(root)

	fmt.Println(maxDepth)
}
