package main

import "fmt"

func moveZeroes(nums []int) {
	length := len(nums)
	for i := 0; i < len(nums); {
		for j := i + 1; j < len(nums); j++ {
			if nums[j-1] == 0 {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
		nums = nums[:len(nums)-1]
	}
	nums = nums[:length]
}

// 没有统一逻辑，想到哪写到哪
func moveZero2(nums []int) {
	length := len(nums)
	c := 0
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			c += 1
		} else {
			i++
		}
	}

	fmt.Println(nums)
	nowL := len(nums)
	nums = nums[:length]
	for i := 0; i < c; i++ {
		nums[nowL+i] = 0
	}
	fmt.Println(nums)
}

// 双指针交换思想
func moveZero3(nums []int) {
	l, r, n := 0, 0, len(nums)
	for r < n {
		if nums[r] != 0 {
			if r != l {
				nums[l], nums[r] = nums[r], nums[l]
			}
			l++
		}
		r++
	}
	fmt.Println(nums)
}

func main() {
	moveZero3([]int{-1, 1, 0, 3, 12})
}
