package dp

/*
*
状态转移方程：
f(i) = 1+ min(f(i-cj)) ,f(i-cj) != -1
f(0) = 0

优化思路：

*
*/
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1, amount+1)
	f[0] = 0
	for balance := 1; balance <= amount; balance++ {
		f[balance] = amount + 1
		for _, v := range coins {
			remain := balance - v
			if remain >= 0 {
				f[balance] = min(f[balance], 1+f[remain])
			}
		}
	}
	if f[amount] < amount+1 {
		return f[amount]
	}
	return -1
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	} else {
		return v2
	}
}
