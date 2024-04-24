98. 验证二叉搜索树
中等
相关标签
相关企业
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

有效 二叉搜索树定义如下：

节点的左子树只包含 小于 当前节点的数。
节点的右子树只包含 大于 当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
 

示例 1：
![alt text](<../../img/tree1 (1).jpg>)

输入：root = [2,1,3]
输出：true
示例 2：


输入：root = [5,1,4,null,null,3,6]
输出：false
解释：根节点的值是 5 ，但是右子节点的值是 4 。
 

提示：

树中节点数目范围在[1, 104] 内
-231 <= Node.val <= 231 - 1


思路：
中序遍历，通过反证法，可以证明每个节点的前序访问节点一定是左子树的最右节点，后序访问节点一定是右子树的最左节点。比较当前值是否大于等于前一个值。如果左、右子树是二叉树，那么根节点的前一个节点必定是左子树的最大值，后一个节点必定是右子树的最小值。
递推法，中序遍历过程中，如果左右子树是二叉搜索树，且节点大于前序节点，小于后序节点，则节点为根的树也是一个二叉树。


踩坑：
需要一个变量标记前序节点值是否已初始化。必须严格递增。