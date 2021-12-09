package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func binaryToDecimal(number string) int {
	var result int = 0
	var bit int = 0
	var n int = len(number) - 1

	for n >= 0 {
		if number[n] == '1' {
			result += (1 << (bit))
		}
		n = n - 1
		bit++
	}

	return result
}

func main() {
	var gammaRate string = ""
	var epsilonRate string = ""

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	holder := [12][]int{}

	for _, line := range lines {
		c := strings.Split(line, "")
		for place, v := range c {
			i, _ := strconv.Atoi(v)
			holder[place] = append(holder[place], i)
		}
	}

	for _, positionArray := range holder {
		zero := 0
		one := 0
		for _, position := range positionArray {
			switch position {
			case 0:
				zero++
			case 1:
				one++
			}
		}
		if zero < one {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			epsilonRate += "1"
			gammaRate += "0"
		}
	}

	fmt.Println(gammaRate, epsilonRate)

	fmt.Println(binaryToDecimal(gammaRate) * binaryToDecimal(epsilonRate))
}

// [00100, 11110, 10110]

// [[0,1,1], [0,1,0], [1,1,1], [0,1,1], [0,0,0]]

// [0,1,1]
