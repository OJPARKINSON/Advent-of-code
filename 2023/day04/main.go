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

	fmt.Println("pt1 answer: ", pt1(getInput()))
	fmt.Println("pt2 answer: ", pt2(getInput()))

}

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var inputText []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputText = append(inputText, scanner.Text())
	}

	return inputText
}

func pt1(inputText []string) string {
	var gameCount int

	for i, line := range inputText {
		newLine := strings.Split(line, ": ")
		inputText[i] = newLine[1]
	}

	for _, line := range inputText {
		var cardCount int
		res := strings.Join(strings.Fields(line), " ")

		winnerAndMine := strings.Split(res, "|")

		winners := strings.Split(winnerAndMine[0], " ")
		mineArray := strings.Split(winnerAndMine[1], " ")

		for _, my := range mineArray {
			mine := strings.Replace(my, " ", "", -1)
			for _, win := range winners {
				if mine == win && mine != " " && win != " " && mine != "" && win != "" {
					if cardCount == 0 {
						cardCount = 1
					} else {
						cardCount *= 2
					}
				}
			}

		}

		gameCount += cardCount
	}

	return strconv.Itoa(gameCount)
}

func pt2(inputText []string) string {
	var cardCount int
	var copies []string

	for i, line := range inputText {
		newLine := strings.Split(line, ": ")
		inputText[i] = newLine[1]
	}

	copies, cardCount = scratchCardCheck(inputText, false, copies, cardCount)

	_, cardCount = scratchCardCheck(copies, true, copies, cardCount)

	return strconv.Itoa(cardCount)
}

func scratchCardCheck(inputText []string, copiesBool bool, copies []string, cardCount int) ([]string, int) {
	for lin, line := range inputText {
		res := strings.Join(strings.Fields(line), " ")

		winnerAndMine := strings.Split(res, "|")

		winners := strings.Split(winnerAndMine[0], " ")
		mineArray := strings.Split(winnerAndMine[1], " ")

		var matchingNumbers int
		for _, my := range mineArray {
			mine := strings.Replace(my, " ", "", -1)
			for _, win := range winners {
				if mine == win && mine != " " && win != " " && mine != "" && win != "" {
					matchingNumbers++
				}
			}

		}

		cardCount++
		if copiesBool {
			break
		}
		for i := 0; i < matchingNumbers; i++ {
			fmt.Println(lin, "adding")
			copies = append(inputText, inputText[i])
		}
	}

	return copies, cardCount
}
