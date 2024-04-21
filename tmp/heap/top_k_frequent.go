package heap

type Counter struct {
	Count int
	Val   int
}

func topKFrequent(nums []int, k int) []int {
	counters := make(map[int]*Counter)
	for _, v := range nums {
		if counter, ok := counters[v]; ok {
			counter.Count++
		} else {
			counters[v] = &Counter{
				Count: 1,
				Val:   v,
			}
		}
	}

	counterList := make([]*Counter, 0)
	for _, v := range counters {
		counterList = append(counterList, v)
	}

	doTopKFrequet(counterList, k, 0, len(counterList)-1)
	res := make([]int, 0)
	for _, v := range counterList[0:k] {
		res = append(res, v.Val)
	}
	return res
}

func doTopKFrequet(nums []*Counter, k, i, j int) {
	if i == j {
		return
	}

	var (
		l = i - 1
		r = j + 1
	)

	partition := nums[i].Count

	for l < r {
		for l++; nums[l].Count > partition; l++ {
		}
		for r--; nums[r].Count < partition; r-- {
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}

	if k <= r {
		doTopKFrequet(nums, k, i, r)
	} else {
		doTopKFrequet(nums, k, r+1, j)
	}
}
