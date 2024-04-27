package graph

// 定义Trie节点
//
//	值
//	子节点Map[子节点值]子节点指针
//	是否单词末尾
type Trie struct {
	Childs [26]*Trie
	IsEnd  bool
}

func Constructor() Trie {
	return Trie{
		Childs: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	// 遍历当前单词的所有字符
	chs := []byte(word)

	matchStartIdx := 0
	curNode := this
	for ; matchStartIdx < len(chs); matchStartIdx++ {
		// 	取当前字符对应在Trie树根的子节点
		nextNode := curNode.Childs[chs[matchStartIdx]-'a']
		// 	子节点存在
		// 		树root，更新为子节点
		// 	子节点不存在
		// 		停止循环
		if nextNode == nil {
			nextNode = &Trie{}
			curNode.Childs[chs[matchStartIdx]-'a'] = nextNode
		}
		curNode = nextNode
	}

	// 更新当前root为单词末尾
	curNode.IsEnd = true
	// 返回
	return
}

func (this *Trie) Search(word string) bool {
	// 遍历当前单词的所有字符
	chs := []byte(word)

	matchStartIdx := 0
	curNode := this
	for ; matchStartIdx < len(chs); matchStartIdx++ {
		// 	取当前字符对应在Trie树根的子节点
		nextNode := curNode.Childs[chs[matchStartIdx]-'a']
		// 	子节点存在
		// 		树root，更新为子节点
		// 	子节点不存在
		// 		停止循环
		if nextNode == nil {
			return false
		}
		curNode = nextNode
	}
	// 如果待匹配字符串起始下标==字符长度且root为单词末端
	// 	返回true
	// 否则
	// 	返回false
	return curNode.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	chs := []byte(prefix)
	matchStartIdx := 0
	curNode := this
	for ; matchStartIdx < len(chs); matchStartIdx++ {
		// 	取当前字符对应在Trie树根的子节点
		nextNode := curNode.Childs[chs[matchStartIdx]-'a']
		// 	子节点存在
		// 		树root，更新为子节点
		// 	子节点不存在
		// 		匹配失败
		if nextNode == nil {
			return false
		}
		curNode = nextNode
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
