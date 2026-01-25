package main

type NumArray struct {
	sumSlice []int
}

func Constructor(nums []int) NumArray {
	sumSlice := make([]int, len(nums)+1)
	for idx, num := range nums {
		sumSlice[idx+1] = sumSlice[idx] + num
	}
	return NumArray{sumSlice}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sumSlice[right+1] - this.sumSlice[left]
}
