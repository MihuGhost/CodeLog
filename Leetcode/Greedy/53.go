package main

import "math"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	res := math.MinInt64
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if res < sum {
			res = sum
		}
		if sum <= 0 {
			sum = 0
		}
	}
	return res
}
