package dp

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	var (
		prev = 1
		cur  = 1
	)

	for n > 1 {
		n--
		prev, cur = cur, prev+cur
	}

	return cur
}
