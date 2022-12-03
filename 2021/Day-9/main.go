package main

import (
	"bytes"
	"fmt"
	"github.com/ojparkinson/Advent-of-code/util"
	"strconv"
	"strings"
)

func part1(heightmap []int) int {

	return 0
}

func main() {
	var grid [][]int
	data := string(bytes.TrimSpace(util.ReadFile("input")))
	heightmap := strings.Split(data, "\n")
	risk := 0

	for _, line := range heightmap {
		var lineValues []int
		bonk := strings.Split(line, "")
		for _, v := range bonk {
			i, _ := strconv.Atoi(v)
			lineValues = append(lineValues, i)
		}
		grid = append(grid, lineValues)
	}

	for x, line := range grid {
		for y, height := range line {
			if (y > 0 && height >= line[y-1]) ||
				(y < len(line)-1 && height >= line[y+1]) ||
				(x > 0 && height >= grid[x-1][y]) ||
				(x < len(grid)-1 && height >= grid[x+1][y]) {
				continue
			}
			risk += height + 1
		}
	}
	fmt.Println(risk)

	//fmt.Println("Part 1: ", part1(heightmap))
	//fmt.Println("Part 2: ", part2(crabPositions))
}

//if (j > 0 && value >= line[j-1]) ||
//(j < len(line)-1 && value >= line[j+1]) ||
//(i > 0 && value >= grid[i-1][j]) ||
//(i < len(grid)-1 && value >= grid[i+1][j]) {
//continue
//}

//[2 1 9 9 9 4 3 2 1 0]
//[3 9 8 7 8 9 4 9 2 1]
//[9 8 5 6 7 8 9 8 9 2]
//[8 7 6 7 8 9 6 7 8 9]
//[9 8 9 9 9 6 5 6 7 8]
