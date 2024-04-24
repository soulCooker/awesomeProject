给你一个二叉树的根节点 root ， 检查它是否轴对称。

 

示例 1：
![示例](../../img/1698026966-JDYPDU-image.png)

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false
 

提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100
 

进阶：你可以运用递归和迭代两种方法解决这个问题吗？

思路：
对称树，子问题，可以用广度优先遍历，判断每一层的节点是否对称
该层节点对称判断：中心对称位置上节点的值相等且位置下标对称。

伪代码：
当前层节点列表
下一层节点列表

如果当前层节点列表不为空
    判断当前层节点中心对称位置上节点的值相等且位置下标对称。
    遍历当前层节点列表，将左右子树添加到下一层节点列表
    更新当前层节点列表为下一层节点列表

边界情况：
层只有一个节点时的特殊情况，对于非第一层必定不对称。推广：层节点数为奇数时，除了第一层，必定不对称。