package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var games []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		games = append(games, scanner.Text())
	}

	// remove the start text
	for i, v := range games {
		stringArray := strings.Split(v, ":")
		games[i] = stringArray[1]
	}

	fmt.Println("pt1 answer: ", pt1(games))
	fmt.Println("pt2 answer: ", pt2(games))

}

func pt2(games []string) string {
	var gamesCount int64 = 0

	for _, v := range games {
		draw := strings.Split(v, " ")

		var red int64
		var green int64
		var blue int64

		for i, n := range draw {

			if strings.Contains(n, "red") && red == 0 {
				cube, err := strconv.ParseInt(draw[i-1], 0, 0)
				if err != nil {
					fmt.Println("error", err)
				}

				red = cube
			}

			if strings.Contains(n, "green") && green == 0 {
				cube, err := strconv.ParseInt(draw[i-1], 0, 0)
				if err != nil {
					fmt.Println("error", err)
				}

				green = cube
			}

			if strings.Contains(n, "blue") && blue == 0 {
				cube, err := strconv.ParseInt(draw[i-1], 0, 0)
				if err != nil {
					fmt.Println("error", err)
				}

				blue = cube
			}

		}

		count := red * green * blue
		gamesCount += count
	}

	return strconv.FormatInt(gamesCount, 10)
}

func pt1(games []string) string {
	var maxRed int64 = 12
	var maxGreen int64 = 13
	var maxBlue int64 = 14

	var count int

	for vi, v := range games {
		draw := strings.Split(v, " ")

		mainBreak := false
		for i, n := range draw {
			var localBreak bool = false

			localBreak = colourCheck(draw, i, n, "red", maxRed, localBreak)
			localBreak = colourCheck(draw, i, n, "green", maxGreen, localBreak)
			localBreak = colourCheck(draw, i, n, "blue", maxBlue, localBreak)

			if localBreak {
				mainBreak = true
				break
			}

		}

		if !mainBreak {
			indexCount := vi + 1
			count += indexCount
		}
	}

	return strconv.Itoa(count)
}

func colourCheck(draw []string, i int, n string, colour string, max int64, newBreak bool) bool {
	if strings.Contains(n, colour) {
		cube, err := strconv.ParseInt(draw[i-1], 0, 0)
		if err != nil {
			fmt.Println("error", err)
		}

		if cube > max {
			newBreak = true
		}
	}

	return newBreak
}
