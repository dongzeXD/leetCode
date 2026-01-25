package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	minPreSum := 0
	maxSum := math.MinInt
	preSum := 0
	for i := 0; i < len(nums); i++ {
		preSum += nums[i]
		maxSum = max(maxSum, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return maxSum
}

func main() {
	r := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(r)
}
