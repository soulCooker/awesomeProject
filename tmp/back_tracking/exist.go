package backtracking

func exist(board [][]byte, word string) bool {
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 处理当前节点：对比格子和当前字符失败返回false
		if board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		// 更新状态：成功则将格子设置为已访问
		board[i][j] |= 1 << 7
		// 递归处理子问题，并根据子问题结果计算出父问题结果： 在四个方向尝试递归搜索下一个字符，只要任何一个方向的回溯说明匹配成功，返回true
		if inArea(i+1, j, board) && (board[i+1][j]&(1<<7)) == 0 && dfs(i+1, j, index+1) {
			return true
		} else if inArea(i, j+1, board) && (board[i][j+1]&(1<<7)) == 0 && dfs(i, j+1, index+1) {
			return true
		} else if inArea(i-1, j, board) && (board[i-1][j]&(1<<7)) == 0 && dfs(i-1, j, index+1) {
			return true
		} else if inArea(i, j-1, board) && (board[i][j-1]&(1<<7)) == 0 && dfs(i, j-1, index+1) {
			return true
		}
		// 恢复状态：将格子设置为未访问，返回匹配失败。
		board[i][j] &= (1<<7 - 1)
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func inArea(i, j int, board [][]byte) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[i])
}
