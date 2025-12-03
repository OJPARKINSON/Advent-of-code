package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// batteries := []string{
	// 	"987654321111111",
	// 	"811111111111119",
	// 	"234234234234278",
	// 	"818181911112111",
	// }

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var batteries []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		batteries = append(batteries, parts...)
	}

	outputJoltage := 0
	for _, battery := range batteries {
		joltageArr := strings.Split(battery, "")
		joltageIntArr := make([]int, len(joltageArr))

		for i := range joltageArr {
			joltageIntArr[i], _ = strconv.Atoi(joltageArr[i])
		}

		maxJoltage := 0
		for i := 0; i < len(joltageIntArr)-1; i++ {
			for j := i + 1; j < len(joltageIntArr); j++ {
				joltage := joltageIntArr[i]*10 + joltageIntArr[j]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		outputJoltage += maxJoltage
	}

	fmt.Println("part1 ", outputJoltage)
}
