package main

import (
	"bytes"
	"fmt"
	"github.com/ojparkinson/Advent-of-code/util"
	"math"
)

func part1(crabPositions []int) int {
	var largest int = 0
	var smallest int = 0

	largest = crabPositions[0]
	for i := 0; i < len(crabPositions); i++ {
		if largest < crabPositions[i] {
			largest = crabPositions[i]
		}
	}

	differences := make([]int, largest+1)
	for i, difference := range differences {
		for _, position := range crabPositions {
			var diff int
			if position-i > 0 {
				diff = position - i
			} else {
				diff = i - position
			}
			difference += diff
		}

		if i == 0 || smallest > difference {
			smallest = difference
		}
	}

	return smallest
}

func part2(crabPositions []int) int {
	var largest int = 0
	var smallest int = 0

	largest = crabPositions[0]
	for i := 0; i < len(crabPositions); i++ {
		if largest < crabPositions[i] {
			largest = crabPositions[i]
		}
	}

	differences := make([]int, largest+1)
	for i, difference := range differences {
		for _, position := range crabPositions {
			steps := int(math.Abs(float64(position - i)))
			difference += steps * (steps+1)/2
		}

		if i == 0 || smallest > difference {
			smallest = difference
		}
		differences[i] = difference
	}

	return smallest
}

func main() {
	data := string(bytes.TrimSpace(util.ReadFile("input")))
	crabPositions := util.Line(data).SubSplitWith(",").AsInts()

	fmt.Println("Part 1: ", part1(crabPositions))
	fmt.Println("Part 2: ", part2(crabPositions))
}
