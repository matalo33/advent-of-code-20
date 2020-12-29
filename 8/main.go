package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type opcode struct {
	instruction string
	num         int
	executed    bool
}

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	opcodesOrig := parseData(buf)
	opcodes := make([]opcode, len(opcodesOrig))

	copy(opcodes, opcodesOrig)
	part1, _ := runBootSequence(opcodes)
	fmt.Printf("Accumulator at time of loop: %v\n", part1)

	copy(opcodes, opcodesOrig)
	fmt.Printf("Accumulator after fixing boot: %v\n", fixBootSequence(opcodes))
}

func parseData(data *bytes.Buffer) (opcodes []opcode) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		opLine := strings.Split(scanner.Text(), " ")
		opArg, _ := strconv.Atoi(opLine[1])
		opcodes = append(opcodes, opcode{opLine[0], opArg, false})
	}
	return opcodes
}

func execute(op *opcode) (string, int, bool) {
	executedStash := op.executed
	op.executed = true
	return op.instruction, op.num, executedStash
}

func runBootSequence(opcodes []opcode) (int, bool) {
	ac, ptr := 0, 0
	for {
		instruction, num, executed := execute(&opcodes[ptr])
		if executed {
			return ac, false
		}
		switch op := instruction; op {
		case "acc":
			ac += num
			ptr++
		case "jmp":
			ptr += num
		case "nop":
			ptr++
		}
		if ptr >= len(opcodes) {
			return ac, true
		}
	}
}

func fixBootSequence(opcodes []opcode) int {
	for seq := 0; seq < len(opcodes); seq++ {
		if opcodes[seq].instruction == "jmp" || opcodes[seq].instruction == "nop" {
			testOpcodes := make([]opcode, len(opcodes))
			copy(testOpcodes, opcodes)
			if testOpcodes[seq].instruction == "jmp" {
				testOpcodes[seq].instruction = "nop"
			} else if testOpcodes[seq].instruction == "nop" {
				testOpcodes[seq].instruction = "jmp"
			}
			ac, result := runBootSequence(testOpcodes)
			if result {
				return ac
			}
		}
	}
	return 0
}
