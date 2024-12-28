package day18

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

func parseInput() []tuple {
	bytes := make([]tuple, 0)

	for _, s := range Tools.ReadByLines("./Day18/input.txt") {
		numbers := strings.Split(s, ",")

		x, _ := strconv.Atoi(numbers[0])
		y, _ := strconv.Atoi(numbers[1])

		bytes = append(bytes, tuple{x, y})
	}

	return bytes
}

// We just use bfs on the grid created by placing the first 1024 bytes
func ByteRainPath() string {
	bytes := parseInput()

	width := 71
	height := 71

	grid := gridFromBytes(bytes, 1024, width, height)

	return strconv.Itoa(bfs(grid, 0, 0, width-1, height-1))
}

// We use binary search on the number of bytes we are placing until we find the last working number
func LatestByteRainPath() string {
	bytes := parseInput()

	width := 71
	height := 71

	left := 0
	right := len(bytes) - 1

	for left < right-1 {
		middle := (left + right) / 2

		grid := gridFromBytes(bytes, middle, width, height)

		// If it is impossible now, try less, otherwise if it is still possible, try more
		if bfs(grid, 0, 0, width-1, height-1) == 0 {
			right = middle
		} else {
			left = middle
		}
	}

	return strconv.Itoa(bytes[left].x) + "," + strconv.Itoa(bytes[left].y)
}

// standard bfs on a 2D-grid
func bfs(grid [][]bool, startX int, startY int, endX int, endY int) int {
	width := len(grid)
	height := len(grid[0])

	positions := make([]tuple, 0)
	positions = append(positions, tuple{startX, startY})

	movements := [4]tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	isInBounds := func(x int, y int) bool {
		return 0 <= x && x < width && 0 <= y && y < height
	}

	visited := make([][]bool, width)
	for i := range width {
		visited[i] = make([]bool, height)
	}
	visited[startX][startY] = true

	steps := 0

	for len(positions) > 0 {
		newPositions := make([]tuple, 0)

		for _, pos := range positions {
			x := pos.x
			y := pos.y

			if x == endX && y == endY {
				return steps
			}

			for _, move := range movements {
				if isInBounds(x+move.x, y+move.y) && !grid[x+move.x][y+move.y] && !visited[x+move.x][y+move.y] {
					newPositions = append(newPositions, tuple{x + move.x, y + move.y})
					visited[x+move.x][y+move.y] = true
				}
			}
		}

		steps++
		positions = newPositions

	}

	return 0
}

func gridFromBytes(bytes []tuple, n int, width int, height int) [][]bool {
	grid := make([][]bool, width)
	for i := range width {
		grid[i] = make([]bool, height)
	}

	for i, b := range bytes {
		if i >= n {
			break
		}
		grid[b.x][b.y] = true
	}

	return grid
}
