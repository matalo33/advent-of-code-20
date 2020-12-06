package main

import "testing"

func TestCountTrees(t *testing.T) {
	tests := []struct {
		input []string
		sets  []struct {
			routes []routes
			output int
		}
	}{
		{
			input: []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			},
			sets: []struct {
				routes []routes
				output int
			}{
				{
					routes: []routes{
						{1, 3},
					},
					output: 7,
				},
				{
					routes: ROUTES(),
					output: 336,
				},
			},
		},
	}

	for _, test := range tests {
		for _, set := range test.sets {
			result := combineRuns(makeMap(test.input), set.routes)
			if result != set.output {
				t.Errorf("Got: %v, want %v", result, set.output)
			}
		}
	}
}
