package main

func main() {}

func canCompleteCircuit(gas []int, cost []int) int {
	start, sum, minSum := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			start = i + 1
			minSum = sum
		}
	}

	if sum < 0 {
		return -1
	}

	if start == len(gas) {
		return 0
	} else {
		return start
	}
}
