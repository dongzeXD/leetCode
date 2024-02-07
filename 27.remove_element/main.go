package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1, 2, 3, 2, 6}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}
