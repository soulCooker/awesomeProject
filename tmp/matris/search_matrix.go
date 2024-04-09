package tmp

import "fmt"

/*
*
[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
*/
func TestSearchMatrix() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 31))
}

func searchMatrix(matrix [][]int, target int) bool {
	var (
		r = 0
		c = len(matrix[0]) - 1
	)

	for r < len(matrix) && c >= 0 {
		if matrix[r][c] < target {
			r++
		} else if matrix[r][c] > target {
			c--
		} else {
			return true
		}
	}

	return false
}
