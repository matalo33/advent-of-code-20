package main

import (
	"testing"
)

func TestForValidPasswords(t *testing.T) {
	tests := []struct {
		input  []string
		output []int
	}{
		{
			[]string{
				"1-3 a: abcde",
				"1-3 b: cdefg",
				"2-9 c: ccccccccc",
			},
			[]int{
				2,
				1,
			},
		},
	}

	for _, test := range tests {
		for i := 0; i < len(test.output); i++ {
			result := countLegalPasswords(parsePasswords(test.input), i+1)
			if result != test.output[i] {
				t.Errorf("With policy %v: got %v, want %v", i+1, result, test.output)
			}
		}
	}
}
