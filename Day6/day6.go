package day6

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type tuple struct {
	a int
	b int
}

func parseInput() ([][]bool, int, int) {
	grid := make([][]bool, 0)
	x := 0
	y := 0

	for i, s := range Tools.ReadByLines("./Day6/input.txt") {
		grid = append(grid, make([]bool, 0))
		for j, c := range strings.Split(s, "") {
			if c == "#" {
				grid[i] = append(grid[i], true)
			} else {
				grid[i] = append(grid[i], false)
			}

			if c == "^" {
				x = i
				y = j
			}
		}
	}

	return grid, x, y
}

// just simulate the guard and sum up the number of visited spaces
func GuardLeave() string {
	grid, x, y := parseInput()

	_, visited := guardPatrol(grid, x, y)

	visitedSpaces := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				visitedSpaces++
			}
		}
	}

	return strconv.Itoa(visitedSpaces)
}

// We only need to place obstructions on spaces in the original path
// For each such space, we simulate the guard on a board with that space obstructed and increment, if that leads to a loop
func Obstruction() string {
	grid, x, y := parseInput()

	_, visited := guardPatrol(grid, x, y)

	total := 0
	for i := range visited {
		for j := range visited[i] {
			if visited[i][j] {
				if i == x && j == y {
					continue
				}

				obstructedGrid := Tools.Clone(grid)
				obstructedGrid[i] = Tools.Clone(grid[i])
				obstructedGrid[i][j] = true

				isLoop, _ := guardPatrol(obstructedGrid, x, y)

				if isLoop {
					total++
				}
			}
		}
	}

	return strconv.Itoa(total)
}

// helper function that charts the guards path
// if it arives at an already seen space with the same orientation, we encountered a loop and we return
// returns a boolean value describing if it encountered a loop or ran out of the board and a copy of the board with spaces visited
func guardPatrol(grid [][]bool, x int, y int) (bool, [][]bool) {
	movements := [4]tuple{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	direction := 0

	alreadySeen := make([][][4]bool, len(grid))

	for i := range grid {
		alreadySeen[i] = make([][4]bool, len(grid[i]))
	}

	onBoard := true

	for onBoard {
		if alreadySeen[x][y][direction] {

			return true, getVisited(alreadySeen)
		}

		alreadySeen[x][y][direction] = true

		nextX := x + movements[direction].a
		nextY := y + movements[direction].b

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			onBoard = false
			break
		}

		if grid[nextX][nextY] {
			direction = (direction + 1) % 4
			continue
		}

		x = nextX
		y = nextY
	}
	return false, getVisited(alreadySeen)
}

// converts the visited datastructure into a simpler copy of the board
func getVisited(v [][][4]bool) [][]bool {
	visited := make([][]bool, len(v))
	for i := range v {
		visited[i] = make([]bool, len(v[i]))

		// squishes the directions into one boolean value
		for j := range v[i] {
			visited[i][j] = v[i][j][0] || v[i][j][1] || v[i][j][2] || v[i][j][3]
		}
	}
	return visited
}
