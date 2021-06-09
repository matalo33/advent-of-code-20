package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	earliestTime, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	timetable := parseTimetable(line)
}

func parseTimetable(string line)
