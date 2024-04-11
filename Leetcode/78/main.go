package main

func main() {}

func subsets(nums []int) [][]int {
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

		for i := startIndex; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(nums, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(nums, 0)
	return result
}
