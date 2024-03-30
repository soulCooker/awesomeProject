package main

type Bucket struct {
	Max int
	Min int
}

/**
100 1 4 3 10 2
**/

//func longestConsecutive(nums []int) int {
//	numSet := map[int]bool{}
//	for _, num := range nums {
//		numSet[num] = true
//	}
//	longestStreak := 0
//	for num := range numSet {
//		if !numSet[num-1] {
//			currentNum := num
//			currentStreak := 1
//			for numSet[currentNum+1] {
//				currentNum++
//				currentStreak++
//			}
//			if longestStreak < currentStreak {
//				longestStreak = currentStreak
//			}
//		}
//	}
//	return longestStreak
//}

func longestConsecutive2(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestConsecutiveCount := 0

	for num := range numSet { // 遍历map，避免倒序case导致的最大时间复杂度
		if !numSet[num-1] {
			curNum := num
			curConsecutiveCount := 1
			for numSet[curNum+1] {
				curNum++
				curConsecutiveCount++
			}

			if curConsecutiveCount > longestConsecutiveCount {
				longestConsecutiveCount = curConsecutiveCount
			}
		}
	}
	return longestConsecutiveCount
}
