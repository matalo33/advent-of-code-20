package main

import (
	"bytes"
	"testing"
)

var testData = struct {
	input string
	part1 int
	part2 int
}{
	`F10
N3
F7
R450
F11`,
	25,
	286,
}

func TestNavigate(t *testing.T) {
	buf := bytes.NewBufferString(testData.input)
	instructions := parseData(buf)
	got := manhattanDistance(navigate(instructions, 90))
	if got != testData.part1 {
		t.Errorf("Part 1 - Got: %v, Want: %v", got, testData.part1)
	}
}

func TestNavigateWithWaypoint(t *testing.T) {
	buf := bytes.NewBufferString(testData.input)
	instructions := parseData(buf)
	got := manhattanDistance(navigateWithWaypoint(instructions, 1, 10))
	if got != testData.part2 {
		t.Errorf("Part 2 - Got : %v, Want: %v", got, testData.part2)
	}
}

func TestRotateRight(t *testing.T) {
	var tests = []struct {
		d int
		e int
		n int
	}{
		{90, 1, -2},
		{180, -2, -1},
		{270, -1, 2},
	}
	for _, test := range tests {
		gotX, gotY := rotateRight(2, 1, test.d)
		if gotX != test.e || gotY != test.n {
			t.Errorf("Expected: %v, %v, Rotating: %v, Got: %v, %v", test.e, test.n, test.d, gotX, gotY)
		}
	}
}

func TestRotateLeft(t *testing.T) {
	var tests = []struct {
		d int
		e int
		n int
	}{
		{90, -1, 2},
		{180, -2, -1},
		{270, 1, -2},
	}
	for _, test := range tests {
		gotX, gotY := rotateLeft(2, 1, test.d)
		if gotX != test.e || gotY != test.n {
			t.Errorf("Expected: %v, %v, Rotating: %v, Got: %v, %v", test.e, test.n, test.d, gotX, gotY)
		}
	}
}
