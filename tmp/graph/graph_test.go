package graph

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	graph := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	res := numIslands(graph)
	fmt.Println(res)
}
