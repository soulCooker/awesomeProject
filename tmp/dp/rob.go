package dp

/*
*
当投窃第k间房屋时，最大收益等于投窃第k-2和第k-3房间的最大收益的较大值+当前房间的收益
f(1) = in(0)
f(2) = max(in(0), in(1))
f(3) = max(in(1), in(0)+in(2))
f(k) = max(max (f(k-2),f(k-3)) + in(k), f(k-1))

坑点：
不能用递归，复杂度是n的三次方
*
*/
func robBackTracking(nums []int) int {

	var dynamicPro func(k int) int

	dynamicPro = func(k int) int {
		if k == 1 {
			return nums[0]
		} else if k == 2 {
			return max(nums[0], nums[1])
		} else if k == 3 {
			return max(nums[1], nums[0]+nums[2])
		} else {
			return max(max(dynamicPro(k-2), dynamicPro(k-3))+nums[k-1], dynamicPro(k-1))
		}
	}
	return dynamicPro(len(nums))
}

func rob(nums []int) int {
	k := len(nums)
	if k == 1 {
		return nums[0]
	} else if k == 2 {
		return max(nums[0], nums[1])
	} else if k == 3 {
		return max(nums[1], nums[0]+nums[2])
	}
	pprev, prev, cur := nums[0], max(nums[0], nums[1]), max(nums[1], nums[0]+nums[2])

	for i := 4; i <= len(nums); i++ {
		pprev, prev, cur = prev, cur, max(max(prev, pprev)+nums[i-1], cur)
	}

	return cur
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
