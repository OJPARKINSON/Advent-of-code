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

	var inputText []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText = append(inputText, scanner.Text())
	}
	fmt.Println("pt1 answer: ", pt1(inputText))
	// fmt.Println("pt2 answer: ", pt2(games))

}

func pt1(inputText []string) string {

	for i, line := range inputText {
		newLine := strings.Split(line, ": ")
		inputText[i] = newLine[1]
	}

	var count int

	for _, line := range inputText {

		fmt.Println(line)

		winnerAndMine := strings.Split(line, "|")

		// winners := strings.Split(winnerAndMine[0], " ")
		mine := strings.Split(winnerAndMine[1], " ")

		fmt.Println(mine)
		for _, my := range mine {
			newmy := strings.Replace(my, " ", "", -1)

			if strings.Contains(winnerAndMine[0], newmy) {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}
