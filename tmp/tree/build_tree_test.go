package tree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	root := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	fmt.Println(PreorderTraversal(root))

	fmt.Println(InorderTraversal(root))
}
