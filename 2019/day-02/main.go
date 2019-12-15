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
	arr := parseInput(input)
	intCode := NewIntcode(arr)

	intCode.Code[1] = 12
	intCode.Code[2] = 2

	intCode.Run()

	return intCode.Code[0]
}

func solve2(input string) int {
	arr := parseInput(input)
	original := make([]int, len(arr))
	copy(original, arr)

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			intCode := NewIntcode(arr)

			intCode.Code[1] = i
			intCode.Code[2] = j

			intCode.Run()

			if intCode.Code[0] == 19690720 {
				return 100*i + j
			}

			copy(arr, original)
		}
	}

	return -1
}

func parseInput(input string) []int {
	input = strings.Replace(input, "\n", "", -1)
	arrString := strings.Split(input, ",")
	arr := make([]int, len(arrString))
	for idx, v := range arrString {
		number, _ := strconv.Atoi(v)
		arr[idx] = number
	}
	return arr
}

const addOperation = 1
const multOperation = 2
const haltOperation = 99

type Intcode struct {
	Code      []int
	idxOpcode int
}

func NewIntcode(arr []int) *Intcode {
	return &Intcode{
		Code:      arr,
		idxOpcode: -4,
	}
}

func (i *Intcode) Next() bool {
	max := len(i.Code)
	for idx := i.idxOpcode + 4; idx < max; idx += 4 {
		opCode := i.Code[idx]
		if opCode == haltOperation {
			break
		}
		if opCode == addOperation || opCode == multOperation {
			i.idxOpcode = idx
			return true
		}
	}
	i.idxOpcode = max
	return false
}

func (i *Intcode) Current() int {
	if i.idxOpcode >= 0 && i.idxOpcode < len(i.Code) {
		return i.Code[i.idxOpcode]
	}
	return i.idxOpcode
}

func (i *Intcode) Add(at, n1, n2 int) {
	i.Code[at] = i.Code[n1] + i.Code[n2]
}

func (i *Intcode) Mult(at, n1, n2 int) {
	i.Code[at] = i.Code[n1] * i.Code[n2]
}

func (i *Intcode) Instructions() (a, b, at int) {
	idx := i.idxOpcode
	return i.Code[idx+1], i.Code[idx+2], i.Code[idx+3]
}

func (i *Intcode) Run() {
	for i.Next() {
		opcode := i.Current()
		switch opcode {
		case addOperation:
			op1, op2, at := i.Instructions()
			i.Add(at, op1, op2)
			break
		case multOperation:
			op1, op2, at := i.Instructions()
			i.Mult(at, op1, op2)
			break
		}
	}
}
