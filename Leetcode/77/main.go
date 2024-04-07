package main

func main() {

}

func combine(n int, k int) [][]int {
	var result [][]int
	var path []int
	var backTracking func(n, k, startIndex int)
	backTracking = func(n, k, startIndex int) {
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := startIndex; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			backTracking(n, k, i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(n, k, 1)
	return result
}
