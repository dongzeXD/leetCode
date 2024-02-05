package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j int

	for i < m+n {
		// nums1 为空
		if i == m+j {
			for ; j < n; j++ {
				nums1[i] = nums2[j]
				i++
			}
		} else if j == n { // nums2 为空
			return
		} else {
			if nums2[j] < nums1[i] {
				rightDrift(nums1, i, m+j)
				nums1[i] = nums2[j]
				j++
			}
			i++
		}
	}
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
