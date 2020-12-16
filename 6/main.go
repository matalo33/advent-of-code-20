package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	p1, p2 := sumAnswers(buf)

	fmt.Printf("Sum of all answers: %v\n", p1)
	fmt.Printf("Sum of common answers: %v\n", p2)
}

func sumAnswers(data *bytes.Buffer) (int, int) {
	scanner := bufio.NewScanner(data)
	sumAllAnswers, sumCommonAnswers := 0, 0
	answersSet := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sumAllAnswers += countAnswers(answersSet)
			sumCommonAnswers += countCommonAnswers(answersSet)
			answersSet = []string{}
		} else {
			answersSet = append(answersSet, line)
		}
	}
	return (sumAllAnswers + countAnswers(answersSet)), (sumCommonAnswers + countCommonAnswers(answersSet))
}

func countAnswers(answersSet []string) int {
	answers := strings.Join(answersSet, "")
	a := []rune(answers)
	b := make(map[rune]bool)
	for _, q := range a {
		b[q] = true
	}
	return len(b)
}

func countCommonAnswers(answersSet []string) int {
	if len(answersSet) == 1 {
		return len(answersSet[0])
	}
	commonSet := []rune(answersSet[0])
	for i := 0; i < len(answersSet); i++ {
		commonSet = intersection(commonSet, []rune(answersSet[i]))
	}
	return len(commonSet)
}

func intersection(a, b []rune) (out []rune) {
	m := make(map[rune]bool)

	for _, q := range a {
		m[q] = true
	}

	for _, q := range b {
		if _, t := m[q]; t {
			out = append(out, q)
		}
	}
	return
}
