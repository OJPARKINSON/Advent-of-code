package main

import (
	"bytes"
	"fmt"
	"github.com/ojparkinson/Advent-of-code/util"
)

func part1(heightmap []int) int {

	return 0
}

func main() {
	data := string(bytes.TrimSpace(util.ReadFile("input")))
	heightmap := util.Line(data).SubSplitWith("\n").AsInts()

	fmt.Println("Part 1: ", part1(heightmap))
	//fmt.Println("Part 2: ", part2(crabPositions))
}
