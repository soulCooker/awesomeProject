package backtracking

import (
	"fmt"
	"testing"
)

func TestLargestTeamNum(t *testing.T) {
	res := findLargestTeamNum([]int{1, 1, 3}, 3)
	fmt.Println("res:", res)
}
