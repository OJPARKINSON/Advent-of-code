package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	var lines []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	counter := 0

	for i := 0; i < len(lines)-3; i++ {
		value := lines[i] + lines[i+1] + lines[i+2]
		value2 := lines[i+1] + lines[i+2] + lines[i+3]
		if value < value2 {
			counter++
		}
	}

	fmt.Println(counter)

}
