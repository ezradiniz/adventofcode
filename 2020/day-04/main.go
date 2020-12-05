package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func parsePassports(input string) [][][]string {
	lines := strings.Split(input, "\n")
	passports := make([][][]string, 0)
	passport := make([][]string, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			passports = append(passports, passport)
			passport = make([][]string, 0)
		} else {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				pair := strings.Split(part, ":")
				passport = append(passport, pair)
			}
		}
	}
	return passports
}

func solve1(input string) int {
	passports := parsePassports(input)
	fields := map[string]bool{
		"byr": true,
		"iyr": true,
		"eyr": true,
		"hgt": true,
		"hcl": true,
		"ecl": true,
		"pid": true,
	}
	ret := 0
	for _, passport := range passports {
		count := 0
		for _, pair := range passport {
			key := pair[0]
			if _, ok := fields[key]; ok {
				count++
			}
		}
		if count == len(fields) {
			ret++
		}
	}
	return ret
}

func solve2(input string) int {
	passports := parsePassports(input)
	ret := 0
	for _, passport := range passports {
		count := 0
		for _, pair := range passport {
			key := pair[0]
			value := pair[1]
			switch key {
			case "byr":
				birthYear, _ := strconv.Atoi(value)
				if birthYear >= 1920 && birthYear <= 2002 {
					count++
				}
				break
			case "iyr":
				issueYear, _ := strconv.Atoi(value)
				if issueYear >= 2010 && issueYear <= 2020 {
					count++
				}
				break
			case "eyr":
				expYear, _ := strconv.Atoi(value)
				if expYear >= 2020 && expYear <= 2030 {
					count++
				}
				break
			case "hgt":
				height, _ := strconv.Atoi(value[:len(value)-2])
				unit := value[len(value)-2:]
				if unit == "cm" && height >= 150 && height <= 193 {
					count++
				} else if unit == "in" && height >= 59 && height <= 76 {
					count++
				}
				break
			case "hcl":
				if value[0] == '#' {
					valid := true
					for _, chr := range value[1:] {
						if !(chr >= '0' && chr <= '9' || chr >= 'a' && chr <= 'f') {
							valid = false
							break
						}
					}
					if valid {
						count++
					}
				}
				break
			case "ecl":
				colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
				for _, color := range colors {
					if value == color {
						count++
						break
					}
				}
				break
			case "pid":
				if len(value) == 9 {
					valid := true
					for _, d := range value {
						if d < '0' || d > '9' {
							valid = false
							break
						}
					}
					if valid {
						count++
					}
				}
				break
			}
		}
		if count == 7 {
			ret++
		}
	}
	return ret
}

func main() {
	var filename = flag.String("filename", "./input.txt", "filename")
	var part = flag.Int("part", 1, "puzzle part")
	flag.Parse()

	in, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	input := string(in)

	switch *part {
	case 1:
		fmt.Println("answer:", solve1(input))
		break
	case 2:
		fmt.Println("answer:", solve2(input))
		break
	default:
		log.Fatalf("part %v not found\n", *part)
		break
	}
}
