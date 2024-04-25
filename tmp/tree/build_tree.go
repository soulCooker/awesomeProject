package tree

/*
*
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/

func BuildTree(preorder []int, inorder []int) *TreeNode {
	return doBuildSubTree(preorder, inorder, 0, len(preorder), 0, len(preorder))
}

func doBuildSubTree(preorder []int, inorder []int, pStart, pEnd, iStart, iEnd int) *TreeNode {
	// 	如果起始下标相等，直接返回空节点
	if iStart == iEnd {
		return nil
	}

	// 	遍历中序找到中序遍历根节点值所在的下标
	curRoot := &TreeNode{
		Val: preorder[pStart],
	}
	split := iStart
	for split < iEnd {
		if inorder[split] == preorder[pStart] {
			break
		}
		split++
	}

	// 	开端到下标就是左子树中旬遍历结果，长度为下标-开端
	// 	下标+1到结尾就是右子树的中序遍历结果
	// 	先序遍历切片开始+1到开始+左子树中旬长度+1的偏移量对应的子序列是左子树的前序遍历
	// 	先序遍历开始+左子树中旬长度的偏移量+1的偏移量开始，到切片结束的子序列是右子树的前序遍历结果

	// 	递归处理左子树的前、中序结果生成左子树
	curRoot.Left = doBuildSubTree(preorder, inorder, pStart+1, pStart+split-iStart+1, iStart, split)
	// 	递归处理右子树的前、中序结果生成右子树
	curRoot.Right = doBuildSubTree(preorder, inorder, pStart+split-iStart+1, pEnd, split+1, iEnd)
	// 返回根节点
	return curRoot
}
