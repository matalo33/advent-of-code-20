package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type password struct {
	num1     int
	num2     int
	char     string
	password string
}

func (p password) isValid(policy int) bool {
	if policy == 1 {
		return strings.Count(p.password, p.char) >= p.num1 && strings.Count(p.password, p.char) <= p.num2
	} else if policy == 2 {
		return (string(p.password[p.num1-1]) == p.char && string(p.password[p.num2-1]) != p.char) || (string(p.password[p.num1-1]) != p.char && string(p.password[p.num2-1]) == p.char)
	}
	return false
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var input []string

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	passwords := parsePasswords(input)

	fmt.Printf("Legal Passwords on first policy: %v\n", countLegalPasswords(passwords, 1))

	fmt.Printf("Legal Passwords on second policy: %v\n", countLegalPasswords(passwords, 2))

}

func parsePasswords(input []string) []password {
	passwords := make([]password, len(input))
	for i, line := range input {
		s1 := strings.Split(line, " ")
		s2 := strings.Split(s1[0], "-")

		min, _ := strconv.Atoi(s2[0])
		max, _ := strconv.Atoi(s2[1])

		passwords[i] = password{
			num1:     min,
			num2:     max,
			char:     string(s1[1][0]),
			password: s1[2],
		}
	}
	return passwords
}

func countLegalPasswords(passwords []password, policy int) int {
	counter := 0
	for _, pw := range passwords {
		if pw.isValid(policy) {
			counter++
		}
	}
	return counter
}
