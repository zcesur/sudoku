package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const boardSize = 9
const boxSize = 3

// Board represents a Sudoku board.
type Board [boardSize][boardSize]int

// Solve attempts to solve the given Sudoku board.
func (b Board) Solve() (Board, bool) {
	for i, row := range b {
		for j, cell := range row {
			if cell != 0 {
				continue
			}

			for val := 1; val <= boardSize; val++ {
				b[i][j] = val

				if !(b.rows().ok() && b.cols().ok() && b.boxs().ok()) {
					continue
				}

				bSolved, ok := b.Solve()
				if !ok {
					continue
				}

				return bSolved, true
			}
			return b, false
		}
	}
	return b, true
}

// ParseBoard interprets the given string as a Sudoku board.
func ParseBoard(s string) (Board, error) {
	var b Board
	if len(s) != boardSize*boardSize {
		return b, fmt.Errorf("Invalid length")
	}

	for i := range b {
		for j := range b[i] {
			pos := i*boardSize + j
			curStr := s[pos : pos+1]

			if curStr == "." {
				continue
			}

			v, err := strconv.Atoi(curStr)
			if err != nil {
				return b, err
			}

			b[i][j] = v
		}
	}

	return b, nil
}

func (b Board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, cell := range row {
			sb.WriteString(strconv.Itoa(cell))
		}
	}
	return sb.String()
}

func (b Board) rows() Board {
	return b
}

func (b Board) cols() Board {
	var res Board
	for i := range b {
		for j := range b[i] {
			res[j][i] = b[i][j]
		}
	}
	return res
}

func (b Board) boxs() Board {
	var res Board
	for i := range b {
		for j := range b[i] {
			iNew := boxSize*(i/boxSize) + j/boxSize
			jNew := boxSize*(i%boxSize) + j%boxSize
			res[iNew][jNew] = b[i][j]
		}
	}
	return res
}

func (b Board) ok() bool {
	for _, row := range b {
		m := map[int]bool{}
		for _, cell := range row {
			if cell == 0 {
				continue
			}

			if _, ok := m[cell]; ok {
				return false
			}

			m[cell] = true
		}
	}
	return true
}

func (b Board) pprint() {
	for _, row := range b {
		for _, cell := range row {
			fmt.Printf("%v ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	var s string
	_, err := fmt.Scanf("%s", &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		return
	}

	b, err := ParseBoard(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ParseBoard: %v\n", err)
		return
	}

	bSolved, ok := b.Solve()
	if !ok {
		fmt.Fprintf(os.Stderr, "No solution exists")
		return
	}

	bSolved.pprint()
}
