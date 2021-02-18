package main

import (
	"bytes"
	"testing"
)

var testMap = struct {
	seatMap string
	part1   int
	part2   int
}{
	`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
	37,
	26,
}

func TestSeatMapFussy(t *testing.T) {
	buf := bytes.NewBufferString(testMap.seatMap)
	seatMap := parseData(buf)

	got := countOccupiedSeats(takeSeats(seatMap, 4, true))
	if got != testMap.part1 {
		t.Errorf("Part 1 - Got: %v, Want: %v\n", got, testMap.part1)
	}
}

func TestSeatMapNotFussy(t *testing.T) {
	buf := bytes.NewBufferString(testMap.seatMap)
	seatMap2 := parseData(buf)

	got := countOccupiedSeats(takeSeats(seatMap2, 5, false))
	if got != testMap.part2 {
		t.Errorf("Part 2 - Got: %v, Want: %v\n", got, testMap.part2)
	}
}
