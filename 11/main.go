package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type adj struct {
	x int
	y int
}

var adjSeats = map[int]adj{
	0: {x: -1, y: -1},
	1: {x: 0, y: -1},
	2: {x: 1, y: -1},
	3: {x: -1, y: 0},
	4: {x: 1, y: 0},
	5: {x: -1, y: 1},
	6: {x: 0, y: 1},
	7: {x: 1, y: 1},
}

var debug = false

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)
	seatMap := parseData(buf)

	seatMapNotFussy := make([][]rune, len(seatMap))
	for x := 0; x < len(seatMap); x++ {
		seatMapNotFussy[x] = make([]rune, len(seatMap[x]))
		copy(seatMapNotFussy[x], seatMap[x])
	}

	seatMap = takeSeats(seatMap, 4, true)
	seatMapNotFussy = takeSeats(seatMapNotFussy, 5, false)

	fmt.Printf("Occupied seats: %v\n", countOccupiedSeats(seatMap))
	fmt.Printf("Occupied seats for fussy people: %v\n", countOccupiedSeats(seatMapNotFussy))
}

func parseData(data *bytes.Buffer) (seatMap [][]rune) {
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		seatRow := scanner.Text()
		seatMapRow := make([]rune, len(seatRow))
		for i, c := range seatRow {
			if c == 'L' {
				// Empty seats become occupied on first pass anyway
				seatMapRow[i] = '#'
			} else {
				seatMapRow[i] = '.'
			}
		}
		seatMap = append(seatMap, seatMapRow)
	}
	return
}

// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// If a seat is occupied (#) and four/five or more seats adjacent to it are also occupied, the seat becomes empty.
func takeSeats(seatMap [][]rune, adjSeatLimit int, fussy bool) [][]rune {
	for {
		prevSeatMap := make([][]rune, len(seatMap))
		for x := 0; x < len(seatMap); x++ {
			prevSeatMap[x] = make([]rune, len(seatMap[x]))
			copy(prevSeatMap[x], seatMap[x])
		}

		// X = DOWN // Y = ALONG //
		for x := 0; x < len(seatMap); x++ {
			for y := 0; y < len(seatMap[x]); y++ {

				if prevSeatMap[x][y] == '.' {
					continue
				}

				thisSeatOccupied := false
				if prevSeatMap[x][y] == '#' {
					thisSeatOccupied = true
				}

				adjSeatCounter := 0

			seatBreach:
				for _, adjSeat := range adjSeats {
					visionDistance := 1
					for {
						nextX := x + (adjSeat.x * visionDistance)
						nextY := y + (adjSeat.y * visionDistance)

						if isOutOfRange(
							nextX,
							nextY,
							len(prevSeatMap),
							len(prevSeatMap[x])) {
							continue seatBreach
						}

						if fussy {
							if prevSeatMap[nextX][nextY] == '#' {
								adjSeatCounter++
							}
							if adjSeatCounter == adjSeatLimit {
								break seatBreach
							}
							continue seatBreach
						} else {
							if prevSeatMap[nextX][nextY] == 'L' {
								continue seatBreach
							}
							if prevSeatMap[nextX][nextY] == '#' {
								adjSeatCounter++
								if adjSeatCounter >= adjSeatLimit {
									break seatBreach
								}
								continue seatBreach
							}
							visionDistance++
						}
					}
				}

				// Sit or stand
				if thisSeatOccupied && (adjSeatCounter >= adjSeatLimit) {
					seatMap[x][y] = 'L'
				} else if !thisSeatOccupied && adjSeatCounter == 0 {
					seatMap[x][y] = '#'
				}
			}
		}

		if reflect.DeepEqual(prevSeatMap, seatMap) {
			printSeatMap(seatMap)
			return seatMap
		}
	}
}

func printSeatMap(seatMap [][]rune) {
	if debug {
		for x := 0; x < len(seatMap); x++ {
			for y := 0; y < len(seatMap[x]); y++ {
				fmt.Printf("%s", string(seatMap[x][y]))
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
}

func isOutOfRange(x, y, lx, ly int) bool {
	if x < 0 || y < 0 {
		return true
	}
	if x >= lx || y >= ly {
		return true
	}
	return false
}

func countOccupiedSeats(seatMap [][]rune) int {
	seats := 0
	for x := 0; x < len(seatMap); x++ {
		for y := 0; y < len(seatMap[x]); y++ {
			if seatMap[x][y] == '#' {
				seats++
			}
		}
	}
	return seats
}
