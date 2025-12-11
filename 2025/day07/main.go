package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	input1 := readInput()
	fmt.Println(part1(input1))
	input2 := readInput()
	fmt.Println(part2(input2))
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
	beams := make(map[int]int)

	startRow := 0
	for i, row := range input {
		if idx := slices.Index(row, "S"); idx != -1 {
			beams[idx] = 1
			startRow = i
			break
		}
	}

	for row := startRow + 1; row < len(input); row++ {
		newBeams := make(map[int]int)

		for col, count := range beams {
			cell := input[row][col]

			if cell == "." {
				newBeams[col] += count
			} else if cell == "^" {
				if col-1 >= 0 {

					newBeams[col-1] += count
				}

				if col+1 < len(input[row]) {
					newBeams[col+1] += count
				}
			}
			fmt.Println(newBeams)
		}
		beams = newBeams
	}

	total := 0

	for _, count := range beams {
		fmt.Println(count)
		total += count
	}

	return total
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
