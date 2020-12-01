package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var expenses []int

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		expenses = append(expenses, i)
	}

	fmt.Printf("Fixed expense report: %v\n", FixExpenseReport(expenses, false))

	fmt.Printf("Fixed expense report: %v\n", FixExpenseReport(expenses, true))
}

// FixExpenseReport foo
func FixExpenseReport(expenses []int, part2 bool) int {
	for _, i := range expenses {
		for _, j := range expenses {
			if part2 {
				for _, k := range expenses {
					if i+j+k == 2020 {
						return i * j * k
					}
				}
			} else {
				if i+j == 2020 {
					return i * j
				}
			}
		}
	}
	return 0
}
