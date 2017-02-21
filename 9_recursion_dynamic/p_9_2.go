package main

import "fmt"

func genSolution(m [][]int, N, M int) {
	if m[N-1][M-1] == -1 {
		return
	}

	m[N-1][M-1] = 1

	for r := N - 1; r >= 0; r-- {
		for c := M - 1; c >= 0; c-- {
			if m[r][c] == -1 {
				m[r][c] = 0
				continue
			}

			var val int
			if r+1 < N { // && m[r+1][c] != -1 {
				val = m[r+1][c]
			}
			if c+1 < M { // && m[r][c+1] != -1 {
				val += m[r][c+1]
			}
			m[r][c] += val
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 0, 0, 0},
		{0, -1, 0, 0},
		{0, 0, 0, 0},
	}
	genSolution(matrix, len(matrix), len(matrix[0]))

	fmt.Println("solution:", matrix[0][0])

	matrix = [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, -1, 0, 0, 0, -1, 0},
		{0, 0, 0, 0, -1, 0, 0, 0},
		{-1, 0, -1, 0, 0, -1, 0, 0},
		{0, 0, -1, 0, 0, 0, 0, 0},
		{0, 0, 0, -1, -1, 0, -1, 0},
		{0, -1, 0, 0, 0, -1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	genSolution(matrix, len(matrix), len(matrix[0]))

	fmt.Println("solution:", matrix[0][0])
}
