package main

import (
	"bytes"
	"testing"
)

var testData = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestRunBootSequence(t *testing.T) {
	buf := bytes.NewBufferString(testData)
	want := 5
	got, _ := runBootSequence(parseData(buf))

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}

func TestFixBootSequence(t *testing.T) {
	buf := bytes.NewBufferString(testData)
	want := 8
	got := fixBootSequence(parseData(buf))

	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
