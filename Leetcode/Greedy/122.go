package main

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(prices))
}
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}
