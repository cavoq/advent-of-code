package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	first  int
	second int
}

func readRules(path string) []Rule {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rules []Rule
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, "|")
		first, err := strconv.Atoi(parts[0])
		second, err2 := strconv.Atoi(parts[1])
		if err != nil || err2 != nil {
			panic(err)
		}
		rules = append(rules, Rule{first, second})
	}
	return rules
}

func readProcodures(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var procedures [][]int
	for i := len(lines) - 1; len(lines[i]) != 0; i-- {
		parts := strings.Split(lines[i], ",")
		var procedure []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			procedure = append(procedure, num)
		}
		procedures = append(procedures, procedure)
	}
	return procedures
}

func isValid(prcd []int, rule Rule) bool {
	firstIdx := -1
	secondIdx := -1
	for i, num := range prcd {
		if num == rule.first {
			firstIdx = i
		}
		if num == rule.second {
			secondIdx = i
		}
		if firstIdx == -1 || secondIdx == -1 {
			continue
		}
		if firstIdx > secondIdx {
			return false
		}
	}
	return true
}

func solve(prcds [][]int, rules *[]Rule) int {
	sum := 0
	for _, prcd := range prcds {
		valid := true
		for _, rule := range *rules {
			if !isValid(prcd, rule) {
				valid = false
				break
			}
		}
		if valid {
			middle := len(prcd) / 2
			sum += prcd[middle]
		}
	}
	return sum
}

func main() {
	rules := readRules("input.txt")
	procedures := readProcodures("input.txt")
	fmt.Println("Part 1:", solve(procedures, &rules))
}
