package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 尾指针
	var tail *TreeNode
	// 初始化树节点栈
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	// 栈不为空时循环处理
	for len(stack) > 0 {
		// 	取栈顶节点
		curRoot := stack[len(stack)-1]
		// 	节点出栈
		stack = stack[0 : len(stack)-1]
		// 	并将右子树入驻并设置节点的右子树指针为空
		if curRoot.Right != nil {
			stack = append(stack, curRoot.Right)
			curRoot.Right = nil
		}
		// 	左子树入栈并将节点的左子树指针设置为空。
		if curRoot.Left != nil {
			stack = append(stack, curRoot.Left)
			curRoot.Left = nil
		}
		// 	将节点用尾插法保存到结果链表
		if tail == nil {
			tail = curRoot
		} else {
			tail.Right = curRoot
			tail = curRoot
		}
	}
}
