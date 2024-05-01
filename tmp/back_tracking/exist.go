package backtracking

func exist(board [][]byte, word string) bool {
	// 处理当前节点：对比格子和当前字符失败返回false
	// 更新状态：成功则将格子设置为已访问
	// 递归处理子问题，并根据子问题结果计算出父问题结果： 在四个方向尝试递归搜索下一个字符，只要任何一个方向的回溯说明匹配成功，返回true
	// 恢复状态：将格子设置为未访问，返回匹配失败。

	return false
}
