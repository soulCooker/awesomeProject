package main

func maxArea(height []int) int {
	var (
		i       = 0
		j       = len(height) - 1
		maxArea = 0
	)

	for i < j {
		min := 0
		if height[i] > height[j] {
			min = height[j]
			j--
		} else {
			min = height[i]
			i++
		}
		area := min * (j - i + 1)
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
