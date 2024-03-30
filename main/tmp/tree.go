package main

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	buckets := make([]map[rune]int, 0)

MyLabel:
	for _, str := range strs {
		bucket := calBucket(str)
		for i, target := range buckets {
			if mapEqual(bucket, target) {
				result[i] = append(result[i], str)
				continue MyLabel
			}
		}
		result = append(result, []string{str})
		buckets = append(buckets, bucket)
	}
	return result
}

func mapEqual(first, second map[rune]int) bool {
	if len(first) != len(second) {
		return false
	}
	for r, c := range first {
		if count, ok := second[r]; ok {
			if count != c {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func calBucket(str string) map[rune]int {
	chars := []rune(str)
	bucket := make(map[rune]int)
	for _, char := range chars {
		bucket[char] += 1
	}
	return bucket
}
