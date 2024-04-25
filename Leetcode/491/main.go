package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 1, 1}
	res := findSubsequences(nums)
	for _, subRes := range res {
		fmt.Println(subRes)
	}
}

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var path []int
	var dfs func(nums []int, startIdx int)
	dfs = func(nums []int, startIdx int) {

		if len(path) > 1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}

		used := make(map[int]bool, len(nums))
		for i := startIdx; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				used[nums[i]] = true
				path = append(path, nums[i])
				dfs(nums, i+1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(nums, 0)
	return res
}
