package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//var input string

func main() {
	//input = strings.TrimRight(input, "\n")
	//calVal := parseInput(input)
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }

  var textArray []string

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {             // internally, it advances token based on sperator
	  fmt.Println(scanner.Text())  // token in unicode-char
	  textArray = append(textArray, scanner.Text())
  }

	p1Answer := part1(textArray)

	fmt.Println("Part 1:", p1Answer)
}


func part1(calVal []string) int {
	var builder strings.Builder
	var fullCount int
	for _, val := range calVal {
		var numberString []string
		for _, r := range val {
			if unicode.IsDigit(r) {
				builder.WriteRune(r)
				numberString = append(numberString, builder.String())
				builder.Reset()
			}
		}

		count, err := strconv.Atoi(numberString[0] + numberString[len(numberString) -1])
		if err != nil {
			fmt.Println("Error during conversion")
		}
		
		fullCount += count
	}
	return fullCount
}
