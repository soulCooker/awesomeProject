package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := mergeSorted(intervals)
	fmt.Println(res)
}

func mergeSorted(intervalsInput [][]int) [][]int {
	sort.Slice(intervalsInput, func(i, j int) bool {
		return intervalsInput[i][0] < intervalsInput[j][0]
	})

	merge := make([][]int, 0)

	for _, interval := range intervalsInput {
		if len(merge) == 0 || merge[len(merge)-1][1] < interval[0] {
			merge = append(merge, interval)
		} else {
			merge[len(merge)-1][1] = max(merge[len(merge)-1][1], interval[1])
		}
	}

	return merge
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	} else {
		return v2
	}
}
