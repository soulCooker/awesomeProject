package tree

/**
当前层节点列表
下一层节点列表

如果当前层节点列表不为空
    判断当前层节点中心对称位置上节点的值相等且位置下标对称。
    遍历当前层节点列表，将左右子树添加到下一层节点列表
    更新当前层节点列表为下一层节点列表
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type PositionNode struct {
	*TreeNode
	Index int
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	curLevel := make([]*PositionNode, 0)
	maxWidth := 1
	curLevel = append(curLevel, &PositionNode{
		TreeNode: root,
		Index:    0,
	})

	for len(curLevel) > 0 {
		if maxWidth > 1 && len(curLevel)%2 == 1 {
			return false
		}
		// 	判断当前层节点中心对称位置上节点的值相等且位置下标对称。
		for i := 0; i < len(curLevel)/2; i++ {
			if curLevel[i].Val != curLevel[len(curLevel)-1-i].Val ||
				curLevel[i].Index != maxWidth-1-curLevel[len(curLevel)-1-i].Index {
				return false
			}
		}
		// 遍历当前层节点列表，将左右子树添加到下一层节点列表
		nextLevel := make([]*PositionNode, 0)
		for _, v := range curLevel {
			if v.Left != nil {
				nextLevel = append(nextLevel, &PositionNode{
					TreeNode: v.Left,
					Index:    2 * v.Index,
				})
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, &PositionNode{
					TreeNode: v.Right,
					Index:    2*v.Index + 1,
				})
			}
		}
		// 更新当前层节点列表为下一层节点列表和最大宽度
		curLevel = nextLevel
		maxWidth *= 2
	}

	return true
}
