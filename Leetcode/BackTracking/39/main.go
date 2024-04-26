package main

func main() {}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	var dfs func(candidates []int, target, sum, index int)
	dfs = func(candidates []int, target, sum, index int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(candidates, target, sum+candidates[i], i)
			path = path[:len(path)-1]
		}

	}
	dfs(candidates, target, 0, 0)
	return result
}
