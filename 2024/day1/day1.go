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

func count(list []int, num int) int {
  count := 0
  for _, elem := range list {
    if elem == num {
      count++
    }
  }
  return count
}

func part2(list1, list2 []int) int {
  score := 0
  i, j := 0, 0

  for i < len(list1) && j < len(list2) {
    elem := list1[i]
    count1 := count(list1[i:], elem)
    count2 := count(list2, elem)
    score += count1 * count2 * elem
    i += count1
    j += count2
  }

  return score
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
  fmt.Println("Part 2:", part2(list1, list2))
}
