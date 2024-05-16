package list

func findIncrementSum(target int) [][]int {
	var (
		i = 1
		j = 1
	)
	res := make([][]int, 0)

	sum := 0
	for j <= target {
		for sum < target {
			sum += j
			j++
		}
		for sum > target {
			sum -= i
			i++
		}

		if sum == target {
			subRes := make([]int, 0)
			for num := i; num < j; num++ {
				subRes = append(subRes, num)
			}
			res = append(res, subRes)

			sum -= i
			i++
		}
	}

	return res
}
