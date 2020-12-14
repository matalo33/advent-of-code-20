package main

import "testing"

func TestReduce(t *testing.T) {
	tests := []struct {
		input  string
		max    int
		output int
	}{
		{
			"FBFBBFF",
			127,
			44,
		},
		{
			"RLR",
			7,
			5,
		},
	}

	for tn, test := range tests {
		result := reduce([]rune(test.input), 0, test.max)
		if result != test.output {
			t.Errorf("Test %v:, got: %v, want: %v", tn, result, test.output)
		}
	}
}

func TestCodeToId(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			"BFFFBBFRRR",
			567,
		},
		{
			"FFFBBBFRRR",
			119,
		},
		{
			"BBFFBBFRLL",
			820,
		},
	}

	for tn, test := range tests {
		result := codeToID(test.input)
		if result != test.output {
			t.Errorf("Test %v:, got: %v, want: %v", tn, result, test.output)
		}
	}
}
