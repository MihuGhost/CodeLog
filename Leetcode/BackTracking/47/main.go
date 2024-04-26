package main

import "sort"

func main() {}

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var path []int
	usedNums := make([]bool, len(nums))

	var dfs func(nums []int, usedNums []bool)
	dfs = func(nums []int, usedNums []bool) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}
		//https://www.programmercarl.com/0047.%E5%85%A8%E6%8E%92%E5%88%97II.html#%E6%80%9D%E8%B7%AF
		//参考树层图中usedNums数组
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && usedNums[i-1] == false {
				continue
			}
			if !usedNums[i] {
				usedNums[i] = true
				path = append(path, nums[i])
				dfs(nums, usedNums)
				usedNums[i] = false
				path = path[:len(path)-1]
			}
		}
	}

	sort.Ints(nums)
	dfs(nums, usedNums)
	return result
}
