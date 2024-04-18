package tmp

import "fmt"

func test() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	var (
		firstRowZero = 0
		firstColZero = 0
	)

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColZero = 1
			break
		}
	}

	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			firstRowZero = 1
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				break
			}
		}
	}

	for j := 1; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				break
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 0; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := 0; j < len(matrix); j++ {
				matrix[j][i] = 0
			}
		}
	}

	if firstColZero == 1 {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
	if firstRowZero == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
}
