package main

import (
	"bytes"
	"testing"
)

func TestFoo(t *testing.T) {
	tests := []struct {
		input string
		part1 int
		part2 int
	}{
		{
			input: `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`,
			part1: 2,
			part2: 2,
		},
		{
			input: `hcl:5d90f0 cid:270 ecl:#66dc9c hgt:62cm byr:1945 pid:63201172 eyr:2026

ecl:amb byr:1943 iyr:2014 eyr:2028
pid:333051831

byr:1971
eyr:2021 iyr:2015 pid:158388040 hcl:#18171d ecl:brn hgt:179cm

byr:1936
pid:707057570 iyr:2014 ecl:amb cid:299 eyr:2030
hcl:#c0946f hgt:186cm

hgt:163cm iyr:2013 ecl:gry hcl:#86e981 byr:1939
eyr:2020 pid:241741372 cid:203

ecl:brn hcl:#341e13
pid:686617364 byr:1929 eyr:2029 hgt:160cm cid:280 iyr:2020

byr:2002 hcl:#623a2f
pid:253005469 iyr:2011 ecl:hzl hgt:184cm eyr:2027

ecl:#bb984b eyr:2040
hgt:188in
iyr:2005 hcl:c5be8e pid:174cm cid:161 byr:2004

ecl:oth iyr:2010 cid:128 hgt:153cm byr:1991
pid:24061445 eyr:2025 hcl:#54d43e

hcl:z
iyr:2023 pid:981178503 ecl:gmt eyr:2038 byr:2004

ecl:gry eyr:2022 iyr:1981 pid:566993828
byr:1941 hcl:#341e13 hgt:176cm
`,
			part1: 8,
			part2: 5,
		},
		{
			input: `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719
`,
			part1: 8,
			part2: 4,
		},
	}

	for tn, test := range tests {
		buf := bytes.NewBufferString(test.input)
		r1, r2 := checkPassports(buf)

		if r1 != test.part1 {
			t.Errorf("Test %v: Got: %v, want: %v", tn, r1, test.part1)
		}

		if r2 != test.part2 {
			t.Errorf("Test %v: Got: %v, want: %v", tn, r2, test.part2)
		}

	}
}
