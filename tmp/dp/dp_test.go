package dp

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	res := robBackTracking([]int{183, 219, 57, 193, 94, 233, 202, 154, 65, 240, 97, 234, 100, 249, 186, 66, 90, 238, 168, 128, 177, 235, 50, 81, 185, 165, 217, 207, 88, 80, 112, 78, 135, 62, 228, 247, 211})
	fmt.Println("res:", res)
}
