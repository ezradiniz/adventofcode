package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func solve1(input string) int64 {
	g := parseInput(input)
	return traverse(g, 3, 1)
}

func solve2(input string) int64 {
	g := parseInput(input)
	return traverse(g, 1, 1) * traverse(g, 3, 1) * traverse(g, 5, 1) * traverse(g, 7, 1) * traverse(g, 1, 2)
}

func traverse(grid [][]string, right int, down int) int64 {
	var ret int64
	j := 0
	for i := 0; i < len(grid); i += down {
		if grid[i][j] == "#" {
			ret++
		}
		j = (j + right) % len(grid[i])

	}
	return ret
}

func parseInput(input string) [][]string {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	grid := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]string, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			grid[i][j] = string(lines[i][j])
		}
	}
	return grid
}
