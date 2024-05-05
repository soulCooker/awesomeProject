package dp

/*
*
边界情况：

*
*/
func generate(numRows int) [][]int {
	res := make([][]int, 0)

	for i := 0; i < numRows; i++ {
		line := make([]int, i+1, i+1) //slice初始化时，长度为i+1
		line[0] = 1
		for j := 1; j < i; j++ { //由于i是行数-1，所以终止条件为j<i
			line[j] = res[len(res)-1][j-1] + res[len(res)-1][j]
		}
		line[len(line)-1] = 1
		res = append(res, line)
	}

	return res
}
