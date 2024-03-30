package main

import "fmt"

func lengthOfLongestSubstring2(s string) int {
	chars := []rune(s)
	bucket := make(map[rune]int)

	var start = 0
	longestSubStringLen := 0
	for end, ch := range chars {
		findIdx, ok := bucket[ch]
		if ok {
			if end-start > longestSubStringLen {
				longestSubStringLen = end - start
			}
			if start < findIdx+1 {
				start = findIdx + 1
			}
			fmt.Println("start:", start, "end:", end, "findIdx:", findIdx)
		}
		bucket[ch] = end
	}
	if len(chars)-start > longestSubStringLen {
		longestSubStringLen = len(chars) - start
	}

	return longestSubStringLen
}
