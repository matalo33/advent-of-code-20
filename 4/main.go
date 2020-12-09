package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	buf := bytes.NewBuffer(data)

	part1, part2 := checkPassports(buf)

	fmt.Printf("Valid passports %v\n", part1)
	fmt.Printf("REALLY valid passports %v\n", part2)
}

func checkPassports(passportData *bytes.Buffer) (int, int) {
	scanner := bufio.NewScanner(passportData)

	validPassports, reallyValidPassports := 0, 0
	passport := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if validatePassport(passport) {
				validPassports++
			}
			if reallyValidatePassport(passport) {
				reallyValidPassports++
			}
			passport = make(map[string]string)
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				f := strings.Split(field, ":")
				passport[f[0]] = f[1]
			}
		}
	}

	// 💩 scanner doesnt give the final empty line so call validate once
	// manually to check the final passport
	if validatePassport(passport) {
		validPassports++
	}
	if reallyValidatePassport(passport) {
		reallyValidPassports++
	}

	return validPassports, reallyValidPassports
}

func validatePassport(passport map[string]string) bool {
	_, byr := passport["byr"]
	_, iyr := passport["iyr"]
	_, eyr := passport["eyr"]
	_, hgt := passport["hgt"]
	_, hcl := passport["hcl"]
	_, ecl := passport["ecl"]
	_, pid := passport["pid"]

	return byr && iyr && eyr && hgt && hcl && ecl && pid
}

func reallyValidatePassport(passport map[string]string) bool {
	byr, _ := passport["byr"]
	iyr, _ := passport["iyr"]
	eyr, _ := passport["eyr"]
	hgt, _ := passport["hgt"]
	hcl, _ := passport["hcl"]
	ecl, _ := passport["ecl"]
	pid, _ := passport["pid"]

	return checkRange(byr, 1920, 2002) &&
		checkRange(iyr, 2010, 2020) &&
		checkRange(eyr, 2020, 2030) &&
		checkHgt(hgt) &&
		checkHcl(hcl) &&
		checkEcl(ecl) &&
		checkPid(pid)
}

func checkRange(input string, min, max int) bool {
	i, _ := strconv.Atoi(input)
	if i >= min && i <= max {
		return true
	}
	return false
}

func checkHgt(input string) bool {
	if strings.HasSuffix(input, "cm") {
		hgt := strings.Split(input, "cm")
		h, _ := strconv.Atoi(hgt[0])
		return h >= 150 && h <= 193
	}
	if strings.HasSuffix(input, "in") {
		hgt := strings.Split(input, "in")
		h, _ := strconv.Atoi(hgt[0])
		return h >= 59 && h <= 76
	}
	return false
}

func checkHcl(input string) bool {
	regex, _ := regexp.MatchString("^#[a-f0-9]{6}$", input)
	return regex
}

func checkEcl(input string) bool {
	for _, c := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if input == c {
			return true
		}
	}
	return false
}

func checkPid(input string) bool {
	regex, _ := regexp.MatchString("^[0-9]{9}$", input)
	return regex
}
