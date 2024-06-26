package greedy

/*
*
思路一：复杂度n2
可用深度遍历+回赎法来解决，同时可以记录已经已经解决的子问题结果。
当遍历的节点可覆盖的范围包含最后一个节点时，找到解。或者到达0节点，不存在解。
对于从节点a开始的dfs，a的值为k，则有k个分支，因此可以考虑用贪婪策略优先遍历更有可能的分支。
问题转化为最小的跳跃次数来到达终点。对与一个节点，现在选择跳跃目标时，选择跳跃次数来到达尽可能远的节点。

思路二：
使用一次遍历解决问题，维护一个可达位置的集合，用右边界位置来表示
依次访问每个节点，如果节点的最远可达位置大于当前可达位置右边界，则更新最远可达位置。
结束条件：如果可达位置大于等于目标位置 或者访问的节点超出当前最远可达位置。
*
*/
func canJump(nums []int) bool {
	fartestPos := nums[0]

	for i := 0; i <= fartestPos; i++ {
		curFartestPos := nums[i] + i
		if curFartestPos >= len(nums)-1 {
			return true
		}
		if curFartestPos > fartestPos {
			fartestPos = curFartestPos
		}
	}

	return false
}
