package main

import (
	"bytes"
	"testing"
)

var testData = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

var testData2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func TestVeryBasicParsing(t *testing.T) {
	input := "drab violet bags contain 1 drab magenta bag, 5 posh orange bags, 1 mirrored brown bag, 4 wavy salmon bags."
	bags := make(map[string]map[string]int)
	parseBagRule(bags, input)

	if (bags["drab violet"]["drab magenta"] != 1) || (bags["drab violet"]["posh orange"] != 5) || (bags["drab violet"]["mirrored brown"] != 1) || (bags["drab violet"]["wavy salmon"] != 4) {
		t.Errorf("Something wrong inside here: %v\n", bags)
	}
}

func TestBagIt(t *testing.T) {
	buf := bytes.NewBufferString(testData)
	wantPt1, wantPt2 := 4, 32
	gotPt1, gotPt2 := bagIt(buf)

	if gotPt1 != wantPt1 {
		t.Errorf("Part 1 Got: %v, want: %v", gotPt1, wantPt1)
	}

	if gotPt2 != wantPt2 {
		t.Errorf("Part 2 Got: %v, want: %v", gotPt2, wantPt2)
	}
}

func TestBagIt2(t *testing.T) {
	buf := bytes.NewBufferString(testData2)
	wantPt1, wantPt2 := 0, 126
	gotPt1, gotPt2 := bagIt(buf)

	if gotPt1 != wantPt1 {
		t.Errorf("Part 1 Got: %v, want: %v", gotPt1, wantPt1)
	}

	if gotPt2 != wantPt2 {
		t.Errorf("Part 2 Got: %v, want: %v", gotPt2, wantPt2)
	}
}
