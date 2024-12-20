package day10

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

func parseInput() ([][]int, []tuple) {
	grid := make([][]int, 0)
	trailheads := make([]tuple, 0)

	for i, l := range Tools.ReadByLines("./Day10/input.txt") {
		grid = append(grid, make([]int, len(l)))
		for j, c := range strings.Split(l, "") {
			n, _ := strconv.Atoi(c)
			grid[i][j] = n
			if n == 0 {
				trailheads = append(trailheads, tuple{i, j})
			}
		}
	}

	return grid, trailheads
}

func HikingTrails() string {
	grid, trailheads := parseInput()

	total := 0

	for _, t := range trailheads {
		total += bfs(grid, t.x, t.y)
	}

	return strconv.Itoa(total)
}

func DistinctHikingTrails() string {
	grid, trailheads := parseInput()

	total := 0

	for _, t := range trailheads {
		total += dfs(grid, t.x, t.y)
	}

	return strconv.Itoa(total)
}

func bfs(grid [][]int, x int, y int) int {
	positions := make(map[tuple]bool, 1)
	positions[tuple{x, y}] = true
	width := len(grid)
	height := len(grid[0])
	movements := [4]tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for range 9 {
		newPositions := make(map[tuple]bool, 0)
		for p := range positions {
			x := p.x
			y := p.y

			for _, m := range movements {
				deltaX := m.x
				deltaY := m.y

				if isInBounds(x+deltaX, y+deltaY, width, height) && grid[x+deltaX][y+deltaY] == grid[x][y]+1 {
					newPositions[tuple{x + deltaX, y + deltaY}] = true
				}
			}
		}

		positions = newPositions
	}

	return len(positions)
}

func isInBounds(x int, y int, width int, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}

func dfs(grid [][]int, x int, y int) int {
	width := len(grid)
	height := len(grid[0])
	movements := [4]tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	total := 0

	if grid[x][y] == 9 {
		return 1
	}

	for _, m := range movements {
		deltaX := m.x
		deltaY := m.y

		if isInBounds(x+deltaX, y+deltaY, width, height) && grid[x+deltaX][y+deltaY] == grid[x][y]+1 {
			total += dfs(grid, x+deltaX, y+deltaY)
		}
	}
	return total
}
