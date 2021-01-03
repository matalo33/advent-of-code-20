package main

import (
	"bytes"
	"testing"
)

func TestAdapters(t *testing.T) {
	testData := []struct {
		input string
		part1 int
		part2 int
	}{
		{
			`16
10
15
5
1
11
7
19
6
12
4
`,
			35,
			8,
		},
		{
			`28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`,
			220,
			19208,
		},
	}

	for _, test := range testData {
		buf := bytes.NewBufferString(test.input)
		adapters := parseInput(buf)

		joltDistribution := findJoltDistribution(0, adapters)
		if joltDistribution != test.part1 {
			t.Errorf("Jolt distribution want: %v, got: %v\n", test.part1, joltDistribution)
		}

		possibleCombinations := findPossibleCombinations(0, adapters)
		if possibleCombinations != test.part2 {
			t.Errorf("Possible combinations want: %v, got: %v\n", test.part2, possibleCombinations)
		}

	}
}
