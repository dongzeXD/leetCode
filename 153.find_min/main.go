package main

// [0,n)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1 // [0, n-1)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func main() {
	//sort.Search()
	// [6 1 2 3 4 5]
	// [3,4,5,1,2]
	findMin([]int{3, 4, 5, 6, 1})
}
