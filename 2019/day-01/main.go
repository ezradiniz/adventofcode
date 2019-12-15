package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var filepath = flag.String("filepath", "./input.txt", "Filepath")
var part = flag.Int("part", 1, "Puzzle part")

func main() {
	flag.Parse()

	in, err := ioutil.ReadFile(*filepath)
	if err != nil {
		panic(err)
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
		fmt.Println("Error: invalid part")
		break
	}
}

func solve1(input string) int {
	modules := parseInput(input)
	res := 0
	for _, module := range modules {
		mass, _ := strconv.Atoi(module)
		res += getFuel(mass)
	}
	return res
}

func solve2(input string) int {
	modules := parseInput(input)
	res := 0
	for _, module := range modules {
		mass, _ := strconv.Atoi(module)
		for {
			fuel := getFuel(mass)
			if fuel <= 0 {
				break
			}
			res += fuel
			mass = fuel
		}
	}
	return res
}

func parseInput(input string) []string {
	in := strings.Split(input, "\n")
	return in[:len(in)-1]
}

func getFuel(mass int) int {
	return mass/3 - 2
}
