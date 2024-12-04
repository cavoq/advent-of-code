package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func readInput(path string) [][]rune {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error reading file")
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mtx := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		mtx = append(mtx, []rune(line))
	}

	return mtx
}

func getDiagonals(mtx [][]rune, bltr bool) [][]rune {
	rows := len(mtx)
	cols := len(mtx[0])

	diagonals := make([][]rune, rows+cols-1)

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if bltr {
				diagonals[row+col] = append(diagonals[row+col], mtx[row][col])
			} else {
				diagonals[col-row+(rows-1)] = append(diagonals[col-row+(rows-1)], mtx[row][col])
			}
		}
	}

	return diagonals
}

func transpose(mtx [][]rune) [][]rune {
	rows := len(mtx)
	cols := len(mtx[0])
	transposed := make([][]rune, cols)

	for i := range transposed {
		transposed[i] = make([]rune, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[j][i] = mtx[i][j]
		}
	}

	return transposed
}

func countOccurences(mtx [][]rune) int {
	sum := 0
	expr := regexp.MustCompile(`XMAS`)
	revExpr := regexp.MustCompile(`SAMX`)

	for line := range mtx {
		sum += len(expr.FindAllString(string(mtx[line]), -1))
		sum += len(revExpr.FindAllString(string(mtx[line]), -1))
	}

	return sum
}

func solve(mtx [][]rune) int {
	diagonals := append(getDiagonals(mtx, true), getDiagonals(mtx, false)...)
	mtxT := transpose(mtx)
	sum := countOccurences(mtx)
	sum += countOccurences(mtxT)
	sum += countOccurences(diagonals)
	return sum
}

func solve2(mtx [][]rune) int {
	occ := 0
	for i := 1; i < len(mtx)-1; i++ {
		for j := 1; j < len(mtx[0])-1; j++ {
			if mtx[i][j] != 'A' {
				continue
			}
			if ((mtx[i-1][j-1] == 'S' && mtx[i+1][j+1] == 'M') || (mtx[i-1][j-1] == 'M' && mtx[i+1][j+1] == 'S')) &&
				((mtx[i-1][j+1] == 'S' && mtx[i+1][j-1] == 'M') || (mtx[i-1][j+1] == 'M' && mtx[i+1][j-1] == 'S')) {
				occ++
			}
		}
	}
	return occ
}

func main() {
	mtx := readInput("input.txt")
	/*mtx := [][]byte{
	    []byte("MMMSXXMASM"),
	    []byte("MSAMXMSMSA"),
	    []byte("AMXSXMAAMM"),
	    []byte("MSAMASMSMX"),
	    []byte("XMASAMXAMM"),
	    []byte("XXAMMXXAMA"),
	    []byte("SMSMSASXSS"),
	    []byte("SAXAMASAAA"),
	    []byte("MAMMMXMMMM"),
	    []byte("MXMXAXMASX"),
	}*/
	occ := solve(mtx)
	occ2 := solve2(mtx)
	fmt.Println("Part 1:", occ)
	fmt.Println("Part 2:", occ2)
}
