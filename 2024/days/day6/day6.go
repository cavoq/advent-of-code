package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	x, y int
}

func readInput(path string) ([][]rune, [2]int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var mtx [][]rune
	var guard [2]int

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		mtx = append(mtx, []rune(line))
		for x, c := range line {
			if c == '^' {
				guard = [2]int{x, y}
			}
		}
	}

	return mtx, guard
}

func switchDirection(dir *Direction) {
	if dir.x == 0 && dir.y == 1 {
		dir.x, dir.y = 1, 0
	} else if dir.x == 1 && dir.y == 0 {
		dir.x, dir.y = 0, -1
	} else if dir.x == 0 && dir.y == -1 {
		dir.x, dir.y = -1, 0
	} else if dir.x == -1 && dir.y == 0 {
		dir.x, dir.y = 0, 1
	}
}

func solve(mtx [][]rune, guard [2]int) int {
	visits := 1
	curX, curY := guard[0], guard[1]
	dir := Direction{0, 1}
	for {
		if string(mtx[curY])[curX] == '.' {
			visits++
		}
		mtx[curY][curX] = 'X'
		nextX, nextY := curX+dir.x, curY-dir.y
		if nextY > len(mtx[0])-1 || nextX > len(mtx)-1 || nextX < 0 || nextY < 0 {
			break
		}
		if string(mtx[nextY])[nextX] == '#' {
			switchDirection(&dir)
			curX, curY = curX+dir.x, curY-dir.y
		} else {
			curX, curY = nextX, nextY
		}
	}
	return visits
}

func main() {
	mtx, guard := readInput("input.txt")
	visits := solve(mtx, guard)
	fmt.Println("Part 1:", visits)
}
