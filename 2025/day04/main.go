package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var examples [][]string

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		examples = append(examples, strings.Split(scanner.Text(), ""))
	}

	p1 := part1(examples)
	p2 := part2(examples)

	fmt.Println(p1)
	fmt.Println(p2)
}

func part1(input [][]string) int {
	rollsCanCollect := 0

	for row := 0; row < len(input); row++ {
		for col := 0; col < len(input); col++ {
			cell := input[row][col]

			if cell == "@" {
				neighbors := countNeighbors(row, col, input)
				if neighbors < 4 {
					rollsCanCollect++
				}
			}

		}
	}

	return rollsCanCollect
}

func gridPass(grid [][]string) int {
	// Find all removable rolls first (don't remove while iterating)
	toRemove := [][2]int{}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			cell := grid[row][col]

			if cell == "@" {
				neighbors := countNeighbors(row, col, grid)
				if neighbors < 4 {
					toRemove = append(toRemove, [2]int{row, col})
				}
			}
		}
	}

	// Now remove them
	for _, pos := range toRemove {
		grid[pos[0]][pos[1]] = "."
	}

	return len(toRemove)
}

func part2(grid [][]string) int {
	grandTotal := 0

	for {
		removed := gridPass(grid)
		if removed == 0 {
			break
		}
		grandTotal += removed
	}

	return grandTotal
}

func safeGet(grid [][]string, row, col int) string {
	if row < 0 || row >= len(grid) {
		return ""
	}

	if col < 0 || col >= len(grid[row]) {
		return ""
	}

	return grid[row][col]
}

func countNeighbors(row, col int, grid [][]string) int {
	neighborsCount := 0

	neighbors := []string{
		// top left
		safeGet(grid, row-1, col-1),
		// top
		safeGet(grid, row-1, col),
		// top right
		safeGet(grid, row-1, col+1),
		// left
		safeGet(grid, row, col-1),
		// right
		safeGet(grid, row, col+1),
		// bottom left
		safeGet(grid, row+1, col-1),
		// bottom
		safeGet(grid, row+1, col),
		// bottom right
		safeGet(grid, row+1, col+1),
	}

	for _, neighbor := range neighbors {
		if neighbor == "@" {
			neighborsCount++
		}
	}
	return neighborsCount
}
