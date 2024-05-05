package dp

/*
*
状态转移方程：h，值：i
将问题转化为若干个的子问题，且从子问题可以推导出问题结果
对于n， 可以由 [1,log n]中的数的平方和组成，遍历这些数j，可以转化成一系列子问题h(i-j*j)
其中最小值+1即是当前问题的解。
*
*/
func numSquares(n int) int {
	return 0
}
