package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println(part1(input))
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
		fmt.Println("---")
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
