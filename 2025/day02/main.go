package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("let's go")

	file, _ := os.Open("./input1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	inputLine := strings.TrimSpace(scanner.Text())

	commaSplitInput := strings.Split(inputLine, ",")

	part1Result := part1(commaSplitInput)

	part2Result := part2(commaSplitInput)

	fmt.Println("part1Result ", part1Result)
	fmt.Println("part2Result ", part2Result)
}

func part1(ranges []string) int {
	count := 0
	for _, input := range ranges {
		splitInput := strings.Split(input, "-")

		input1, _ := strconv.Atoi(splitInput[0])
		input2, _ := strconv.Atoi(splitInput[1])

		for i := input1; i <= input2; i++ {
			iAsString := strconv.Itoa(i)
			if iAsString[len(iAsString)/2:] == iAsString[:len(iAsString)/2] {
				count += i
			}
		}
	}

	return count
}

func part2(ranges []string) int {
	total := 0

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			length := len(str)

			for patternLen := 1; patternLen <= length/2; patternLen++ {
				if length%patternLen != 0 {
					continue
				}

				pattern := str[:patternLen]
				repetitions := length / patternLen

				if strings.Repeat(pattern, repetitions) == str {
					total += i
					break
				}
			}
		}
	}

	return total
}
