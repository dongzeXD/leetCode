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

func main() {
	result := twoSum([]int{3, 2, 4}, 6)
	if !reflect.DeepEqual([]int{1, 2}, result) {
		panic(fmt.Errorf("want %v but got %v", []int{1, 2}, result))
	}
}
