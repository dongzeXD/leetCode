package main

import "fmt"

func trap(height []int) int {
	leftMaxMap := make(map[int]int)
	rightMaxMap := make(map[int]int)

	for i := 0; i < len(height); i++ {
		left := i - 1
		if left < 0 {
			leftMaxMap[i] = 0
		} else {
			if leftMaxMap[left] > height[left] {
				leftMaxMap[i] = leftMaxMap[left]
			} else {
				leftMaxMap[i] = height[left]
			}
		}
	}

	for j := len(height) - 1; j >= 0; j-- {
		right := j + 1
		if right == len(height) {
			rightMaxMap[j] = 0
		} else {
			if rightMaxMap[right] > height[right] {
				rightMaxMap[j] = rightMaxMap[right]
			} else {
				rightMaxMap[j] = height[right]
			}
		}
	}

	fmt.Println(rightMaxMap, leftMaxMap)

	result := 0
	for i := 0; i < len(height); i++ {
		left := leftMaxMap[i]
		right := rightMaxMap[i]
		r := min(left, right) - height[i]
		if r < 0 {
			continue
		}
		result += r
	}

	return result
}

func main() {
	r := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println(r)
}
