package main

import "fmt"

func spiralArray(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	n := 0
	s := len(array) - 1
	w := 0
	e := len(array[0]) - 1
	dst := true
	result := make([]int, 0, len(array))
	for n <= s && w <= e {
		if dst {
			for i := w; i <= e; i++ {
				fmt.Println(array[n][i])
				result = append(result, array[n][i])
			}
			n++
			for i := n; i <= s; i++ {
				fmt.Println(array[i][e])
				result = append(result, array[i][e])
			}
			e--

		} else {
			for i := e; i >= w; i-- {
				fmt.Println(array[s][i])
				result = append(result, array[s][i])
			}
			s--
			for i := s; i >= n; i-- {
				fmt.Println(array[i][w])
				result = append(result, array[i][w])
			}
			w++
		}
		dst = !dst
	}
	return result
}

func main() {
	//result := spiralArray([][]int{
	//	{1, 2, 3},
	//	{8, 9, 4},
	//	{7, 6, 5},
	//})

	//result := spiralArray([][]int{})

	result := spiralArray([][]int{
		{7},
		{9},
		{6},
	})
	fmt.Println(result)

}
