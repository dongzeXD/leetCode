package main

import "fmt"

func minWindow(s string, t string) string {
	wantMap := make(map[byte]int)
	for _, b := range []byte(t) {
		wantMap[b] += 1
	}

	left, right, valid := 0, 0, 0
	result := ""

	for right < len(s) {
		char := s[right]
		if w, ok := wantMap[char]; ok {
			wantMap[char] = w - 1
			if w-1 == 0 {
				valid++
			}
		}
		right++

		for valid == len(wantMap) {
			tmp := s[left:right]
			if len(tmp) < len(result) || result == "" {
				result = tmp
			}

			char = s[left]
			if w, ok := wantMap[char]; ok {
				wantMap[char] = w + 1
				if w == 0 {
					valid--
				}
			}
			left++
		}
	}
	return result
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}
