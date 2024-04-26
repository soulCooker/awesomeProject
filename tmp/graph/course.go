package graph

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 将边按照起点进行分组，得到节点和所有的出边分组
	// 初始化拓扑编号
	// 循环取下一个未遍历节点时
	// 	拓扑编号递增
	// 	设置节点拓扑编号
	// 	初始化当前层队列为当前节点
	// 	队列不会空循环处理
	// 		遍历当前层队列
	// 			节点拓扑编号>0
	// 				如果出边节点拓扑编号==拓扑编号，存在环
	// 				如果出边节点拓扑编号 < 拓扑编号，忽略
	// 			节点未访问==0
	// 				更新节点拓扑编号
	// 				节点进下一层队列
	// 		更新当前层= 下一层队列
}
