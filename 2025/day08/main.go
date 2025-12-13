package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	input, len := readInput()
	fmt.Println(part1(input, len))
}

func part1(pts [][3]int, n int) int {
	type pair struct {
		i, j int
		d    float64
	}
	var pairs []pair
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx, dy, dz := float64(pts[i][0]-pts[j][0]), float64(pts[i][1]-pts[j][1]), float64(pts[i][2]-pts[j][2])
			pairs = append(pairs, pair{i, j, math.Sqrt(dx*dx + dy*dy + dz*dz)})
		}
	}
	sort.Slice(pairs, func(a, b int) bool { return pairs[a].d < pairs[b].d })

	// Union-Find
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i], size[i] = i, 1
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	conns := min(1000, len(pairs))
	for i := 0; i < conns; i++ {
		a, b := find(pairs[i].i), find(pairs[i].j)
		if a != b {
			parent[a] = b
			size[b] += size[a]
		}
	}
	var sizes []int
	for i := range parent {
		if find(i) == i {
			sizes = append(sizes, size[i])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	result := 1
	for i := 0; i < min(3, len(sizes)); i++ {
		result *= sizes[i]
	}

	return result
}

func readInput() ([][3]int, int) {
	data, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	n := len(lines)

	pts := make([][3]int, n)
	for i, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &pts[i][0], &pts[i][1], &pts[i][2])
	}

	return pts, n
}
