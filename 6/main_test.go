package main

import (
	"bytes"
	"testing"
)

func TestSumAllAnswers(t *testing.T) {
	tests := []struct {
		input string
		part1 int
		part2 int
	}{
		{
			`abc

a
b
c

ab
ac

a
a
a
a

b
`,
			11,
			6,
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString(test.input)
		p1, p2 := sumAnswers(buf)

		if p1 != test.part1 {
			t.Errorf("Got: %v, want: %v", p1, test.part1)
		}

		if p2 != test.part2 {
			t.Errorf("Got %v, want: %v", p2, test.part2)
		}
	}
}
