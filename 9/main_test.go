package main

import (
	"bytes"
	"testing"
)

var testData = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestFindInvalidNum(t *testing.T) {
	buf := bytes.NewBufferString(testData)
	want := 127
	numbs := parseData(buf)
	got := findInvalidNum(numbs, 5)

	if got != want {
		t.Errorf("Finding invalid number, Got: %v, want %v\n", got, want)
	}

	encryptionWeakness := findEncryptionWeakness(numbs, 127)
	wantWeakness := 62

	if encryptionWeakness != wantWeakness {
		t.Errorf("Finding weakness, got: %v, want: %v\n", encryptionWeakness, wantWeakness)
	}
}
