package main

import "fmt"

func findPeakElement(nums []int) int {
	// 正常来讲二分查找区间 [0,n) 但存在边界情况，不将最后一个元素划入查找范围就不会遇到
	left, right := 0, len(nums)-1 // [0, n-1)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
}
