package main

func main() {}

func permute(nums []int) [][]int {
	var result [][]int
	var path []int
	usedNums := make([]int, len(nums))

	var dfs func(nums []int, usedNums []int)
	dfs = func(nums []int, usedNums []int) {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		for i := 0; i < len(nums); i++ {
			if usedNums[i] == 1 {
				continue
			}
			usedNums[i] = 1
			path = append(path, nums[i])
			dfs(nums, usedNums)
			usedNums[i] = 0
			path = path[:len(path)-1]
		}
	}
	dfs(nums, usedNums)
	return result
}
