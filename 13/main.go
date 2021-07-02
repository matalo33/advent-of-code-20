package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')
	earliestTime, err := strconv.Atoi(strings.TrimSuffix(line, "\n"))
	if err != nil {
		panic(err)
	}

	line, _ = reader.ReadString('\n')
	timetable := parseTimetable(line)

	fmt.Printf("Earliest bus ID: %v", findEarliestBusId(timetable, earliestTime))
}

func parseTimetable(line string) (busses []int) {
	for _, bus := range strings.Split(line, ",") {
		busNum, err := strconv.Atoi(bus)
		if err == nil {
			busses = append(busses, busNum)
		}
	}
	return
}

func findEarliestBus(busses []int, earliestTime int) (id int) {
	earliestBus := math.MaxInt32

	for _, bus := range busses {
		if earliestTime%bus == 0 {
			id = (earliestTime / bus)
		} else {
			id = (earliestTime / bus) + 1
		}

		if id < earliestBus {
			earliestBus = id
		}
	}
}
