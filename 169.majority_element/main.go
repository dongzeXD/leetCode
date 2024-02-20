package main

import "sort"

// 也可使用map记录整数出现的次数，并用中间变量避免遍历改map，空间复杂度为O(n)
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
