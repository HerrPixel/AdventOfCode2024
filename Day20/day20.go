package day20

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() ([][]int, int, int, int, int) {
	distance := make([][]int, 0)
	startX := 0
	startY := 0
	endX := 0
	endY := 0

	for i, s := range Tools.ReadByLines("./Day20/input.txt") {
		distance = append(distance, make([]int, 0))
		for j, c := range strings.Split(s, "") {
			if c == "S" {
				startX = i
				startY = j
			}

			if c == "E" {
				endX = i
				endY = j
			}

			if c == "#" {
				distance[i] = append(distance[i], -1)
			} else {
				distance[i] = append(distance[i], 0)
			}
		}
	}

	return distance, startX, startY, endX, endY
}

// We assume there is a unique path from start to finish
// We calculate this path and then check for each pair of points on the path, if their manhattan distance is 2 or less but their indices differ by at least 100
// Since the start and end point of a cheat must be points on the unique path and with each cheat, we can travel at most a manhattan distance of 2, this suffices as conditions for cheats
// Similarily, since the path is unique, the number of saved seconds equals the differences in indices of the two points on the path.
func ShortShortcuts() string {
	distance, startX, startY, endX, endY := parseInput()

	path := getPath(distance, startX, startY, endX, endY)

	return strconv.Itoa(countShortcuts(path, 2))
}

// Same as Part 1 but we test if the manhattan distance is 20 or less
func LongShortcuts() string {
	distance, startX, startY, endX, endY := parseInput()

	path := getPath(distance, startX, startY, endX, endY)

	return strconv.Itoa(countShortcuts(path, 20))
}

// returns the unique path from start to finish as a list of 2D coordinates
func getPath(grid [][]int, startX int, startY int, endX int, endY int) [][2]int {
	x := startX
	y := startY

	// we mark visited spaces with their distance
	distance := 1

	path := make([][2]int, 0)

	isInBounds := func(i int, j int) bool {
		return 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i])
	}

	for !(x == endX && y == endY) {
		grid[x][y] = distance
		path = append(path, [2]int{x, y})

		// only if the space is inbounds and unmarked, we visit it
		// except for the endspace, there is always exactly one such space
		if isInBounds(x+1, y) && grid[x+1][y] == 0 {
			x = x + 1
		} else if isInBounds(x, y+1) && grid[x][y+1] == 0 {
			y = y + 1
		} else if isInBounds(x-1, y) && grid[x-1][y] == 0 {
			x = x - 1
		} else if isInBounds(x, y-1) && grid[x][y-1] == 0 {
			y = y - 1
		} else {
		}

		distance += 1
	}

	grid[x][y] = distance
	path = append(path, [2]int{x, y})

	return path
}

// Tests each pair of 2D points for manhattan Distance < shortcutLength and difference of indices > 100
func countShortcuts(path [][2]int, shortcutLength int) int {
	shortcuts := 0

	for i, start := range path {
		for j, end := range path {
			length := Tools.Abs(start[0]-end[0]) + Tools.Abs(start[1]-end[1])

			if j-i < 100+length {
				continue
			}

			if length <= shortcutLength {
				shortcuts++
			}
		}
	}

	return shortcuts
}
