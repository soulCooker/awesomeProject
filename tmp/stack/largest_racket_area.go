package stack

/*
*
按照高的方式枚举所有可能的矩形，已位置i的柱子为高，向右扩散的过程中，维护一个单调递增的栈。栈的元素包括，柱子的高度和柱子的位置。在栈里维护矩形的左边界，当低于栈顶柱子出现时，其一定是右边界。
对于高度<栈顶高度，找到了栈顶柱子的右边界，对栈顶出栈：

	计算栈顶柱子的面积=（当前位置-柱子自己的坐标）*柱子的高度
	如果大于当前矩形面积最大值，则更新

当找到高度>=栈顶柱子的高度时，将柱子入栈。

对于栈的所有剩余元素，计算len(柱子数组总长)-柱子坐标*柱子高度，如果大于当前矩形面积最大值，则更新
*
*/

type Node struct {
	Height int
	Pos    int
}

func largestRectangleArea(heights []int) int {
	stack := make([]*Node, 0)
	res := 0
	for pos, h := range heights {
		topH := 0
		leftMostPos := pos
		for len(stack) > 0 {
			topH = stack[len(stack)-1].Height
			nodePos := stack[len(stack)-1].Pos
			if h < topH {
				res = max((pos-nodePos)*topH, res)
				stack = stack[0 : len(stack)-1]
				leftMostPos = nodePos //每次出栈时，需要更新当前柱子所在矩形的最左位置为出栈柱子的位置
			} else {
				break
			}
		}

		if len(stack) == 0 || h > topH {
			stack = append(stack, &Node{
				Height: h,
				Pos:    leftMostPos,
			})
		}
	}

	for len(stack) > 0 {
		topH := stack[len(stack)-1].Height
		nodePos := stack[len(stack)-1].Pos
		res = max((len(heights)-nodePos)*topH, res)
		stack = stack[0 : len(stack)-1]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
