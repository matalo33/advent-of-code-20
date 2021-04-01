package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type instruction struct {
	direction byte
	quantity  int
}

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)
	instructions := parseData(buf)

	fmt.Printf("Manhattan Distance: %v\n", manhattanDistance(navigate(instructions, 90)))
	fmt.Printf("Manhattan Distance with waypoint: %v\n", manhattanDistance(
		navigateWithWaypoint(instructions, 1, 10)))
}

func parseData(data *bytes.Buffer) (instructionList []instruction) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		distance, _ := strconv.Atoi(line[1:])
		instructionList = append(instructionList, instruction{
			direction: byte(line[0]),
			quantity:  distance,
		})
	}
	return
}

func navigate(instructions []instruction, direction int) (northing, easting int) {
	for _, instruction := range instructions {
		switch op := instruction.direction; op {
		case 'N':
			northing += instruction.quantity
		case 'S':
			northing -= instruction.quantity
		case 'E':
			easting += instruction.quantity
		case 'W':
			easting -= instruction.quantity
		case 'R':
			direction = modulo((direction + instruction.quantity), 360)
		case 'L':
			direction = modulo((direction - instruction.quantity), 360)
		case 'F':
			switch dir := direction; dir {
			case 0:
				northing += instruction.quantity
			case 90:
				easting += instruction.quantity
			case 180:
				northing -= instruction.quantity
			case 270:
				easting -= instruction.quantity
			}
		}
	}
	return northing, easting
}

func navigateWithWaypoint(instructions []instruction, offsetY, offsetX int) (shipNorthing, shipEasting int) {
	for _, instruction := range instructions {
		switch op := instruction.direction; op {
		case 'N':
			offsetY += instruction.quantity
		case 'S':
			offsetY -= instruction.quantity
		case 'E':
			offsetX += instruction.quantity
		case 'W':
			offsetX -= instruction.quantity
		case 'R':
			offsetX, offsetY = rotateRight(offsetX, offsetY, instruction.quantity)
		case 'L':
			offsetX, offsetY = rotateLeft(offsetX, offsetY, instruction.quantity)
		case 'F':
			shipEasting += offsetX * instruction.quantity
			shipNorthing += offsetY * instruction.quantity
		}
	}
	return shipNorthing, shipEasting
}

func manhattanDistance(northing, easting int) int {
	return abs(northing) + abs(easting)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func rotateRight(x, y, d int) (int, int) {
	if d/90 > 1 {
		x, y = rotateRight(x, y, d-90)
	}
	return y, -x
}

func rotateLeft(x, y, d int) (int, int) {
	if d/90 > 1 {
		x, y = rotateLeft(x, y, d-90)
	}
	return -y, x
}

func modulo(x, m int) int {
	return (x%m + m) % m
}
