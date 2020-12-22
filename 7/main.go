package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

/* consts */
var rxpOuterBag = regexp.MustCompile(`^([a-z].*?) bags`)
var rxpInnerBags = regexp.MustCompile(`([0-9]) ([a-z].*?) bag`)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	p1, p2 := bagIt(buf)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func bagIt(data *bytes.Buffer) (int, int) {
	bags := make(map[string]map[string]int)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		bagRule := scanner.Text()
		parseBagRule(bags, bagRule)
	}

	return numBagsThatFit(bags, "shiny gold"), howManyBagsInBag(bags, "shiny gold") - 1
}

func parseBagRule(bags map[string]map[string]int, bagRule string) {
	outer := rxpOuterBag.FindStringSubmatch(bagRule)[1]
	bags[outer] = make(map[string]int)

	for _, b := range rxpInnerBags.FindAllStringSubmatch(bagRule, -1) {
		bagCount, _ := strconv.Atoi(b[1])
		bags[outer][b[2]] = bagCount
	}
}

func numBagsThatFit(bagList map[string]map[string]int, myBag string) (bagCount int) {
	for bigBag := range bagList {
		if doesBagFitInBag(bagList, bigBag, myBag) {
			if bigBag == myBag {
				continue
			}
			bagCount++
		}
	}
	return bagCount
}

func doesBagFitInBag(bagList map[string]map[string]int, currentBag, targetBag string) bool {
	// Bag directly fits in the current bag
	if _, ok := bagList[currentBag][targetBag]; ok {
		return true
	} else {
		for bag := range bagList[currentBag] {
			if doesBagFitInBag(bagList, bag, targetBag) {
				return true
			}
		}
	}
	return false
}

func howManyBagsInBag(bagList map[string]map[string]int, currentBag string) (count int) {
	for bag, capacity := range bagList[currentBag] {
		count += capacity * howManyBagsInBag(bagList, bag)
	}
	return count + 1
}
