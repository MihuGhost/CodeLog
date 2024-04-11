package main

import "sort"

func main() {}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var result [][]int
	var path []int

	var dfs func(nums []int, startIndex int)
	dfs = func(nums []int, startIndex int) {

		temp := make([]int, len(path))
		copy(temp, path)
		result = append(result, temp)

		if startIndex == len(nums) {
			return
		}
		//同层不可重复
		for i := startIndex; i < len(nums); i++ {
			if i != startIndex && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}

	}
	dfs(nums, 0)
	return result
}
