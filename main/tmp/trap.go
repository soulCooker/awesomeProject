package main

func trap(height []int) int {
	leftMost := make([]int, len(height))
	rightMost := make([]int, len(height))

	leftMost[0] = height[0]
	rightMost[len(height)-1] = height[len(height)-1]

	var i = 1

	for i < len(height) {
		leftMost[i] = max(leftMost[i-1], height[i])
		i++
	}

	var j = len(height) - 2
	for j >= 0 {
		rightMost[j] = max(rightMost[j+1], height[j])
		j--
	}

	var trap = 0
	for i, h := range height {
		trap += min(leftMost[i], rightMost[i]) - h
	}
	return trap
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func trap2(height []int) int {
	var (
		i = 0
		j = len(height) - 1
	)

	trapSpace := 0

	for i < j {
		i = findEdge(height, 1, i, j)
		if i == -1 {
			break
		}
		j = findEdge(height, -1, i, j)
		if j == -1 {
			break
		}
		if height[i] < height[j] {
			blockInBucket := 0
			index := i + 1
			for index <= j {
				if height[index] >= height[i] {
					trapSpace += height[i]*(index-i-1) - blockInBucket
					i = index
					// fmt.Printf("new i:%d, trapSpace:%d\n", i, trapSpace)
					break
				} else {
					blockInBucket += height[index]
				}
				index++
			}
		} else {
			blockInBucket := 0
			index := j - 1
			for index >= i {
				if height[index] >= height[j] {
					trapSpace += height[j]*(j-index-1) - blockInBucket
					j = index
					// fmt.Printf("new j:%d, trapSpace:%d\n", j, trapSpace)
					break
				} else {
					blockInBucket += height[index]
				}
				index--
			}
		}
	}
	return trapSpace
}

func findEdge(height []int, direction int, start, end int) int {
	// fmt.Printf("start:%d, end:%d,direction:%d\n", start, end, direction)
	for start < end {
		if direction == 1 {
			start++
			if height[start] < height[start-1] {
				return start - 1
			}
		} else if direction == -1 {
			end--
			if height[end] < height[end+1] {
				return end + 1
			}
		}
	}
	return -1
}
