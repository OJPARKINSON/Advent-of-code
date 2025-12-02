package main

import "testing"

func TestPart1(t *testing.T) {
	// given
	example := []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659",
		"824824821-824824827", "2121212118-2121212124"}

	// when
	result := part1(example)

	// then
	if result != 1227775554 {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, 1227775554)
	}
}

func TestPart2(t *testing.T) {
	// given
	example := []string{"11-22", "95-115", "998-1012", "1188511880-1188511890", "222220-222224",
		"1698522-1698528", "446443-446449", "38593856-38593862", "565653-565659",
		"824824821-824824827", "2121212118-2121212124"}

	// when
	result := part2(example)

	// then
	if result != 4174379265 {
		t.Errorf("Result was incorrect, got: %d, wanted: %d.", result, 4174379265)
	}
}
