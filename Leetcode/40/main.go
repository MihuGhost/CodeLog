package main

import "sort"

func main() {}

// 排序剪枝
func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var path []int

	var dfs func(candidates []int, target, sum, index int)
	dfs = func(candidates []int, target, sum, index int) {
		if sum > target {
			return
		}

		if target == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := index; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			dfs(candidates, target, sum+candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}

	sort.Ints(candidates)
	dfs(candidates, target, 0, 0)
	return result
}
