package util

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type File []byte

type Lines []Line

type Line string

func ReadFile(path string) File {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return input
}

func (l Line) SubSplitWith(separator string) Lines {
	return split(string(l), separator)
}

func split(s, sep string) Lines {
	stringsArray := strings.Split(s, sep)
	lines := make([]Line, len(stringsArray))
	for i, s := range stringsArray {
		lines[i] = Line(s)
	}
	return lines
}

func (lines Lines) AsInts() []int {
	res := make([]int, len(lines))
	for i, l := range lines {
		res[i] = l.AsInt()
	}
	return res
}

func (l Line) AsInt() int {
	return ToInt(string(l))
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
