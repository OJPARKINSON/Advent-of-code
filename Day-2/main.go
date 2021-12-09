package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var horizontalPosition int = 0
	var depth int = 0
	var aim int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(lines); i++ {
		direction := strings.Split(lines[i], " ")
		directionAmount, _ := strconv.Atoi(direction[1])
		switch direction[0] {
		case "forward":
			horizontalPosition += directionAmount
			depth += aim * directionAmount
		case "up":
			aim -= directionAmount
		case "down":
			aim += directionAmount
		}
	}

	fmt.Println(horizontalPosition * depth)
}
