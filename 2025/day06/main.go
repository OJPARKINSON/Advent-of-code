package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input := readInput()

	fmt.Println("part1: ", part1(input))
	fmt.Println("part2: ", part2(input))

}

func part1(problems [][]string) int {
	grandTotal := 0

	trimmedProblems := make([][]string, len(problems))

	for i, p := range problems {
		trimmedProblems[i] = make([]string, 0)
		for _, nest := range p {
			if nest != " " && nest != "   " && len(nest) > 0 {
				trimmedProblems[i] = append(trimmedProblems[i], nest)
			}
		}
	}

	for col := 0; col < len(trimmedProblems[0]); col++ {
		op := trimmedProblems[len(trimmedProblems)-1][col]
		colTotal := 0
		for row := 0; row < len(trimmedProblems)-1; row++ {
			num, _ := strconv.Atoi(trimmedProblems[row][col])
			switch op {
			case "*":
				if row == 0 {
					colTotal = num
				} else {
					colTotal *= num
				}
			case "+":
				colTotal += num
			}
		}
		grandTotal = grandTotal + colTotal
	}
	return grandTotal
}

func part2(problems [][]string) int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if len(lines) == 0 {
		return 0
	}

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	grid := make([][]rune, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		for len(runes) < maxLen {
			runes = append(runes, ' ')
		}
		grid[i] = runes
	}

	grandTotal := 0
	var currentNums []int
	currentOp := ""

	for col := len(grid[0]) - 1; col >= 0; col-- {
		numStr := ""
		for row := 0; row < len(grid); row++ {
			char := grid[row][col]

			if unicode.IsDigit(char) {
				numStr += string(char)
			} else if char == '*' || char == '+' {
				currentOp = string(char)
			}
		}

		if numStr != "" {
			number, _ := strconv.Atoi(numStr)
			currentNums = append(currentNums, number)
		} else if len(currentNums) > 0 {
			result := currentNums[0]
			for i := 1; i < len(currentNums); i++ {
				if currentOp == "*" {
					result *= currentNums[i]
				} else {
					result += currentNums[i]
				}
			}
			fmt.Println("Problem:", currentNums, currentOp, "=", result)
			grandTotal += result
			currentNums = nil
			currentOp = ""
		}
	}

	if len(currentNums) > 0 {
		result := currentNums[0]
		for i := 1; i < len(currentNums); i++ {
			if currentOp == "*" {
				result *= currentNums[i]
			} else {
				result += currentNums[i]
			}
		}
		fmt.Println("Problem:", currentNums, currentOp, "=", result)
		grandTotal += result
	}

	return grandTotal
}

func readInput() [][]string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), " "))
	}

	return lines
}

// 123 328  51 64
//  45 64  387 23
//   6 98  215 314
// *   +   *   +
