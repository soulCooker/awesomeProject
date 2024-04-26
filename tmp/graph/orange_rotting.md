994. 腐烂的橘子
中等
相关标签
相关企业
在给定的 m x n 网格 grid 中，每个单元格可以有以下三个值之一：

值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子。
每分钟，腐烂的橘子 周围 4 个方向上相邻 的新鲜橘子都会腐烂。

返回 直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1 。

 

示例 1：
![alt text](../../img/oranges.png)


输入：grid = [[2,1,1],[1,1,0],[0,1,1]]
输出：4
示例 2：

输入：grid = [[2,1,1],[0,1,1],[1,0,1]]
输出：-1
解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个方向上。
示例 3：

输入：grid = [[0,2]]
输出：0
解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。
 

提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 10
grid[i][j] 仅为 0、1 或 2


思路：
两轮遍历
第一轮处理腐烂橘子及其四个方向进行腐烂dfs递归遍历，最终这个橘子堆完全腐烂的时间为判断四个方向的腐烂时长最大值+1，
遍历如果遇到新鲜侧继续遍历，如果遇到腐烂橘子，说明这条腐烂路径会从两端开始腐烂，速度快一倍，因此(腐烂时长+1)/2。如果遇到空白格，说明腐烂流程结束。

第二轮，查看是否还有未腐烂的橘子，如有，则返回-1

伪代码：
初始化当前最短分钟数=0
初始化当前节点队列
第一次循环遍历行
    第二次循环遍历列
        将腐烂的橘子入队

深度遍历但前节点队列
    递增遍历深度

第一次循环遍历行
    第二次循环遍历列
        如果存在未腐烂的橘子，返回-1
        
返回BFS遍历深度-1


坑点：
1、橘子堆中可能有多个橘子，会加快腐败速度。因此不能从一个腐败橘子开始深度遍历，需要同时处理所有腐败橘子。
2、可以使用广度优先遍历，同时对所有腐败橘子进行遍历。
3、注意不要将未腐败橘子重复入队，可以在入队时，就将橘子标记为腐败

边界条件：
没有橘子