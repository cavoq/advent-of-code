package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	Values []int
	Safe   bool
}

func readRecords(path string) []Record {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records := []Record{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		values := []int{}

		for _, part := range parts {
			value, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			values = append(values, value)
		}

		records = append(records, Record{Values: values, Safe: true})
	}

	return records
}

func diff(num1, num2 int) (int, int) {
	if num1 > num2 {
		return num1 - num2, -1
	}
	return num2 - num1, 1
}

func isSafe(minV, maxV int, slice []int) bool {
	direction := 1
	if slice[0] > slice[1] {
		direction = -1
	}
	for i := 0; i < len(slice)-1; i++ {
		distance, newDir := diff(slice[i], slice[i+1])
		if distance < minV || distance > maxV || newDir != direction {
			return false
		}
	}
	return true
}

func isSafe2(minV, maxV int, values []int) bool {
	if isSafe(minV, maxV, values) {
		return true
	}

	for i := 0; i < len(values); i++ {
		slice := make([]int, len(values))
		copy(slice, values)
		slice = append(slice[:i], slice[i+1:]...)
		if isSafe(minV, maxV, slice) {
			return true
		}
	}
	return false
}

func main() {
	records := readRecords("input.txt")

	safeRecords := 0
	for i := 0; i < len(records); i++ {
		if isSafe(1, 3, records[i].Values) {
			safeRecords++
		}
	}
	fmt.Println("Part 1:", safeRecords)

	safeRecords = 0
	for i := 0; i < len(records); i++ {
		if isSafe2(1, 3, records[i].Values) {
			safeRecords++
		}
	}
	fmt.Println("Part 2:", safeRecords)
}
