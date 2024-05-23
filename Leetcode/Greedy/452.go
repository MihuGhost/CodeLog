package main

import "sort"

func main() {
	println(findMinArrowShots([][]int{
		{9, 12},
		{1, 10},
		{4, 11},
		{8, 12},
		{3, 9},
		{6, 9},
		{6, 7},
	}))
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		a, b := points[i], points[j]
		return a[1] < b[1] || a[1] == b[1] && a[0] < b[0]
	})
	res := 1
	last := points[0][1]
	for i := 1; i < len(points); i++ {
		if last < points[i][0] {
			last = points[i][1]
			res++
		}
	}
	return res
}
