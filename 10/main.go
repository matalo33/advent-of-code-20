package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	adapters := parseInput(buf)

	joltDistribution := findJoltDistribution(0, adapters)
	fmt.Printf("Jolt distribution: %v\n", joltDistribution)

	possibleCombinations := findPossibleCombinations(0, adapters)
	fmt.Printf("Possible combinations: %v\n", possibleCombinations)
}

func parseInput(data *bytes.Buffer) (adapters []int) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, num)
	}
	sort.Ints(adapters)
	return
}

func findJoltDistribution(startJolt int, adapters []int) int {
	jolt, jmpOneJolt, jmpThreeJolt := 0, 0, 0

	for i := 0; i < len(adapters); i++ {
		nextJoltJump := adapters[i] - jolt
		if nextJoltJump > 3 {
			break
		} else if nextJoltJump == 3 {
			jmpThreeJolt++
			jolt += 3
		} else if nextJoltJump == 2 {
			jolt += 2
		} else if nextJoltJump == 1 {
			jmpOneJolt++
			jolt++
		}
	}

	return jmpOneJolt * (jmpThreeJolt + 1)
}

func findPossibleCombinations(startJolt int, adapters []int) int {
	combinations := make(map[int]int)
	combinations[0] = 1
	for _, adapter := range adapters {
		combinations[adapter] = combinations[adapter-1] + combinations[adapter-2] + combinations[adapter-3]
	}
	return combinations[adapters[len(adapters)-1]]
}
