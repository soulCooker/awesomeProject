package main

/**
输入：matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
输出：[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
*/

func rotate(matrix [][]int) {
	var len = len(matrix)

	if len == 1 {
		return
	}

	startI := 0
	startJ := 0

	for len > 0 {
		cur := matrix[startI][startJ]
		for i := 0; i < len-1; i++ {
			for j := 0; j < 4; j++ {
				newI := startJ
				newJ := len - 1 - startI
				tmp := matrix[newI][newJ]
				matrix[newI][newJ] = cur
				cur = tmp
				startI = newI
				startJ = newJ
			}
		}

		len -= 2
	}
}
