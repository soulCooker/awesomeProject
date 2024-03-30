package tree

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

//[1,3,2,5,3,null,9]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelNode struct {
	*TreeNode
	Index int
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func widthOfBinaryTree(root *TreeNode) int {
	var (
		currentLevel []*LevelNode
		nextLevel    []*LevelNode
		maxWidth     int
	)

	if root == nil {
		return 0
	}

	currentLevel = append(currentLevel, &LevelNode{
		TreeNode: root,
		Index:    0,
	})

	for len(currentLevel) > 0 {
		maxWidth = max(currentLevel[len(currentLevel)-1].Index-currentLevel[0].Index+1, maxWidth)
		for _, node := range currentLevel {
			if node == nil {
				continue
			}

			if node.Left != nil {
				nextLevel = append(nextLevel, &LevelNode{
					TreeNode: node.Left,
					Index:    2 * node.Index,
				})
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, &LevelNode{
					TreeNode: node.Right,
					Index:    2*node.Index + 1,
				})
			}
		}

		currentLevel = nextLevel

		nextLevel = make([]*LevelNode, 0)
	}

	return maxWidth
}
