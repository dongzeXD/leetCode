package main

import (
	"fmt"
	"reflect"
)

func twoSum(nums []int, target int) []int {
	targetMap := make(map[int]int)

	for firstIDX, num := range nums {
		if secondIDX, ok := targetMap[target-num]; ok {
			return []int{secondIDX, firstIDX}
		}
		targetMap[num] = firstIDX
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	toTargetMap := make(map[int]int)

	for idx, num := range nums {
		if firstIdx, ok := toTargetMap[num]; ok {
			return []int{firstIdx, idx}
		}
		want := target - num
		toTargetMap[want] = idx
	}
	return nil
}

func main() {
	result := twoSum2([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual([]int{1, 2}, result) {
		panic(fmt.Errorf("want %v but got %v", []int{1, 2}, result))
	}
}
