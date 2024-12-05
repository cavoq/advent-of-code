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

func topologicalSort(rules []Rule) ([]int, error) {
	graph := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, rule := range rules {
		a, b := rule.first, rule.second
		graph[a] = append(graph[a], b)
		inDegree[b]++
		if _, exists := inDegree[a]; !exists {
			inDegree[a] = 0
		}
	}

	var queue []int
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var result []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(result) != len(inDegree) {
		return nil, fmt.Errorf("cycle detected in the input rules")
	}

	return result, nil
}

func getRelevantRules(prcd []int, rules *[]Rule) []Rule {
	var relevantRules []Rule
	for _, rule := range *rules {
		firstIdx := -1
		secondIdx := -1
		for i, num := range prcd {
			if num == rule.first {
				firstIdx = i
			}
			if num == rule.second {
				secondIdx = i
			}
		}
		if firstIdx != -1 && secondIdx != -1 {
			relevantRules = append(relevantRules, rule)
		}
	}
	return relevantRules
}

func solve(prcds [][]int, rules *[]Rule) (int, int) {
	sum := 0
	var prcdsToFix [][]int
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
		} else {
			prcdsToFix = append(prcdsToFix, prcd)
		}
	}

	sum2 := 0
	for _, prcd := range prcdsToFix {
		relevantRules := getRelevantRules(prcd, rules)
		sorted, err := topologicalSort(relevantRules)
		if err != nil {
			panic(err)
		}
		middle := len(sorted) / 2
		sum2 += sorted[middle]
	}

	return sum, sum2
}

func main() {
	rules := readRules("input.txt")
	procedures := readProcodures("input.txt")
	sum, sum2 := solve(procedures, &rules)
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
