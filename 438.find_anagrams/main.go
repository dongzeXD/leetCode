package main

import "fmt"

func findAnagrams(s string, p string) []int {
	findMap := make(map[byte]int)
	for _, ch := range p {
		matchCount := findMap[byte(ch)]
		findMap[byte(ch)] = matchCount + 1
	}

	result := make([]int, 0, len(p))

	srcSlice := []byte(s)
	lIdx, rIdx := 0, 0
	length := len(srcSlice)
	wantLength := len(p)
	valid := 0
	wantValid := len(findMap)

	for rIdx < length {
		// 右指针移动
		if c, ok := findMap[srcSlice[rIdx]]; ok {
			// c<0 匹配多了，左指针移动
			findMap[srcSlice[rIdx]] = c - 1
			if c-1 == 0 {
				valid += 1
			}
		}
		fmt.Println(">>>>>>右指针移动至: ", rIdx, "窗口字符为:", string(srcSlice[lIdx:rIdx+1]), "valid:", valid)
		rIdx++

		// 左指针移动并计算结果
		for rIdx-lIdx >= wantLength {
			if valid == wantValid {
				result = append(result, lIdx)
			}
			c, ok := findMap[srcSlice[lIdx]]
			if ok {
				findMap[srcSlice[lIdx]] = c + 1
				if c == 0 {
					valid -= 1
				}
			}
			lIdx += 1
			fmt.Println("<<<<<<<左指针移动至: ", lIdx, "窗口内字符为:", string(srcSlice[lIdx:rIdx]), "valid:", valid)
		}
	}
	return result
}

func main() {
	r := findAnagrams("baa", "aa")
	fmt.Println(r)
}
