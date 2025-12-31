package main

import "fmt"

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	m := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		if area > m {
			m = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return m
}

func main() {
	tS := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	r := maxArea(tS)
	fmt.Println(r)
}
