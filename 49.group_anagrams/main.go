package main

import (
	"fmt"
	"maps"
	"slices"
)

func groupAnagrams2(strs []string) [][]string {
	baseMap := make(map[string][]string)

	for _, str := range strs {
		byteSlice := []byte(str)
		slices.Sort(byteSlice)
		baseMap[string(byteSlice)] = append(baseMap[string(byteSlice)], str)
	}
	result := make([][]string, 0)
	for _, slice := range baseMap {
		result = append(result, slice)
	}
	return result
}

func groupAnagrams(strs []string) [][]string {
	matchMapIndx := make(map[int]map[rune]int, len(strs))
	indexSlice := make([]int, len(strs))
	result := make([][]string, 0, len(strs))

	for idx, str := range strs {
		matchMap := make(map[rune]int)
		for _, ch := range str {
			if _, ok := matchMap[ch]; ok {
				matchMap[ch] = matchMap[ch] + 1
			} else {
				matchMap[ch] = 1
			}
		}
		matchMapIndx[idx] = matchMap
		indexSlice[idx] = idx
	}

	for len(indexSlice) > 0 {
		idx := indexSlice[0]
		baseMatch := matchMapIndx[idx]
		r := []string{strs[idx]}
		indexSlice = indexSlice[1:]
		for j := 0; j < len(indexSlice); {
			wantIdx := indexSlice[j]
			wantMatch := matchMapIndx[wantIdx]
			if maps.Equal(baseMatch, wantMatch) {
				r = append(r, strs[wantIdx])
				indexSlice = append(indexSlice[:j], indexSlice[j+1:]...)
			} else {
				j++
			}
		}
		result = append(result, r)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams2(strs))

	strs = []string{"zzzzzzy", "yzzzzzz"}
	fmt.Println(groupAnagrams2(strs))
}
