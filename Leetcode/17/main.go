package main

func main() {

}

func letterCombinations(digits string) []string {
	numbers := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var result []string
	var path []byte

	var dfs func(digits string, index int)
	dfs = func(digits string, index int) {
		if index == len(digits) {
			temp := string(path)
			result = append(result, temp)
			return
		}

		digit := int(digits[index] - '0')
		letter := numbers[digit]
		for i := 0; i < len(letter); i++ {
			path = append(path, letter[i])
			dfs(digits, index+1)
			path = path[:len(path)-1]
		}
	}

	if digits == "" {
		return nil
	}

	dfs(digits, 0)
	return result
}
