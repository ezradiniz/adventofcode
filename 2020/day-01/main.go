package main

import (
	"errors"
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
		output, err := solve1(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("answer:", output)
		break
	case 2:
		output, err := solve2(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("answer:", output)
		break
	default:
		log.Fatalf("part %v not found\n", *part)
		break
	}
}

func solve1(input string) (int, error) {
	entries := parseInput(input)
	seen := make(map[int]struct{})
	for _, entry := range entries {
		seen[entry] = struct{}{}
	}
	for _, entry := range entries {
		comp := 2020 - entry
		if _, ok := seen[comp]; ok {
			return entry * comp, nil
		}
	}
	return 0, errors.New("Entries not found")
}

func solve2(input string) (int, error) {
	entries := parseInput(input)
	seen := make(map[int]struct{})
	for i := 0; i < len(entries); i++ {
		cur := 2020 - entries[i]
		for j := i + 1; j < len(entries); j++ {
			comp := cur - entries[j]
			if _, ok := seen[comp]; ok {
				return entries[i] * entries[j] * comp, nil
			}
			seen[entries[j]] = struct{}{}
		}
	}
	return 0, errors.New("Entries not found")
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	entries := make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		entry, _ := strconv.Atoi(line)
		entries[i] = entry
	}
	return entries
}
