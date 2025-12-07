package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	freshness, ingredients := readInput()

	fmt.Println("par1: ", part1(freshness, ingredients))

	fmt.Println("par2: ", part2(freshness))

}

func part1(freshness [][2]int, ingredients []int) int {
	totalFresh := len(ingredients)
	for _, ingredient := range ingredients {
		notInAnyFreshRange := 0
		for _, fresh := range freshness {
			if ingredient < fresh[0] || ingredient > fresh[1] {
				notInAnyFreshRange++
			}
		}

		if notInAnyFreshRange == len(freshness) {
			totalFresh--
		}
	}

	return totalFresh
}

func part2(freshness [][2]int) int {
	slices.SortFunc(freshness, func(a, b [2]int) int {
		if a[0] < b[0] {
			return -1
		} else if a[0] > b[0] {
			return 1
		}
		return 0
	})

	merged := [][2]int{freshness[0]}
	for _, r := range freshness[1:] {
		last := &merged[len(merged)-1]
		if r[0] <= last[1]+1 {
			if r[1] > last[1] {
				last[1] = r[1]
			}
		} else {
			merged = append(merged, r)
		}

	}
	total := 0
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}
	return total
}

func readInput() ([][2]int, []int) {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	fresh := [][2]int{}
	ingredients := []int{}

	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, "-") {
			freshSplit := strings.Split(text, "-")
			start, _ := strconv.Atoi(freshSplit[0])
			end, _ := strconv.Atoi(freshSplit[1])
			fresh = append(fresh, [2]int{start, end})
		} else if text != "" {
			ingredient, _ := strconv.Atoi(text)

			ingredients = append(ingredients, ingredient)
		}
	}
	return fresh, ingredients
}
