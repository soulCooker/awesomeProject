package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type DiameterNode struct {
	*TreeNode
	Heigth  int
	Visited int
	LeftD   *DiameterNode
	RightD  *DiameterNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 0

	stack := make([]*DiameterNode, 0)
	dNode := &DiameterNode{
		TreeNode: root,
	}

	stack = append(stack, dNode)

	for len(stack) > 0 {
		// 取出栈顶节点
		curRoot := stack[len(stack)-1]
		//     节点为未访问
		if curRoot.Visited == 0 {
			//         左子节点，右子节点进栈，标记节点为已访问
			if curRoot.Left != nil {
				leftD := &DiameterNode{
					TreeNode: curRoot.Left,
				}
				stack = append(stack, leftD)
				curRoot.LeftD = leftD
			}
			if curRoot.Right != nil {
				rightD := &DiameterNode{
					TreeNode: curRoot.Right,
				}
				stack = append(stack, rightD)
				curRoot.RightD = rightD
			}
			curRoot.Visited = 1
		} else {
			//     节点为已访问
			//         节点出栈
			stack = stack[0 : len(stack)-1]

			leftHeight := 0
			rightHeight := 0
			if curRoot.LeftD != nil {
				leftHeight = curRoot.LeftD.Heigth
			}
			if curRoot.RightD != nil {
				rightHeight = curRoot.RightD.Heigth
			}

			//         节点高度更新为左、右侧高度的最大值+1
			if leftHeight > rightHeight {
				curRoot.Heigth = leftHeight + 1
			} else {
				curRoot.Heigth = rightHeight + 1
			}

			//         计算节点为根的树最大直径更新为 左侧高度+右侧的高度与当前最大直径 的最大值
			if maxWidth < leftHeight+rightHeight {
				maxWidth = leftHeight + rightHeight
			}
		}
	}

	return maxWidth
}
