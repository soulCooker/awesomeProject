package main

import (
	"container/heap"
	"fmt"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	res := merge(intervals)
	fmt.Println(res)
}

var intervals [][]int

type IntervalHeap struct {
	hpArr []int
}

func (hp IntervalHeap) Len() int {
	return len(hp.hpArr)
}

func (hp IntervalHeap) Less(i, j int) bool {
	return intervals[hp.hpArr[i]][0] < intervals[hp.hpArr[j]][0]
}

func (hp IntervalHeap) Swap(i, j int) {
	tmp := hp.hpArr[i]
	hp.hpArr[i] = hp.hpArr[j]
	hp.hpArr[j] = tmp
}

func (hp *IntervalHeap) Push(x any) {
	hp.hpArr = append(hp.hpArr, x.(int))
}

func (hp *IntervalHeap) Pop() any {
	res := hp.hpArr[len(hp.hpArr)-1]
	hp.hpArr = hp.hpArr[0 : len(hp.hpArr)-1]
	return res
}

func merge(intervalsInput [][]int) [][]int {
	intervals = intervalsInput
	hp := &IntervalHeap{hpArr: make([]int, 0)}

	for i := 0; i < len(intervals); i++ {
		hp.Push(i)
	}
	heap.Init(hp)

	res := make([][]int, 0)
	cur := intervals[heap.Pop(hp).(int)]
	for hp.Len() > 0 {
		next := intervals[heap.Pop(hp).(int)]
		if next[0] <= cur[1] {
			if next[1] <= cur[1] {
				continue
			} else {
				cur = []int{cur[0], next[1]}
			}
		} else {
			res = append(res, cur)
			cur = next
		}
		if hp.Len() == 0 {
			res = append(res, cur)
		}
	}

	return res
}
