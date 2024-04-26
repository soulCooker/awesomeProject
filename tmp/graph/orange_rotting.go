package graph

func orangesRotting(grid [][]int) int {
	// 初始化当前节点队列
	curLevel := make([][2]int, 0)

	// 第一次循环遍历行
	for i := range grid {
		// 	第二次循环遍历列
		for j := range grid[i] {
			// 		将腐烂的橘子入队
			if grid[i][j] == 2 {
				//标记为已处理
				grid[i][j] = 3
				curLevel = append(curLevel, [2]int{i, j})
			}
		}
	}

	// 深度遍历但节点队列
	res := 0
	for len(curLevel) > 0 {
		// 	递增遍历深度
		nextLevel := make([][2]int, 0)
		for _, pos := range curLevel {
			i, j := pos[0], pos[1]
			//将四个子节点的未腐烂橘子标记为已处理并入队
			if ok, node := checkInAreaAndNotRottingThenMark(i-1, j, grid); ok {
				nextLevel = append(nextLevel, *node)
			}
			if ok, node := checkInAreaAndNotRottingThenMark(i+1, j, grid); ok {
				nextLevel = append(nextLevel, *node)
			}
			if ok, node := checkInAreaAndNotRottingThenMark(i, j-1, grid); ok {
				nextLevel = append(nextLevel, *node)
			}
			if ok, node := checkInAreaAndNotRottingThenMark(i, j+1, grid); ok {
				nextLevel = append(nextLevel, *node)
			}
		}
		res++
		curLevel = nextLevel
	}

	// 第一次循环遍历行
	for i := range grid {
		// 	第二次循环遍历列
		for j := range grid[i] {
			//	如果存在未腐烂的橘子，返回-1
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	// 返回BFS遍历深度-1
	if res > 0 {
		return res - 1
	}
	return 0
}

func checkInAreaAndNotRottingThenMark(i, j int, grid [][]int) (bool, *[2]int) {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[i]) && grid[i][j] == 1 {
		grid[i][j] = 3
		return true, &[2]int{i, j}
	}
	return false, nil
}
