package day4

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() [][]string {
	matrix := make([][]string, 0)

	for i, s := range Tools.ReadByLines("./Day4/input.txt") {
		matrix = append(matrix, make([]string, 0))
		for _, c := range strings.Split(s, "") {
			if c == "" {
				continue
			}
			matrix[i] = append(matrix[i], c)
		}
	}

	return matrix
}

// We pivot around the "X" and check each direction if it forms a valid XMAS starting with the given "X"
func XMAS() string {
	matrix := parseInput()
	total := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "X" {
				if upwards(matrix, i, j) {
					total++
				}
				if downwards(matrix, i, j) {
					total++
				}
				if right(matrix, i, j) {
					total++
				}
				if left(matrix, i, j) {
					total++
				}
				if diagonal_left_downwards(matrix, i, j) {
					total++
				}
				if diagonal_left_upwards(matrix, i, j) {
					total++
				}
				if diagonal_right_donwards(matrix, i, j) {
					total++
				}
				if diagonal_right_upwards(matrix, i, j) {
					total++
				}
			}
		}
	}
	return strconv.Itoa(total)
}

// We pivot around the "A" character
// Since X-Mas's don't overlap, we can naively check if an "A" is a X-Mas
func X_MAS() string {
	matrix := parseInput()
	total := 0

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == "A" {
				if isXMas(matrix, i, j) {
					total++
				}
			}
		}
	}
	return strconv.Itoa(total)
}

func upwards(m [][]string, i int, j int) bool {
	if j < 3 {
		return false
	}

	return (m[i][j-1] == "M") && (m[i][j-2] == "A") && (m[i][j-3] == "S")
}

func downwards(m [][]string, i int, j int) bool {
	if j > len(m)-4 {
		return false
	}

	return (m[i][j+1] == "M") && (m[i][j+2] == "A") && (m[i][j+3] == "S")
}

func right(m [][]string, i int, j int) bool {
	if i > len(m[0])-4 {
		return false
	}

	return (m[i+1][j] == "M") && (m[i+2][j] == "A") && (m[i+3][j] == "S")
}

func left(m [][]string, i int, j int) bool {
	if i < 3 {
		return false
	}

	return (m[i-1][j] == "M") && (m[i-2][j] == "A") && (m[i-3][j] == "S")
}

func diagonal_right_upwards(m [][]string, i int, j int) bool {
	if (i > len(m[0])-4) || (j < 3) {
		return false
	}

	return (m[i+1][j-1] == "M") && (m[i+2][j-2] == "A") && (m[i+3][j-3] == "S")
}

func diagonal_left_upwards(m [][]string, i int, j int) bool {
	if (i < 3) || (j < 3) {
		return false
	}

	return (m[i-1][j-1] == "M") && (m[i-2][j-2] == "A") && (m[i-3][j-3] == "S")
}

func diagonal_right_donwards(m [][]string, i int, j int) bool {
	if (i > len(m[0])-4) || (j > len(m)-4) {
		return false
	}

	return (m[i+1][j+1] == "M") && (m[i+2][j+2] == "A") && (m[i+3][j+3] == "S")
}

func diagonal_left_downwards(m [][]string, i int, j int) bool {
	if (i < 3) || j > (len(m)-4) {
		return false
	}

	return (m[i-1][j+1] == "M") && (m[i-2][j+2] == "A") && (m[i-3][j+3] == "S")
}

// diagonal elements must be "M" or "S" and different, then the anti-diagonal elements must match exactly one of the diagonal elements
func isXMas(m [][]string, i int, j int) bool {
	if (i < 1) || (j < 1) || (i > len(m[0])-2) || (j > len(m)-2) {
		return false
	}

	c := m[i-1][j-1]
	d := m[i+1][j+1]

	if (c == "M" && d == "S") || (c == "S" && d == "M") {
		return (m[i+1][j-1] == c && m[i-1][j+1] == d) || (m[i-1][j+1] == c && m[i+1][j-1] == d)
	}
	return false
}
