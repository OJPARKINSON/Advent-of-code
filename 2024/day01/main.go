package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	// Open the file
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Initialize slices for left and right values
	var left, right []int

	// Read each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		// Validate that we have exactly two numbers on each line
		if len(parts) != 2 {
			log.Fatalf("Invalid format. Each line must contain two space-separated numbers. Found: %s", line)
		}

		// Parse the numbers as integers
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("Error parsing numbers on line: %s", line)
		}

		// Append to respective lists
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Sort both lists
	sort.Ints(left)
	sort.Ints(right)

	// Calculate total distance using smallest pairing logic
	totalDistance := 0
	for i := range left {
		totalDistance += int(math.Abs(float64(left[i] - right[i])))
	}

	// Output the total distance
	fmt.Printf("Total distance: %d\n", totalDistance)
}

func part2() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Initialize slices for left and right values
	var left, right []int

	// Read each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		// Validate that we have exactly two numbers on each line
		if len(parts) != 2 {
			log.Fatalf("Invalid format. Each line must contain two space-separated numbers. Found: %s", line)
		}

		// Parse the numbers as integers
		leftNum, err1 := strconv.Atoi(parts[0])
		rightNum, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			log.Fatalf("Error parsing numbers on line: %s", line)
		}

		// Append to respective lists
		left = append(left, leftNum)
		right = append(right, rightNum)

		similarity := 0

		for _, num := range left {
			var counter int

			for _, num2 := range right {
				if num == num2 {
					counter += 1
				}
			}

			similarity += num * counter
		}

		fmt.Println("test", similarity)
	}

	fmt.Println(left)
	fmt.Println(right)
}

func main() {
	// part1()
	part2()
}
