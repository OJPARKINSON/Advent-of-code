package main

import (
	"fmt"
	"strings"

	"github.com/jdrst/adventofgo/util"
	"github.com/ojparkinson/Advent-of-code/util"
)

type BingoNumber struct {
	num    int
	marked bool
}
type Board [5][5]BingoNumber

func partOne(file util.File) int {
	lines := strings.Split(strings.TrimSpace(string(file)), "\n\n")
	nums := util.Line(lines[0]).SubSplitWith(",").AsInts()
	boards := createBoardsFrom(lines[1:])
	for _, num := range nums {
		for boardIndex := range boards {
			boards[boardIndex].MarkNum(num)
			if boards[boardIndex].HasWon() {
				return boards[boardIndex].CalculateScore(num)
			}
		}
	}
	return 0
}

func main() {
	fmt.Printf("First part: %v\n", partOne(util.ReadFile("input.txt")))
}

func (b *Board) CalculateScore(winNr int) int {
	sum := 0
	for _, r := range b {
		for _, n := range r {
			if !n.marked {
				sum += n.num
			}
		}
	}
	return sum * winNr
}

func (boards *Board) MarkNum(num int) {
	for i, r := range boards {
		for j, board := range r {
			if board.num == num {
				boards[i][j].marked = true
			}
		}
	}
}

func (boards Board) HasWon() bool {
nextRow:
	for rIndex, r := range boards {
		for _, board := range r {
			if !board.marked {
				goto cols
			}
		}
		return true
	cols:
		for cIndex := 0; cIndex < len(boards); cIndex++ {
			if !boards[cIndex][rIndex].marked {
				continue nextRow
			}
		}
		return true
	}
	return false
}

func createBoardsFrom(a []string) []Board {
	boards := make([]Board, len(a))

	for i, s := range a {
		var b Board
		for j, r := range strings.Split(s, "\n") {
			for k, n := range strings.Fields(r) {
				b[j][k] = BingoNumber{util.ToInt(n), false}
			}
		}
		boards[i] = b
	}

	return boards
}
