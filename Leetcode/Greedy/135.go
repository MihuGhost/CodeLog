package main

func main() {
	println(candy([]int{1, 0, 2}))
}
func candy(ratings []int) int {
	res := 0
	candys := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		candys[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i] < ratings[i+1] {
			candys[i+1] = candys[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candys[i-1] = max(candys[i-1], candys[i]+1)
		}
	}

	for i := 0; i < len(ratings); i++ {
		res += candys[i]
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
