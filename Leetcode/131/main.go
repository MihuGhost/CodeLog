package main

func main() {
	s := "aab"
	res := partition(s)
	println(res)
}

func partition(s string) [][]string {
	var result [][]string
	var path []string

	var dfs func(s string, startIndex int)
	dfs = func(s string, startIndex int) {
		//结束条件
		if startIndex == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			result = append(result, temp)
		}

		for i := startIndex; i < len(s); i++ {
			if isPalindrome(s[startIndex : i+1]) {
				path = append(path, s[startIndex:i+1])
			} else {
				continue
			}
			dfs(s, i+1)
			path = path[:len(path)-1]
		}

	}

	dfs(s, 0)
	return result
}
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
