package main

import (
	"bytes"
	"fmt"
	"github.com/ojparkinson/Advent-of-code/util"
	"strings"
)

type Cords struct {
	from struct {
		x, y int
	}
	to struct {
		x, y int
	}
}

func partOne(file util.File) int {
	inputFile := util.ReadFile("input")
	data := string(bytes.TrimSpace(inputFile))
	lines := strings.Split(strings.TrimSpace(data), "\n")
	vents := [1000][1000]int{}
	horAndVerLines := make([]Cords, 0)

	for _, line := range lines {
		cordsArray := strings.Split(line, " -> ")
		start := util.ExtractInts(cordsArray[0])
		end := util.ExtractInts(cordsArray[1])
		cord := Cords{from: struct {
			x, y int
		}{start[0], start[1]}, to: struct {
			x, y int
		}{end[0], end[1]}}

		if cord.from.x == cord.to.x {
			horAndVerLines = append(horAndVerLines, cord)
		} else if cord.from.y == cord.to.y {
			horAndVerLines = append(horAndVerLines, cord)
		}
	}
	for _, cord := range horAndVerLines {
		addX := 0
		addY := 0
		if cord.from.x > cord.to.x {
			addX = -1
		}
		if cord.from.x < cord.to.x {
			addX = 1
		}
		if cord.from.y > cord.to.y {
			addY = -1
		}
		if cord.from.y < cord.to.y {
			addY = 1
		}

		startX := cord.from.x
		startY := cord.from.y
		targetX := cord.to.x
		targetY := cord.to.y
		for startX != targetX || startY != targetY {
			vents[startY][startX] += 1
			startX += addX
			startY += addY
		}
		vents[startY][startX] += 1
	}
	sum := 0
	for i := range vents {
		for k := range vents[i] {
			if vents[i][k] >= 2 {
				sum += 1
			}
		}
	}

	return sum
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input")))
}
