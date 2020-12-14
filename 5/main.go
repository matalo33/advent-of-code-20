package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	p1, p2 := findSeatIDs(input)

	fmt.Printf("Highest seat ID found: %v\n", p1)
	fmt.Printf("My seat ID: %v\n", p2)
}

func findSeatIDs(passes []string) (int, int) {
	highSeatID := 0
	seatList := make([]int, (128*8 + 8))
	for _, s := range passes {
		id := codeToID(s)
		seatList[id] = id
		if id > highSeatID {
			highSeatID = id
		}
	}

	// This feels wrong, but 🤷‍♂️
	// Sort the slice containing all seen seat IDs
	sort.Ints(seatList)

	// Find where the seats begin
	i := 0
	for s := range seatList {
		if seatList[s] != 0 {
			i = s
			break
		}
	}

	// Then iterate to find the missing seat
	lastSeenSeat := seatList[i]
	for _, s := range seatList[i:] {
		if s-lastSeenSeat > 1 {
			break
		}
		lastSeenSeat = s
	}

	return highSeatID, lastSeenSeat + 1
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
		}
		return max
	}

	// To ♾ and beyond
	if c == "F" || c == "L" {
		return reduce(s[1:], min, max-reduceBy)
	}
	return reduce(s[1:], min+reduceBy, max)
}
