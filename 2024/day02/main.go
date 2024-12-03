package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {
	if len(report) < 2 {
		return false
	}

	increasing := true
	decreasing := true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}

func part1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read each line
	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		var report []int
		for _, part := range parts {
			cur, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Invalid number in report:", part)
				return
			}
			report = append(report, cur)
		}
		reports = append(reports, report)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}

	fmt.Println("Part 1: ", safeCount)
}

func isSafe2(report []int) bool {
	increasing, decreasing := true, true

	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if diff > 0 {
			decreasing = false
		} else if diff < 0 {
			increasing = false
		}
	}

	return increasing || decreasing
}

func canBeMadeSafe(report []int) bool {
	n := len(report)
	for skip := 0; skip < n; skip++ {
		increasing, decreasing := true, true
		prev := -1 // Previous valid index

		for i := 0; i < n; i++ {
			if i == skip {
				continue // Skip the current level
			}

			if prev != -1 { // Compare with the previous valid level
				diff := report[i] - report[prev]
				if diff < -3 || diff > 3 || diff == 0 {
					increasing, decreasing = false, false
					break
				}
				if diff > 0 {
					decreasing = false
				} else if diff < 0 {
					increasing = false
				}
			}

			prev = i // Update the previous valid index
		}

		if increasing || decreasing {
			return true
		}
	}
	return false
}

func part2() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		levels := strings.Fields(line)
		var report []int
		for _, level := range levels {
			num, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Invalid number in report:", level)
				return
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	safeCount := 0
	for _, report := range reports {
		if isSafe(report) || canBeMadeSafe(report) {
			safeCount++
		}
	}

	fmt.Println("part 2: ", safeCount)
}

func main() {
	// part1()
	part2()
}
