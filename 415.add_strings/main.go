package main

import "strconv"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carryFlag := 0
	result := ""

	for i >= 0 || j >= 0 || carryFlag > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		add := x + y + carryFlag
		carryFlag = add / 10
		result = strconv.Itoa(add%10) + result
		i--
		j--
	}
	return result
}

func main() {

}
