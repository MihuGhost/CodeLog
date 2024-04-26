package main

import (
	"sort"
)

func main() {
}

func findItinerary(tickets [][]string) []string {
	var result []string
	targets := make(map[string][]string, 0)

	for _, ticket := range tickets {
		targets[ticket[0]] = append(targets[ticket[0]], ticket[1])
	}

	for _, target := range targets {
		sort.Strings(target)
	}

	var dfs func(fromAirport string)
	dfs = func(fromAirport string) {
		for len(targets[fromAirport]) > 0 {
			toAirport := targets[fromAirport][0]
			targets[fromAirport] = targets[fromAirport][1:]
			dfs(toAirport)
		}
		result = append([]string{fromAirport}, result...)
	}
	dfs("JFK ")
	return result
}
