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

// -1 -2 4 -5 2 9 -2
// m:-1 c:-1 max:-1
// m:-3 c:-3 max: -1
// m:-3 c:1  max: 4
// m:-4 c:-4 max:4
// m:-4 c:-2 max:4
//
//	m:-4 c: 7 max: 11
func maxSubArray2(nums []int) int {
	minPre := 0
	maxSum := math.MinInt
	curSum := 0
	for _, num := range nums {
		curSum += num
		maxSum = max(maxSum, curSum-minPre)
		minPre = min(minPre, curSum)
	}
	return maxSum
}

func main() {
	r := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(r)
}
