package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	res := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range res {
			if res[i] == nil {
				spaces--
				if spaces == 0 {
					res[i] = person
					break
				}
			}
		}
	}
	return res
}

func main() {
	people := [][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}
	t := reconstructQueue(people)
	for _, t := range t {
		for _, v := range t {
			print(v)
		}
		println()
	}
}
