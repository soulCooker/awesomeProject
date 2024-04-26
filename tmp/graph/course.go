package graph

type GraphNode struct {
	Node    int
	Visited int
	OutNode map[int]interface{}
	InNode  map[int]interface{}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 	遍历所有边，构建每个节点的入边和出边集合
	nodeMap := make(map[int]*GraphNode)

	for _, v := range prerequisites {
		var (
			node *GraphNode
			ok   bool
		)
		if node, ok = nodeMap[v[0]]; !ok {
			node = &GraphNode{
				Node:    v[0],
				OutNode: make(map[int]interface{}),
				InNode:  make(map[int]interface{}),
			}
			nodeMap[v[0]] = node
		}
		node.OutNode[v[1]] = 1

		if node, ok = nodeMap[v[1]]; !ok {
			node = &GraphNode{
				Node:    v[1],
				OutNode: make(map[int]interface{}),
				InNode:  make(map[int]interface{}),
			}
			nodeMap[v[1]] = node
		}
		node.InNode[v[0]] = 1
	}
	//     初始化当前层队列为当前节点
	curLevel := make([]*GraphNode, 0)
	for _, v := range nodeMap {
		if len(v.InNode) == 0 {
			curLevel = append(curLevel, v)
		}
	}

	remainCount := len(nodeMap)
	//     队列不会空循环处理
	for len(curLevel) > 0 {
		//         遍历当前层队列
		nextLevel := make([]*GraphNode, 0)
		for _, v := range curLevel {
			for outNodeIndex := range v.OutNode {
				if outNode, ok := nodeMap[outNodeIndex]; ok {
					//删除子节点中已当前节点开始的入边
					delete(outNode.InNode, v.Node)
					//如果子节点的入度为0，将子节点加入下一层节点列表
					if len(outNode.InNode) == 0 {
						nextLevel = append(nextLevel, outNode)
					}
				}
			}
		}
		remainCount -= len(curLevel)
		//         更新当前层= 下一层队列
		curLevel = nextLevel
	}

	return remainCount == 0
}
