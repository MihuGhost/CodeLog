package main

func Swap(i, j []string) {
	i, j = j, i
}

func main() {
	tickets := [][]string{
		{"MUC", "LHR"},
		{"MUA", "LHR"},
		{"MUG", "LHR"},
	}

	for i := 1; i < len(tickets); i++ {
		if tickets[i-1][0] > tickets[i][0] {
			Swap(tickets[i-1], tickets[i])
		}
	}

}

/*func findItinerary(tickets [][]string) []string {

}*/
