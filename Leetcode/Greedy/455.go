package main

import "sort"

// g-children

func main() {
	g := []int{1, 2, 3}
	s := []int{1, 2}
	println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	idx := len(s) - 1
	for i := len(g) - 1; i >= 0; i-- {
		if idx >= 0 && s[idx] >= g[i] {
			res++
			idx--
		}
	}
	return res
}
