package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readInput(path string) ([]int, []int, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, nil, fmt.Errorf("Error opening file: %v", err)
  }
  defer file.Close()

  var list1, list2 []int
  scanner := bufio.NewScanner(file)

  for scanner.Scan() { 
    line := scanner.Text() 
    numbers := strings.Fields(line)

    if len(numbers) != 2 {
      return nil, nil, fmt.Errorf("Invalid line format: %v", line)
    }
    
    num1, err1 := strconv.Atoi(numbers[0])
    num2, err2 := strconv.Atoi(numbers[1])

    if err1 != nil || err2 != nil {
      return nil, nil, fmt.Errorf("Invalid number format: %v", numbers)
    }

    list1 = append(list1, num1)
    list2 = append(list2, num2)
  }

  if err := scanner.Err(); err != nil {
    return nil, nil, fmt.Errorf("Error scanning file: %v", err)
  }

  return list1, list2, nil
}

func diff(num1, num2 int) int {
  if num1 > num2 {
    return num1 - num2
  }
  return num2 - num1
}

func part1(list1, list2 []int) int {
  if len(list1) != len(list2) {
    return -1
  }
  
  sum := 0
  for i := 0; i < len(list1); i++ {
    sum += diff(list1[i], list2[i])
  }

  return sum
}

func main() {
  path := "input.txt"
  list1, list2, err := readInput(path)
  if err != nil {
    fmt.Println(err)
    return
  }

  sort.Ints(list1)
  sort.Ints(list2)

  fmt.Println("Part 1:", part1(list1, list2))
}
