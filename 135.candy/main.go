package main

import (
	"fmt"
)

// 1 2 2 1 => 1 2 1 1  1 1 2 1
func candy(ratings []int) int {
	left := make([]int, len(ratings))
	right := make([]int, len(ratings))
	for idx, rating := range ratings {
		if idx > 0 && rating > ratings[idx-1] {
			left[idx] = left[idx-1] + 1
		} else {
			left[idx] = 1
		}
	}

	result := 0
	for j := len(ratings) - 1; j >= 0; j-- {
		if j < len(ratings)-1 && ratings[j] > ratings[j+1] {
			right[j] = right[j+1] + 1
		} else {
			right[j] = 1
		}
		result += max(right[j], left[j])
	}
	fmt.Println(left)
	fmt.Println(right)
	return result
}

func main() {
	fmt.Println(candy([]int{2, 2, 1}))
}
