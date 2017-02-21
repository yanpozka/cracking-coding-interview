package main

import (
	"fmt"
	"os"
)

type piece byte

const (
	empty piece = iota
	queen
	occup
)

func (i piece) String() string {
	switch i {
	case empty:
		return "_"
	case queen:
		return "Q"
	case occup:
		return "*"
	}
	return "no"
}

var count int

func queens(m [][]piece, row, col, level int) {
	fmt.Println(row, col, level)

	mark(m, row, col)
	m[row][col] = queen

	if level == 8 {
		for _, r := range m {
			fmt.Println(r)
		}
		fmt.Println()
		count++
		if count > 6 {
			os.Exit(0)
		}
		return
	}

	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m); c++ {
			if m[r][c] == empty {
				newMtx := cloneM(m)
				queens(newMtx, r, c, level+1)
			}
		}
	}
}

func main() {
	matrix := make([][]piece, 8)
	for ix := range matrix {
		matrix[ix] = make([]piece, 8)
	}

	queens(matrix, 0, 0, 1)
}

func cloneM(m [][]piece) [][]piece {
	matrix := make([][]piece, len(m))
	for ix := range matrix {
		matrix[ix] = make([]piece, len(m))

		for k := 0; k < len(m); k++ {
			matrix[ix][k] = m[ix][k]
		}
	}

	return matrix
}

func mark(m [][]piece, row, col int) {
	for ix := 0; ix < len(m); ix++ {
		m[row][ix] = occup
	}

	for ix := 0; ix < len(m); ix++ {
		m[ix][col] = occup
	}

	for r, c := row-1, col-1; r >= 0 && c >= 0; r-- { // <\
		m[r][c] = occup
		c--
	}
	for r, c := row+1, col+1; r < len(m) && c < len(m); r++ { // \>
		m[r][c] = occup
		c++
	}

	for r, c := row+1, col-1; r < len(m) && c >= 0; r++ { // </
		m[r][c] = occup
		c--
	}
	for r, c := row-1, col+1; r >= 0 && c < len(m); r-- { // />
		m[r][c] = occup
		c++
	}
}
