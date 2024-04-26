package main

import (
	"strconv"
	"strings"
)

func main() {}

func restoreIpAddresses(s string) []string {

	var result []string
	var path []string

	var dfs func(s string, startIndex int)
	dfs = func(s string, startIndex int) {
		if len(path) == 4 {
			if startIndex == len(s) {
				result = append(result, strings.Join(path, "."))
			}
		}

		for i := startIndex; i < len(s); i++ {
			if i != startIndex && s[startIndex] == '0' {
				break
			}
			num, _ := strconv.Atoi(s[startIndex : i+1])
			if num <= 255 && num >= 0 {
				path = append(path, s[startIndex:i+1])
				dfs(s, i+1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	dfs(s, 0)
	return result
}
