package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var lIDX, rIDX int
	inMap := make(map[uint8]struct{})
	maxLen := 0
	for lIDX <= rIDX && rIDX != len(s) {
		var rptChar uint8
		// 向右扩张
		for rIDX < len(s) {
			_, ok := inMap[s[rIDX]]
			inMap[s[rIDX]] = struct{}{}
			if ok {
				rptChar = s[rIDX]
				break
			} else {
				rIDX++
			}
		}

		length := rIDX - lIDX
		if length > maxLen {
			maxLen = length
		}

		// 左侧收缩
		for lIDX < len(s) {
			lChar := s[lIDX]
			lIDX++
			delete(inMap, lChar)
			if lChar == rptChar {
				break
			}
		}

	}
	return maxLen
}

func main() {
	r := lengthOfLongestSubstring(" ")
	fmt.Println(r)
}
