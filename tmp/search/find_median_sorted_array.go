package search

/*
*
解法一：复杂度不稳定，极端情况还是o（n）
借鉴归并排序算法，不需要正常归并，只需维护归并排序的待插入下标，当下标到达中间位置时，停止归并
用双指针p,q ，每次更新较小值的所在数组的下标。
用二分法优化：
如果 sIdx+lIdx < 中位数下标 且sIdx在数组范围内

	每次在较小值的所在数组通过二分查找较大值+1值的插入位置
	sIdx = start

如果 sIdx+lIdx < 中位数下标

	更新lIdx = median - sIdx

如果 sIdx == len(sNums)

	如果sIdx 大于median
	中位数在sNum中
	中位数在lNums中
		如果和长度为奇数，中位数为lIdx
		如果和长度为偶数，中位数为(lIdx,lIdx+1)	/2

否则

	如果和长度为奇数，中位数在sIdx和lIdx的较大值
	如果和长度为偶数，中位数为

解法二：
转换为计算两个有序数组中第k大的数
*
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		lIdx  int
		sIdx  int
		lNums = nums1
		sNums = nums2
	)

	median := (len(nums1) + len(nums2) - 1) / 2

	if len(nums1) == 0 {
		if len(nums2)%2 == 0 {
			return float64(nums2[median]+nums2[median+1]) / 2
		} else {
			return float64(nums2[median])
		}
	} else if len(nums2) == 0 {
		if len(nums1)%2 == 0 {
			return float64(nums1[median]+nums1[median+1]) / 2
		} else {
			return float64(nums1[median])
		}
	}

	for lIdx+sIdx < median {
		if sNums[sIdx] > lNums[lIdx] {
			sIdx, lIdx = lIdx, sIdx
			sNums, lNums = lNums, sNums
		}

		var (
			start = sIdx
			end   = len(sNums) - 1
		)

		for start <= end {
			mid := (start + end) / 2
			if lNums[lIdx]+1 >= sNums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}

		sIdx = min(start-1, median-lIdx)
	}

	var (
		s1 int
		s2 int
		l2 int
		l1 int
	)

	s1 = sNums[sIdx]
	l2 = lNums[lIdx]

	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(min(s1, l2))
	}

	l1 = lNums[lIdx-1]

	if s1 < l2 {
		if sIdx+1 < len(sNums) {
			s2 = sNums[sIdx+1]
			if s2 < l1 {
				return float64(s2+l2) / 2
			} else {
				return float64(l2+l1) / 2
			}
		} else {

		}
	} else {
		if l1 > s2 {
			return float64(s2+s1) / 2
		} else {
			return float64(s1+l1) / 2
		}
	}
	return -1
}

func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	totalNums := len(nums1) + len(nums2)
	midIndex := totalNums / 2
	if totalNums%2 == 1 {
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		return float64(getKthElement(nums1, nums2, midIndex) + getKthElement(nums1, nums2, midIndex+1))
	}
}

func getKthElement(nums1, nums2 []int, k int) int {
	var (
		idx1 = 0
		idx2 = 0
	)

	for {
		if idx1 == len(nums1) {
			return nums2[idx2+k-1]
		}
		if idx2 == len(nums2) {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}

		half := k / 2
		newIdx1 := min(idx1+half, len(nums1)) - 1
		newIdx2 := min(idx2+half, len(nums2)) - 1

		pivot1 := nums1[newIdx1]
		pivot2 := nums2[newIdx2]

		if pivot1 <= pivot2 {
			k -= (newIdx1 - idx1 + 1)
			idx1 = newIdx1 + 1
		} else {
			k -= (newIdx2 - idx2 + 1)
			idx2 = newIdx2 + 1
		}
	}
}

// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/258842/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func min(val1, val2 int) int {
	if val1 > val2 {
		return val2
	} else {
		return val1
	}
}
