108. 将有序数组转换为二叉搜索树
简单
相关标签
相关企业
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵平衡二叉搜索树。

示例 1：
![alt text](../../img/btree1.jpg)

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：


输入：nums = [1,3]
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
 

提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 按 严格递增 顺序排列

思路：
取当前切片中位数作为根节点，对下标小于和大于中位数的部分分别做递归调用，递归调用结果分别作为根节点的左右子树。
如果切片长度 <= 0，直接返回nil值

边界情况：长度为1、2的切片，左右子切片的长度为0