package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input := readInput()
	fmt.Println(part1(input))
}

func part1(input [][]string) int {
	var beams []int

	for i, row := range input {
		if idx := slices.Index(row, "S"); idx != -1 {
			beams = append(beams, idx)
			input[i][idx] = "|"
			break
		}
	}

	splitCount := 0
	for row := 1; row < len(input); row++ {
		var newBeams []int

		for _, col := range beams {
			cell := input[row][col]

			if cell == "." {
				input[row][col] = "|"
				newBeams = append(newBeams, col)
			} else if cell == "^" {
				splitCount++
				if col-1 >= 0 {
					input[row][col-1] = "|"
					newBeams = append(newBeams, col-1)
				}
				if col+1 < len(input[row]) {
					input[row][col+1] = "|"
					newBeams = append(newBeams, col+1)
				}
			}
		}

		slices.Sort(newBeams)
		beams = slices.Compact(newBeams)
	}

	return splitCount
}
func part2(input [][]string) int {
	var beams []int

	for i, row := range input {
		if idx := slices.Index(row, "S"); idx != -1 {
			beams = append(beams, idx)
			input[i][idx] = "|"
			break
		}
	}

	splitCount := 0
	for row := 1; row < len(input); row++ {
		var newBeams []int

		for _, col := range beams {
			cell := input[row][col]

			if cell == "." {
				input[row][col] = "|"
				newBeams = append(newBeams, col)
			} else if cell == "^" {
				splitCount++
				if col-1 >= 0 {
					input[row][col-1] = "|"
					newBeams = append(newBeams, col-1)
				}
				if col+1 < len(input[row]) {
					input[row][col+1] = "|"
					newBeams = append(newBeams, col+1)
				}
			}
		}

		slices.Sort(newBeams)
		beams = slices.Compact(newBeams)
	}

	return splitCount
}

func readInput() [][]string {
	var inputGrid [][]string

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		inputGrid = append(inputGrid, strings.Split(scanner.Text(), ""))
	}

	return inputGrid
}
