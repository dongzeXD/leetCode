package main

import (
	"fmt"
	"slices"
)

// 思路: 滑动窗口 但不满足需求 无法处理 0 1 1 2 3 输出4的预期
func longestConsecutive(nums []int) int {
	slices.Sort(nums)

	fmt.Println(nums)
	pre := 0
	start := 0
	m := 0

	for i := 0; i < len(nums); i++ {

		if i == 0 {
			pre = nums[i]
			m = 1
			continue
		}

		if pre+1 != nums[i] {
			start = i
		}
		pre = nums[i]
		if (i - start + 1) > m {
			m = i - start + 1
		}
	}

	return m
}

func longestConsecutive2(nums []int) int {
	slices.Sort(nums)
	result := make(map[int]map[int]struct{}) // next:pre

	for _, num := range nums {
		pre := num - 1
		if _, ok := result[num]; !ok {
			result[num] = make(map[int]struct{})
			result[num][num] = struct{}{}
		}
		if _, ok := result[pre]; ok {
			result[num] = result[pre]
			result[num][num] = struct{}{}
		}
	}
	fmt.Println(result)
	m := 0
	for _, item := range result {
		l := len(item)
		if l > m {
			m = l
		}
	}
	return m
}

func longestConsecutive3(nums []int) int {
	numSet := map[int]struct{}{}
	for _, num := range nums {
		numSet[num] = struct{}{}
	}
	longestStreak := 0
	for num := range numSet {
		if _, ok := numSet[num-1]; ok {
			continue
		}
		currentNum := num
		currentStreak := 1
		_, ok := numSet[num+1]
		for ok {
			currentNum++
			currentStreak++
			_, ok = numSet[currentNum+1]
		}
		if longestStreak < currentStreak {
			longestStreak = currentStreak
		}
	}
	return longestStreak
}

func main() {
	t := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive3(t))
}
