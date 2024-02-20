package main

func removeDuplicates(nums []int) int {
	existMap := make(map[int]int)
	var left int
	for _, num := range nums {
		if count, ok := existMap[num]; !ok {
			existMap[num] = 1
			nums[left] = num
			left++
		} else if count < 2 {
			existMap[num] = count + 1
			nums[left] = num
			left++
		}
	}
	return left
}

// 力扣官方题解
//func removeDuplicates(nums []int) int {
//	n := len(nums)
//	if n <= 2 {
//		return n
//	}
//	slow, fast := 2, 2
//	for fast < n {
//		if nums[slow-2] != nums[fast] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//		fast++
//	}
//	return slow
//}
