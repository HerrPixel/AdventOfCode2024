package day15

import (
	"strconv"
	"strings"

	"git.jonasseiler.de/Jonas/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

func parseInput() ([][]string, int, int, []int) {
	input := strings.Split(Tools.Read("./Day15/input.txt"), "\n\n")

	grid := make([][]string, 0)
	x := 0
	y := 0

	for i, s := range strings.Split(input[0], "\n") {
		grid = append(grid, make([]string, 0))

		for j, c := range strings.Split(s, "") {
			if c == "@" {
				x = i
				y = j
			}
			grid[i] = append(grid[i], c)
		}
	}

	moves := make([]int, 0)
	for _, c := range strings.Split(input[1], "") {
		if c == "^" {
			moves = append(moves, 0)
		} else if c == ">" {
			moves = append(moves, 1)
		} else if c == "v" {
			moves = append(moves, 2)
		} else if c == "<" {
			moves = append(moves, 3)
		}
	}

	return grid, x, y, moves
}

func Part1() string {
	grid, x, y, moves := parseInput()

	for _, i := range moves {
		if canMove(grid, x, y, i) {
			grid, x, y = move(grid, x, y, i)
		}
	}

	return strconv.Itoa(gpsScore(grid))
}

func Part2() string {
	grid, x, y, moves := parseInput()

	y *= 2

	grid = doubleWarehouse(grid)

	for _, i := range moves {
		if canMove(grid, x, y, i) {
			grid, x, y = move(grid, x, y, i)
		}
	}

	return strconv.Itoa(gpsScore(grid))
}

func canMove(grid [][]string, x int, y int, direction int) bool {
	if grid[x][y] == "." {
		return true
	}

	if grid[x][y] == "#" {
		return false
	}

	directions := [4]tuple{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	x += directions[direction].x
	y += directions[direction].y

	successorCanMove := canMove(grid, x, y, direction)

	if grid[x][y] == "[" && direction%2 == 0 {
		successorCanMove = successorCanMove && canMove(grid, x, y+1, direction)
	}

	if grid[x][y] == "]" && direction%2 == 0 {
		successorCanMove = successorCanMove && canMove(grid, x, y-1, direction)
	}

	return successorCanMove
}

func move(grid [][]string, x int, y int, direction int) ([][]string, int, int) {
	if grid[x][y] == "." {
		return grid, x, y
	}

	directions := [4]tuple{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	newX := x + directions[direction].x
	newY := y + directions[direction].y

	grid, _, _ = move(grid, newX, newY, direction)

	if grid[x][y] == "[" && direction%2 == 0 {
		grid, _, _ = move(grid, newX, newY+1, direction)
		grid[newX][newY+1] = "]"
		grid[x][y+1] = "."
	}

	if grid[x][y] == "]" && direction%2 == 0 {
		grid, _, _ = move(grid, newX, newY-1, direction)
		grid[newX][newY-1] = "["
		grid[x][y-1] = "."
	}

	grid[newX][newY] = grid[x][y]
	grid[x][y] = "."

	return grid, newX, newY
}

func doubleWarehouse(grid [][]string) [][]string {
	warehouse := make([][]string, 0)

	for i := range grid {
		warehouse = append(warehouse, make([]string, 0))

		for _, c := range grid[i] {
			if c == "@" {
				warehouse[i] = append(warehouse[i], "@", ".")
			} else if c == "." {
				warehouse[i] = append(warehouse[i], ".", ".")
			} else if c == "#" {
				warehouse[i] = append(warehouse[i], "#", "#")
			} else if c == "O" {
				warehouse[i] = append(warehouse[i], "[", "]")
			}
		}
	}

	return warehouse
}

func gpsScore(grid [][]string) int {
	total := 0

	for i, s := range grid {
		for j, c := range s {
			if c == "O" || c == "[" {
				total += 100*i + j
			}
		}
	}
	return total
}
