package tmp

/*
*
5 2 9 1 8 4 1
*/
func quickSort(arr []int) {
	doQuickSort(arr, 0, len(arr))
}

func doQuickSort(arr []int, start, end int) {
	var direction int
	var pivot = start
	start++
	for start < end {
		if direction == 0 {
			if arr[pivot] > arr[start] {
				swap(arr, pivot, start)
			} else {
				swap(arr, pivot, start)
				start++
			}
		} else if direction == 1 {

		}
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
