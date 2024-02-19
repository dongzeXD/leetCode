package main

func removeDuplicates(nums []int) int {
	existMap := make(map[int]struct{})
	var left int
	for _, num := range nums {
		if _, ok := existMap[num]; !ok {
			existMap[num] = struct{}{}
			nums[left] = num
			left++
		}
	}
	return left
}
