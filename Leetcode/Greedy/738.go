package main

import "strconv"

func main() {
	a := 10
	monotoneIncreasingDigits(a)
}

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	for i := len(ss) - 1; i > 0; i-- {
		if ss[i] < ss[i-1] {
			ss[i] = '9'
			ss[i-1]--
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
