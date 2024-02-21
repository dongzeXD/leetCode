package main

import "fmt"

// 倒循环,正循环结果为取余运算
// 注:该解法收题目描述影响，模拟了真实循环运动过程，无法满足时间要求，应该推测运动后的结果
//func rotate_(nums []int, k int) {
//	length := len(nums)
//	if length < 2 {
//		return
//	}
//
//	k = length - (k % length)
//
//	for c := 0; c < k; c++ {
//		cache := nums[0]
//
//		for i := 0; i < length; i++ {
//			nextIDX := (i + length + 1) % length
//			nums[i] = nums[nextIDX]
//		}
//		nums[length-1] = cache
//	}
//}

// 找一下正反转搞不清楚的原因
func rotate(nums []int, k int) {
	length := len(nums)
	k = (length - k%length) % length
	n := append(nums[k:], nums[:k]...)
	fmt.Println(n)
	copy(nums, n)
	fmt.Println(nums)
}

func main() {

	//s := []int{-1, -100, 3, 99}

	s := []int{1, 2}
	// len = 7 K = 3
	// 【1，2，3，4，5，6，7，0 0 0 0 0 0 0】
	// 【 5，6，7，1，2，3，4】
	rotate(s, 3)
	fmt.Println(s)
}
