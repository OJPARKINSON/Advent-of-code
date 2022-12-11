package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimRight(input, "\n")

	p1Answer := part1(input)

	fmt.Println("Part 1:", p1Answer)
}

func part1(input string) int {
	elves := parseInput(input)

	totals := []int{}
	for _, items := range elves {
		var sum int
		for _, item := range items {
			sum += item
		}
		totals = append(totals, sum)
	}

	biggestNum := totals[0]
	for _, num := range totals {
		if num > biggestNum {
			biggestNum = num
		}
	}

	return biggestNum
}

func parseInput(input string) (ans [][]int) {
	for _, group := range strings.Split(input, "\n\n") {
		row := []int{}
		for _, line := range strings.Split(group, "\n") {
			intLine, _ := strconv.Atoi(line)
			row = append(row, intLine)
		}
		ans = append(ans, row)
	}
	return ans
}
