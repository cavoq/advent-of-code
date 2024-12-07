package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation func(a, b int) int

var ADD Operation = func(a, b int) int { return a + b }
var MULT Operation = func(a, b int) int { return a * b }
var CONCAT Operation = func(a, b int) int {
	concatStr := fmt.Sprintf("%d%d", a, b)
	concatNum, err := strconv.Atoi(concatStr)
	if err != nil {
		fmt.Println("Error converting concatenated string to integer:", err)
		os.Exit(1)
	}
	return concatNum
}

type Equation struct {
	testValue int
	numbers   []int
	ops       []Operation
}

func readInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, strings.Fields(line))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return lines
}

func NewEquation(line []string, ops []Operation) Equation {
	testValue, _ := strconv.Atoi(strings.TrimSuffix(line[0], ":"))
	numbers := make([]int, len(line)-1)
	for i, numStr := range line[1:] {
		val, _ := strconv.Atoi(numStr)
		numbers[i] = val
	}

	return Equation{
		testValue: testValue,
		numbers:   numbers,
		ops:       ops,
	}
}

func (eq *Equation) generateOperationSequences(length int) [][]Operation {
	if length == 0 {
		return [][]Operation{{}}
	}

	subSequences := eq.generateOperationSequences(length - 1)
	var sequences [][]Operation

	for _, seq := range subSequences {
		for _, op := range eq.ops {
			newSeq := append([]Operation{}, seq...)
			newSeq = append(newSeq, op)
			sequences = append(sequences, newSeq)
		}
	}

	return sequences
}

func (eq *Equation) seekEquation() *int {
	length := len(eq.numbers) - 1
	opSequences := eq.generateOperationSequences(length)

	for _, ops := range opSequences {
		total := eq.numbers[0]
		for i, op := range ops {
			total = op(total, eq.numbers[i+1])
		}
		if total == eq.testValue {
			return &eq.testValue
		}
	}

	return nil
}

func part1(input [][]string) int {
	total := 0
	for _, line := range input {
		eq := NewEquation(line, []Operation{ADD, MULT})
		if result := eq.seekEquation(); result != nil {
			total += *result
		}
	}
	return total
}

func part2(input [][]string) int {
	total := 0
	for _, line := range input {
		eq := NewEquation(line, []Operation{ADD, MULT, CONCAT})
		if result := eq.seekEquation(); result != nil {
			total += *result
		}
	}
	return total
}

func main() {
	input := readInput("input.txt")
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
