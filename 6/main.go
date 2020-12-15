package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	fmt.Printf("Sum of all answers: %v\n", sumAllAnswers(buf))
}

func sumAllAnswers(data *bytes.Buffer) int {
	scanner := bufio.NewScanner(data)
	sumAllAnswers, answersSet := 0, ""

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sumAllAnswers += countAnswers(answersSet)
			answersSet = ""
		} else {
			answersSet += line
		}
	}
	return sumAllAnswers + countAnswers(answersSet)
}

func countAnswers(answers string) int {
	a := []rune(answers)
	b := make(map[rune]bool)
	for _, q := range a {
		b[q] = true
	}
	return len(b)
}
