package main

import "fmt"

// 1 3 3 5 0 0 0   2 4 6
// 1 2 3 3 5 0 0   4 6
func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j, c int
	sorted := make([]int, 0, m+n)

	for ; c < m+n; c++ {
		if i == m {
			sorted = append(sorted, nums2[j:]...)
			break
		}
		if j == n {
			sorted = append(sorted, nums1[i:]...)
			break
		}
		if nums1[i] <= nums2[j] {
			sorted = append(sorted, nums1[i])
			i++
		} else {
			sorted = append(sorted, nums2[j])
			j++
		}
	}

	copy(nums1, sorted)
}

func rightDrift(src []int, left, right int) {
	for i := right; i > left; i-- {
		src[i] = src[i-1]
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
