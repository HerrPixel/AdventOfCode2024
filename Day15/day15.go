package day15

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
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

// We simulate each move as a two-step process
// First we test if the move is valid by recursively testing each pushed chest if it could move in that direction
// If any chest can't, the move is invalid
// Only if it is valid, we actually push the chests
// We do this by first recursively push the last chest and then in rolling up the push stack, we push the previous chests.
func GPSCoordinates() string {
	grid, x, y, moves := parseInput()

	for _, i := range moves {
		if canMove(grid, x, y, i) {
			grid, x, y = move(grid, x, y, i)
		}
	}

	return strconv.Itoa(gpsScore(grid))
}

// Same as Part 1
// We first test the move if it is valid by testing each affected chest if it can be pushed
// And only execute the move if that is the case
func DoubledChests() string {
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

// Recursively test if a move is valid by testing if we can move a space in that direction and every pushed object can as well
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

	// If we move up or down and the chest is a double chest, we also need to check if we can push the other half
	if grid[x][y] == "[" && direction%2 == 0 {
		successorCanMove = successorCanMove && canMove(grid, x, y+1, direction)
	}

	if grid[x][y] == "]" && direction%2 == 0 {
		successorCanMove = successorCanMove && canMove(grid, x, y-1, direction)
	}

	return successorCanMove
}

// Move the chests, assuming there is enough free space
func move(grid [][]string, x int, y int, direction int) ([][]string, int, int) {
	if grid[x][y] == "." {
		return grid, x, y
	}

	directions := [4]tuple{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	newX := x + directions[direction].x
	newY := y + directions[direction].y

	grid, _, _ = move(grid, newX, newY, direction)

	// If we move up or down and the chest is a double chest, we also need to push the other half
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

// stretches the warehouse horizontally
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
