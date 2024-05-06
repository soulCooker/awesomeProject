package dp

import "strings"

/*
*
S(j) = S(j-len(w)) + w
因此维护一个已解决子问题的集合，遍历这个子问题集合不断扩展，直到原始问题解决。
子问题的集合->可以表示为下标集合，每个下标集合都重复解决判断是否存在一个单词可以匹配从该位置开始的字符串，存在则添加到一个新的子问题结果集合中。
终止条件：子问题无法继续扩展，或者新的子问题等于原始问题时。
*
*/
func wordBreak(s string, wordDict []string) bool {
	wordBreak := make([]int, len(s), len(s))
	wordBreak[0] = 1
	for idx := 0; idx < len(s); idx++ {
		if wordBreak[idx] == 0 {
			continue
		}

		for _, word := range wordDict {
			if len(word)+idx <= len(s) && strings.Compare(s[idx:idx+len(word)], word) == 0 {
				if idx+len(word) == len(s) {
					return true
				} else {
					wordBreak[idx+len(word)] = 1
				}
			}
		}
	}

	return false
}
