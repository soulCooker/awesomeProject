package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CommonAncestorNode struct {
	*TreeNode
	Visited   int
	PAncestor bool
	QAncestor bool
	CLeft     *CommonAncestorNode
	CRight    *CommonAncestorNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := make([]*CommonAncestorNode, 0)
	// 初始化节点栈
	stack = append(stack, &CommonAncestorNode{
		TreeNode: root,
	})

	// 节点栈不为空时循环处理
	for len(stack) > 0 {
		// 	取栈顶节点
		curRoot := stack[len(stack)-1]
		// 	节点未访问
		if curRoot.Visited == 0 {
			// 		将节点的左右子树入栈
			if curRoot.Left != nil {
				l := &CommonAncestorNode{
					TreeNode: curRoot.Left,
				}
				curRoot.CLeft = l
				stack = append(stack, l)
			}
			if curRoot.Right != nil {
				r := &CommonAncestorNode{
					TreeNode: curRoot.Right,
				}
				curRoot.CRight = r
				stack = append(stack, r)
			}
			// 		节点设置为已访问
			curRoot.Visited = 1
		} else {
			// 	节点已访问
			// 	   节点是否为p祖先 = 左右子树是否为p的祖先 或 节点值等于p
			curRoot.PAncestor = p.Val == curRoot.Val
			curRoot.QAncestor = q.Val == curRoot.Val
			if curRoot.CLeft != nil {
				curRoot.PAncestor = curRoot.PAncestor || curRoot.CLeft.PAncestor
				curRoot.QAncestor = curRoot.QAncestor || curRoot.CLeft.QAncestor
			}
			if curRoot.CRight != nil {
				curRoot.PAncestor = curRoot.PAncestor || curRoot.CRight.PAncestor
				curRoot.QAncestor = curRoot.QAncestor || curRoot.CRight.QAncestor
			}
			//     判断 左右子树是否为q的祖先 或 节点值等于q，为q祖先
			if curRoot.PAncestor && curRoot.QAncestor {
				return curRoot.TreeNode
			}

			//     否则节点出栈
			stack = stack[0 : len(stack)-1]
		}

	}
	// 返回空
	return nil
}
