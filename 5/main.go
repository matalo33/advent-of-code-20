package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Highest seat ID found: %v\n", findHighSeatID(input))
}

func findHighSeatID(passes []string) int {
	highSeatID := 0
	for _, s := range passes {
		id := codeToID(s)
		if id > highSeatID {
			highSeatID = id
		}
	}
	return highSeatID
}

func codeToID(s string) int {
	c := []rune(s)
	return reduce(c[0:6], 0, 127)*8 + reduce(c[7:9], 0, 7)
}

func reduce(s []rune, min, max int) int {
	c := string(s[0:1])
	reduceBy := ((max - min) + 1) / 2

	// Time to 🚪
	if reduceBy == 1 {
		if c == "F" || c == "L" {
			return min
		} else {
			return max
		}
	}

	// To ♾ and beyond
	if c == "F" || c == "L" {
		return reduce(s[1:], min, max-reduceBy)
	} else {
		return reduce(s[1:], min+reduceBy, max)
	}
}
