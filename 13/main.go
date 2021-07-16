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

	fmt.Printf("Earliest bus ID: %v", part1(findEarliestBusId(timetable, earliestTime), earliestTime))
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

func part1(bus, earliestTime int) int {
	nextDep := earliestTime + bus - 1
	nextDep -= nextDep % bus
	return (nextDep - earliestTime) * bus
}

func findEarliestBusId(busses []int, earliestTime int) (id int) {
	earliestBus, earliestDep := math.MaxInt32, math.MaxInt32

	for _, bus := range busses {
		nextDep := earliestTime + bus - 1
		nextDep -= nextDep % bus
		if nextDep < earliestDep {
			earliestDep = nextDep
			earliestBus = bus
		}
	}

	return earliestBus
}
