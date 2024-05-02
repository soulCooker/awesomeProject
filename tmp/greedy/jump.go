package greedy

/*
*
知道第k跳的最远位置，计算第k+1跳可以到达最远位置。k从1开始，第一个可以到达终点的k就是最小跳跃次数

初始化当前最远可到达位置为 farest = nums[i]
跳跃次数k=1
当前位置i小于farest时

	初始化newFarest
	k++
	遍历从当前位置到第farest最远位置的中间位置;i++
		计算该位置可以到达的最远位置=i+nums[i]
		如果最远位置大于数组长度，返回k+1
		否则，如果大于newFarest，更新newFarest
	更新farest为newFarest

return -1

坑点：
注意长度为1的特殊值处理，已经在终点，跳跃次数为0
注意长度k=1的特殊情况，也就是起点处的特殊情况处理
两层循环要注意循环终止条件
*
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var (
		count  = 1
		farest = nums[0]
		index  = 0
	)

	if farest >= len(nums)-1 {
		return count
	}

	for index <= farest {
		count++
		newFarest := 0
		for ; index <= farest; index++ {
			curFarest := index + nums[index]
			if curFarest >= len(nums)-1 {
				return count
			}
			if curFarest > newFarest {
				newFarest = curFarest
			}
		}
		if newFarest > farest {
			farest = newFarest
		}
	}
	return count
}
