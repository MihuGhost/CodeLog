package main

func main() {}

func combinationSum3(k int, n int) [][]int {
	var result [][]int
	var path []int
	var dfs func(k, n, sum, startIndex int)
	dfs = func(k, n, sum, startIndex int) {
		if sum == n && len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := startIndex; i <= 9; i++ {
			if sum+i > n || 9-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(k, n, sum+i, i+1)
			path = path[:len(path)-1]
		}

	}
	dfs(k, n, 0, 1)
	return result
}
