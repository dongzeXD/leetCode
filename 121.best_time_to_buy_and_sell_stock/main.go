package main

// 暴力
func maxProfit_(prices []int) int {
	money := 0

	for idx, buy := range prices {
		for i := idx + 1; i < len(prices); i++ {
			take := prices[i] - buy
			if take > money {
				money = take
			}
			money = take
		}
	}
	return money
}
