package main

import "fmt"

// [3,4,−1,1]
func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; {
		if nums[i] < 1 || nums[i] > length {
			i++
			continue
		}

		currentValue := nums[i]
		correctValue := i + 1
		if currentValue == correctValue {
			i++
			continue
		}
		correctValue = nums[currentValue-1]
		if correctValue == currentValue {
			i++
			continue
		}
		nums[currentValue-1], nums[i] = nums[i], nums[currentValue-1]
	}

	for idx, num := range nums {
		if idx+1 != num {
			return idx + 1
		}
	}
	return length + 1
}

func main() {
	s := []int{7, 8, 9, 11, 12}
	r := firstMissingPositive(s)
	fmt.Println(r, s)

}
