package tmp

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)

	var (
		rowSize = len(matrix)
		colSize = len(matrix[0])
		colIdx  = 0
		rowIdx  = 0
		round   = 0
	)

	for rowSize > 0 && colSize > 0 {
		colIdx = round
		rowIdx = round

		for i := 0; i < colSize-1; i++ {
			res = append(res, matrix[rowIdx][colIdx])
			colIdx++
		}

		if rowSize == 1 {
			res = append(res, matrix[rowIdx][colIdx])
			break
		}

		for i := 0; i < rowSize-1; i++ {
			res = append(res, matrix[rowIdx][colIdx])
			rowIdx++
		}

		if colSize == 1 {
			res = append(res, matrix[rowIdx][colIdx])
			break
		}

		for i := 0; i < colSize-1; i++ {
			res = append(res, matrix[rowIdx][colIdx])
			colIdx--
		}
		for i := 0; i < rowSize-1; i++ {
			res = append(res, matrix[rowIdx][colIdx])
			rowIdx--
		}

		rowSize -= 2
		colSize -= 2
		round++
	}

	return res
}
