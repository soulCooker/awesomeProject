package graph

func numIslands(grid [][]byte) int {
	res := 0
	// 第一次循环遍历行
	for i := range grid {
		// 	第二次循环遍历列
		for j := range grid[i] {
			//每个未访问的陆地点
			if grid[i][j] == '1' {
				// 扩展相邻陆地点
				visitIsland(grid, i, j)
				// 		岛屿数量+1
				res++
			}
		}
	}

	// 返回岛屿数量
	return res
}

func visitIsland(grid [][]byte, i, j int) {
	// 如果不是未探索陆地点，直接返回
	if grid[i][j] != '1' {
		return
	}
	// 			更新陆地点为已访问
	grid[i][j] = '2'
	// 对上下左右地点递归扩展相邻点
	if i-1 >= 0 {
		visitIsland(grid, i-1, j)
	}
	if i+1 < len(grid) {
		visitIsland(grid, i+1, j)
	}
	if j-1 >= 0 {
		visitIsland(grid, i, j-1)
	}
	if j+1 < len(grid[i]) {
		visitIsland(grid, i, j+1)
	}
}
