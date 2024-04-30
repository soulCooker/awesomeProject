package sort

import "fmt"

/*
*
5 2 9 1 8 4 1
*/
func quickSort(arr []int) {
	doQuickSort(arr, 0, len(arr)-1)
}

func doQuickSort(arr []int, start, end int) {
	fmt.Printf("start:%d, end:%d\n", start, end)
	//如果 start>=end，返回
	if start >= end {
		return
	}
	//第一个元素也就start作为初始partition值
	partition := arr[start]
	// 初始化l=start-1,r=end+1
	l := start - 1
	r := end + 1

	//循环处理，对于l<r
	for l < r {
		//循环不变量为 l之前的数<=pivot， r之后的数>=pivot
		for l++; arr[l] < partition; l++ {
		}
		// l++，对于l下标值 < partition值, l++
		// r--，对于r > partition, r--
		for r--; arr[r] > partition; r-- {
		}
		// l指向的值>=partition，r指向的值<=partition
		if l < r {
			arr[l], arr[r] = arr[r], arr[l]
		}
		// 如果l<r
		//	将l与r交换
	}
	//对start，r
	//对r+1，end
	doQuickSort(arr, start, r)
	doQuickSort(arr, r+1, end)
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
