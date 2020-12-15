package main

import (
	"bytes"
	"testing"
)

func TestSumAllAnswers(t *testing.T) {
	tests := []struct {
		input  string
		output int
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
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString(test.input)
		got := sumAllAnswers(buf)

		if got != test.output {
			t.Errorf("Got: %v, want: %v", got, test.output)
		}
	}
}
