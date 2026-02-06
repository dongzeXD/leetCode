package main

import "fmt"

// [0,n)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1 // [0, n-1)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func lowerBound(src []int, left, right, target int) int {
	for left < right {
		mid := left + (right-left)/2
		if src[mid] > target {
			right = mid
		} else if src[mid] < target {
			left = mid + 1
		} else if src[mid] == target {
			right = mid
		}
	}

	if src[left] != target {
		return -1
	}
	return left
}

func search(nums []int, target int) int {
	i := findMin(nums)
	if target > nums[len(nums)-1] { // target 在第一段
		return lowerBound(nums, 0, i-1, target) // [0, i-1)
	}
	// target 在第二段
	return lowerBound(nums, i, len(nums)-1, target) // [i, n-1)
}

// [0,n-1)
// 3 3 4 5 1 2 3 3 3 3 3 3 3
func search2(nums []int, target int) int {
	flag := nums[len(nums)-1]
	left, right := 0, len(nums)-1
	middle := 0

	for left < right {
		middle = left + (right-left)/2
		if nums[middle] <= flag && target > flag {
			right = middle
		} else if nums[middle] > flag && target <= flag {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle
		} else if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] == target {
			right = middle
		}
	}
	if nums[left] != target {
		return -1
	}
	return left
}

func main() {
	fmt.Println(search2([]int{3, 3, 4, 5, 1, 2, 3, 3, 3}, 3))
	// 153,确定最小值位置划分区间，然后对区间进行二分查找
	src := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(src, 1))
}
