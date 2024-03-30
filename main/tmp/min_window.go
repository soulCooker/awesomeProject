package main

import "fmt"

func main() {
	fmt.Println(int('z') - int('Z'))
	res := minWindow("cabwefgewcwaefgcf", "cae")
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	var (
		tBucket = [64]int{}
		sBucket = [64]int{}
	)

	for _, ch := range t {
		tBucket[ch-'A']++
	}

	var (
		i = 0
		j = 0
	)

	var res string

	for ; j < len(s); j++ {
		sBucket[s[j]-'A']++
		if cover(tBucket, sBucket) {
			fmt.Println("index:", i, j)
			for ; i <= j; i++ {
				sBucket[s[i]-'A']--
				if !cover(tBucket, sBucket) {
					fmt.Println("res:", res)
					if res == "" || len(res) > j-i+1 {
						res = s[i : j+1]
					}
					i++
					break
				}
			}
		}
	}

	return res
}

func cover(arr1, arr2 [64]int) bool {
	for i, val := range arr1 {
		if val > arr2[i] {
			return false
		}
	}
	return true
}
