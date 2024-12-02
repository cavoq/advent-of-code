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
  Safe bool
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

    records = append(records, Record{Values: values})
  }

  return records
}

func reverse(values []int) []int {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
	return values
}

func (r *Record) isSafe(minV, maxV int) bool {
  if r.Values[0] > r.Values[1] {
    r.Values = reverse(r.Values)
  }
  for i := 0; i < len(r.Values) - 1; i++ {  
    distance := r.Values[i+1] - r.Values[i]
    if distance < minV || distance > maxV {
      return false
    }
  }
  return true
}

func main() {
  records := readRecords("input.txt")
  safeRecords := 0
  for _, record := range records {
    if record.isSafe(1, 3) {
      safeRecords++
    }
  }
  fmt.Println("Part 1:", safeRecords)
}
