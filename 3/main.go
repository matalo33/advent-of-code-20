package main

import (
	"bufio"
	"fmt"
	"os"
)

type routes struct {
	y int
	x int
}

func ROUTES() []routes {
	return []routes{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	tMap := makeMap(input)

	fmt.Printf("Trees encountered on first route: %v\n", countTrees(tMap, 1, 3))
	fmt.Printf("Trees encountered on other toute: %v\n", combineRuns(tMap, ROUTES()))
}

func makeMap(input []string) [][]bool {
	tMap := make([][]bool, len(input))
	for y, line := range input {
		tMap[y] = make([]bool, len(line))
		for x, cell := range line {
			if cell == '#' {
				tMap[y][x] = true
			}
		}
	}
	return tMap
}

func countTrees(tMap [][]bool, steerDown int, steerRight int) int {
	x, trees := 0, 0
	for y := 0; y < len(tMap); y += steerDown {
		if x/len(tMap[y]) != 0 {
			x = x % len(tMap[y])
		}
		if tMap[y][x] {
			trees++
		}
		x += steerRight
	}
	return trees
}

func combineRuns(tMap [][]bool, routes []routes) int {
	total := 1
	for _, route := range routes {
		total = total * countTrees(tMap, route.y, route.x)
	}
	return total
}
