package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

func parseInput(input string) [][]int {
	regex := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return regex.FindAllStringSubmatchIndex(input, -1)
}

func parseSkips(input string) [][]int {
	regex := regexp.MustCompile(`don't\(\)|do\(\)`)
	matches := regex.FindAllStringSubmatchIndex(input, -1)

	var skips [][]int
	var skipStart, skipEnd int
	skipping := false

	for i, m := range matches {
		startIndex, endIndex := m[0], m[1]
		match := input[startIndex:endIndex]
		if match == "don't()" && !skipping {
			skipping = true
			skipStart = startIndex
		}
		if match == "do()" && skipping {
			skipping = false
			skipEnd = endIndex
			skips = append(skips, []int{skipStart, skipEnd})
		}
		if i == len(matches)-1 && skipping {
			skips = append(skips, []int{skipStart, len(input)})
		}
	}
	return skips
}

func solve(expr [][]int, input string) int {
	sum := 0
	for _, e := range expr {
		a, err := strconv.Atoi(input[e[2]:e[3]])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return 0
		}
		b, err := strconv.Atoi(input[e[4]:e[5]])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return 0
		}
		sum += a * b
	}
	return sum
}

func solve2(expr [][]int, skips [][]int, input string) int {
	sum := 0
outerloop:
	for _, e := range expr {
		for _, s := range skips {
			if e[0] >= s[0] && e[1] <= s[1] {
				continue outerloop
			}
		}
		a, err := strconv.Atoi(input[e[2]:e[3]])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return 0
		}
		b, err := strconv.Atoi(input[e[4]:e[5]])
		if err != nil {
			fmt.Println("Error converting string to int", err)
			return 0
		}
		sum += a * b
	}
	return sum
}

func main() {
	input := readInput()
	expr := parseInput(input)
	result := solve(expr, input)
	fmt.Println("Part 1:", result)
	skips := parseSkips(input)
	result2 := solve2(expr, skips, input)
	fmt.Println("Part 2:", result2)
}
