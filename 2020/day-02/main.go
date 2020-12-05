package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
		output := solve1(input)
		fmt.Println("answer:", output)
		break
	case 2:
		output := solve2(input)
		fmt.Println("answer:", output)
		break
	default:
		log.Fatalf("part %v not found\n", *part)
		break
	}
}

func solve1(input string) int {
	lines := parseInput(input)
	ret := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		policy := strings.Split(parts[0], "-")
		lowest, _ := strconv.Atoi(policy[0])
		highest, _ := strconv.Atoi(policy[1])
		letter := parts[1][0]
		password := parts[2]
		count := 0
		for _, c := range password {
			if c == rune(letter) {
				count++
			}
		}
		if count >= lowest && count <= highest {
			ret++
		}
	}
	return ret
}

func solve2(input string) int {
	lines := parseInput(input)
	ret := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		policy := strings.Split(parts[0], "-")
		l, _ := strconv.Atoi(policy[0])
		h, _ := strconv.Atoi(policy[1])
		l--
		h--
		letter := parts[1][0]
		password := parts[2]
		if l >= 0 && h < len(password) && ((password[l] == letter && password[h] != letter) || (password[l] != letter && password[h] == letter)) {
			ret++
		}
	}
	return ret
}

func parseInput(input string) []string {
	lines := strings.Split(input, "\n")
	return lines[:len(lines)-1]
}
