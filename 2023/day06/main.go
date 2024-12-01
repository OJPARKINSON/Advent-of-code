package main

import "fmt"

func calculateWays(duration, record int) int {
	ways := 0

	for holdTime := 1; holdTime <= duration; holdTime++ {
		speed := holdTime
		travelTime := duration - holdTime
		if travelTime < 0 {
			break
		}
		distance := speed * travelTime
		if distance > record {
			ways++
		}
	}

	return ways
}

func main() {
	// Example input: durations and record distances
	races := []struct {
		duration int
		record   int
	}{
		{49, 356},
		{87, 1378},
		{78, 1502},
		{95, 1882},
	}

	totalWays := 1 // Start with 1 for multiplication
	for _, race := range races {
		ways := calculateWays(race.duration, race.record)
		totalWays *= ways // Multiply the number of ways for each race
	}
	fmt.Println("Total ways to win:", totalWays)
}
