package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	input = strings.TrimRight(input, "\n")
	elves := parseInput(input)

	totals := []int{}
	for _, items := range elves {
		var sum int
		for _, item := range items {
			sum += item
		}
		totals = append(totals, sum)
	}

	p1Answer := part1(totals)
	p2Answer := part2(totals)

	fmt.Println("Part 1:", p1Answer)
	fmt.Println("Part 2:", p2Answer)
}

func part2(totals []int) int {
	sort.Ints(totals)

	for i, j := 0, len(totals)-1; i < j; i, j = i+1, j-1 {
		totals[i], totals[j] = totals[j], totals[i]
	}

	top3Total := totals[0]
	top3Total += totals[1]
	top3Total += totals[2]

	return top3Total
}

func part1(totals []int) int {
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
