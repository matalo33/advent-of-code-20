package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	numbs := parseData(buf)

	invalidNum := findInvalidNum(numbs, 25)
	fmt.Printf("First invalid number: %v\n", invalidNum)
	encryptionWeakness := findEncryptionWeakness(numbs, invalidNum)
	fmt.Printf("Encryption weakness: %v\n", encryptionWeakness)
}

func parseData(data *bytes.Buffer) (numbs []int) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbs = append(numbs, num)
	}
	return
}

func findInvalidNum(numbs []int, preamble int) int {
	for pos := preamble; pos < len(numbs); pos++ {
		found := false
	out:
		for innerPos := pos - preamble; innerPos < pos-1; innerPos++ {
			for add := innerPos + 1; add < innerPos+preamble; add++ {
				if numbs[innerPos]+numbs[add] == numbs[pos] {
					found = true
					break out
				}
			}
		}
		if !found {
			return numbs[pos]
		}
	}
	return 0
}

func findEncryptionWeakness(numbs []int, target int) int {
	for pos := 0; pos < len(numbs); pos++ {
		acc := numbs[pos]
		for add := pos + 1; add < len(numbs); add++ {
			acc += numbs[add]
			if acc == target {
				min, max := minAndMax(numbs[pos : add+1])
				return min + max
			} else if acc > target {
				break
			}
		}
	}
	return 0
}

func minAndMax(in []int) (min int, max int) {
	min, max = in[0], in[0]
	for _, value := range in {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
