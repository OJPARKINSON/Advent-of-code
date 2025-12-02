package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the advent of code 2025 season!")

	// part1()

	part2()

}

func part1() {
	file, err := os.Open("./input1.txt")
	if err != nil {
		panic("failed to open input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dialMovements []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		dialMovements = append(dialMovements, parts...)
	}

	zeroCount := 0

	count := 50

	for _, movement := range dialMovements {
		direction := movement[0:1]
		moveAmount := movement[1:]

		if direction == "L" {
			amount, err := strconv.Atoi(moveAmount)
			if err != nil {
				fmt.Println("failed to convert int L: ", amount)
				continue
			}

			count = (count - amount + 100) % 100
		} else {
			amount, err := strconv.Atoi(moveAmount)
			if err != nil {
				fmt.Println("failed to convert int R: ", amount)
				continue
			}

			count = (count + amount) % 100
		}

		if count == 0 {
			zeroCount += 1
		}
	}
	fmt.Println(zeroCount)
}

func part2() {
	// dialMovements := []string{"L68",
	// 	"L30",
	// 	"R48",
	// 	"L5",
	// 	"R60",
	// 	"L55",
	// 	"L1",
	// 	"L99",
	// 	"R14",
	// 	"L82",
	// 	"R1000"}

	file, err := os.Open("./input1.txt")
	if err != nil {
		panic("failed to open input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var dialMovements []string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)

		dialMovements = append(dialMovements, parts...)
	}

	zeroCount := 0

	count := 50

	for _, movement := range dialMovements {
		direction := movement[0:1]
		moveAmount := movement[1:]

		amount, err := strconv.Atoi(moveAmount)
		if err != nil {
			fmt.Println("failed to convert to int: ", movement)
			continue
		}

		if direction == "L" {

			if amount > count {
				zeroCount += (amount - count + 99) / 100
			}
			count = (count - amount%100 + 100) % 100
		} else {
			zeroCount += (count + amount) / 100

			count = (count + amount) % 100
		}
	}
	fmt.Println(zeroCount)

}
