package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	x, y int
}

type Position struct {
	x, y int
	dir  Direction
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

func getPath(mtx *[][]rune, guard *[2]int, withDir bool) []Position {
	x, y := (*guard)[0], (*guard)[1]
	dir := Direction{0, 1}
	visited := []Position{{x, y, dir}}
	for {
		nextX, nextY := x+dir.x, y-dir.y
		if withDir && (*mtx)[y][x] == '.' {
			visited = append(visited, Position{x, y, dir})
			(*mtx)[y][x] = 'X'
		} else if !withDir {
			visited = append(visited, Position{x, y, dir})
		}
		if nextY >= len((*mtx)[0]) || nextX >= len(*mtx) || nextX < 0 || nextY < 0 {
			break
		}
		if string((*mtx)[nextY])[nextX] == '#' {
			switchDirection(&dir)
		} else {
			x, y = nextX, nextY
		}
	}
	return visited
}

func loop(mtx [][]rune, start Position) bool {
	visited := make(map[Position]bool)
	pos := start
	for steps := 0; steps <= len(mtx)*len(mtx[0]); steps++ {
		if visited[pos] {
			return true
		}
		visited[pos] = true
		nextX, nextY := pos.x+pos.dir.x, pos.y-pos.dir.y
		if nextX < 0 || nextY < 0 || nextY >= len(mtx) || nextX >= len(mtx[0]) {
			return false
		}
		if mtx[nextY][nextX] == '#' {
			switchDirection(&pos.dir)
		} else {
			pos.x, pos.y = nextX, nextY
		}
	}
	return false
}

func obstacleCandidates(mtx [][]rune, positions []Position) int {
	visitedCandidates := make(map[string]bool)
	candidates := []Position{}
	for _, pos := range positions {
		nextX, nextY := pos.x+pos.dir.x, pos.y-pos.dir.y
		candidateKey := fmt.Sprintf("%d,%d", nextX, nextY)
		if visitedCandidates[candidateKey] {
			continue
		}
		visitedCandidates[candidateKey] = true
		if nextX < 0 || nextY < 0 || nextY >= len(mtx) || nextX >= len(mtx[0]) {
			continue
		}
		tempMtx := make([][]rune, len(mtx))
		for i := range mtx {
			tempMtx[i] = append([]rune{}, mtx[i]...)
		}
		tempMtx[nextY][nextX] = '#'
		if loop(tempMtx, pos) {
			candidates = append(candidates, Position{nextX, nextY, pos.dir})
		}
	}
	return len(candidates)
}

func validObstacle(mtx [][]rune, x, y int) bool {
	if x < 0 || y < 0 || y >= len(mtx) || x >= len(mtx[0]) {
		return false
	}
	return mtx[y][x] == '.'
}

func main() {
	mtx, guard := readInput("input.txt")
	/*mtx := [][]rune{
		[]rune("....#....."),
		[]rune(".........#"),
		[]rune(".........."),
		[]rune("..#......."),
		[]rune(".......#.."),
		[]rune(".........."),
		[]rune(".#..^....."),
		[]rune("........#."),
		[]rune("#........."),
		[]rune("......#..."),
	}
	guard := [2]int{4, 6}*/
	positions1 := getPath(&mtx, &guard, true)
	fmt.Println("Part 1:", len(positions1))
	positions := getPath(&mtx, &guard, false)
	variations := obstacleCandidates(mtx, positions)
	fmt.Println("Part 2:", variations)
}
