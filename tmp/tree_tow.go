package tmp

import (
	"sort"
)

func groupAnagrams2(strs []string) [][]string {
	buckets := make(map[string][]string)
	results := make([][]string, 0)

	for _, str := range strs {
		key := sortStr(str)
		if bucket, ok := buckets[key]; ok {
			buckets[key] = append(bucket, str)
		} else {
			buckets[key] = []string{str}
		}
	}
	for _, strings := range buckets {
		results = append(results, strings)
	}
	return results
}

func sortStr(str string) string {
	result := []rune(str)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return string(result)
}
