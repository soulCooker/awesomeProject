543. 二叉树的直径
简单
相关标签
相关企业
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。

 

示例 1：
![示例一](../../img/diamtree.jpg)

输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。

思路：
节点树的最大长度为经过根节点的最长路径、左子树的最大直径、右子树的最大直径中的最大值。经过根节点的最长路径为左子树的高度+右子树的高度。
树的高度=节点到叶子节点的最大距离，也即左、右子树的高度的最大值+1。

深度优先遍历，后序遍历
取出栈顶节点
    节点为未访问
        左子节点，右子节点进栈，标记节点为已访问
    节点为已访问
        节点出栈
        节点高度更新为左、右子树的高度的最大值+1
        计算节点为根的树最大直径更新为 左子树的高度+右子树的高度的最大值

坑点：
需要利用左右子树的，因此需要在树节点中维护左右子树的处理状态。