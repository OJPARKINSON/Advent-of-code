package main

import (
	"bytes"
	"fmt"
	"github.com/ojparkinson/Advent-of-code/util"
)

func getFinalFishes(days int) int {
	data := string(bytes.TrimSpace(util.ReadFile("input")))
	fishes := util.Line(data).SubSplitWith(",").AsInts()

	queue := make([]int, 9)
	for _, fish := range fishes {
		queue[fish]++
	}

	for day := 0; day < days; day++ {
		currentFishes := queue[0]

		for j := 1; j < len(queue); j++ {
			queue[j-1] = queue[j]
		}

		queue[8] = currentFishes
		queue[6] += currentFishes
	}

	sum := 0
	for _, ele := range queue {
		sum += ele
	}

	return sum
}

func main() {
	fmt.Println("After 80 days: ", getFinalFishes(80))
	fmt.Println("After 256 days: ", getFinalFishes(256))
}
