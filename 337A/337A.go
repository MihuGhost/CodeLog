package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// number of students, puzzles
	var student int
	var num int
	n, err := fmt.Scanln(&student, &num)
	if err != nil {
		fmt.Println(err)
	}
	if n != 2 {
		fmt.Println("Error: Incorrect number of elements entered.")
		return
	}
	// pieces of puzzle
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputString := scanner.Text()
	puzzle := make([]int, num)
	elements := strings.Fields(inputString)
	if len(elements) != num {
		fmt.Println("Error: Incorrect number of elements entered.")
		return
	}
	for i := 0; i < num; i++ {
		element, err := strconv.Atoi(elements[i])
		if err != nil {
			fmt.Println("Error converting input to integers:", err)
			return
		}
		puzzle[i] = element
	}
	fmt.Println(puzzles(student, puzzle))
}

func puzzles(student int, puzzle []int) int {
	SortArray(puzzle)
	res := 1000
	for i := 0; i < len(puzzle)-student+1; i++ {
		sub := puzzle[i+student-1] - puzzle[i]
		if res > sub {
			res = sub
		}
	}
	return res
}

func SortArray(puzzle []int) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle)-1-i; j++ {
			if puzzle[j] > puzzle[j+1] {
				puzzle[j+1], puzzle[j] = puzzle[j], puzzle[j+1]
			}
		}
	}
}
