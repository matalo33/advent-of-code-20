package main

import (
	"testing"
)

var testData = struct {
	earliestTime int
	timetable    string
	part1        int
}{
	939,
	"7,13,x,x,59,x,31,19",
	295,
}

func TestFindEarliestBusId(t *testing.T) {
	got := part1(findEarliestBusId(parseTimetable(testData.timetable), testData.earliestTime), testData.earliestTime)
	if got != testData.part1 {
		t.Errorf("Part 1, got: %v, want: %v\n", got, testData.part1)
	}
}