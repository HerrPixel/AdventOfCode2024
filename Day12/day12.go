package day12

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

type zone struct {
	area      int
	perimeter int
	sides     int
}

func readInput() ([][]string, int, int) {
	grid := make([][]string, 0)

	for i, s := range Tools.ReadByLines("./Day12/input.txt") {
		grid = append(grid, make([]string, 0))
		grid[i] = append(grid[i], strings.Split(s, "")...)
	}

	width := len(grid)
	height := len(grid[0])

	return grid, width, height
}

func PerimeterFences() string {
	grid, width, height := readInput()

	zones := floodFill(grid, width, height)

	total := 0

	for _, z := range zones {
		total += z.area * z.perimeter
	}

	return strconv.Itoa(total)
}

func SideFences() string {
	grid, width, height := readInput()

	zones := floodFill(grid, width, height)

	total := 0

	for _, z := range zones {
		total += z.area * z.sides
	}

	return strconv.Itoa(total)
}

func floodFill(grid [][]string, width int, height int) []zone {
	filled := make([][]bool, width)
	for i := range width {
		filled[i] = make([]bool, height)
	}

	isInBounds := func(a int, b int) bool {
		return 0 <= a && a < width && 0 <= b && b < height
	}

	movements := [4]tuple{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	zones := make([]zone, 0)

	for i := range width {
		for j := range height {
			if filled[i][j] {
				continue
			}

			var q Tools.Queue[tuple]

			q = Tools.Enqueue(q, tuple{i, j})

			filled[i][j] = true

			area := 0
			perimeter := 0
			sides := 0

			for !Tools.IsEmpty(q) {

				var plot tuple
				plot, q, _ = Tools.Dequeue(q)

				x := plot.x
				y := plot.y
				c := grid[x][y]

				area += 1
				sides += cornerMultiplicity(grid, x, y, width, height)

				for _, movement := range movements {
					deltaX := movement.x
					deltaY := movement.y

					if !isInBounds(x+deltaX, y+deltaY) || grid[x+deltaX][y+deltaY] != c {
						perimeter++
						continue
					}

					if !filled[x+deltaX][y+deltaY] {
						q = Tools.Enqueue(q, tuple{x + deltaX, y + deltaY})
						filled[x+deltaX][y+deltaY] = true
					}
				}
			}

			zones = append(zones, zone{area, perimeter, sides})
		}
	}

	return zones
}

func cornerMultiplicity(grid [][]string, x int, y int, width int, height int) int {
	multiplicity := 0

	isInBounds := func(a int, b int) bool {
		return 0 <= a && a < width && 0 <= b && b < height
	}

	isDifferent := func(a int, b int) bool {
		return !isInBounds(a, b) || grid[a][b] != grid[x][y]
	}

	orthogonalSpaces := [4][2]tuple{{{1, 0}, {0, 1}},
		{{0, 1}, {-1, 0}},
		{{-1, 0}, {0, -1}},
		{{0, -1}, {1, 0}}}

	for _, c := range orthogonalSpaces {
		a := c[0]
		b := c[1]

		if isDifferent(x+a.x, y+a.y) && isDifferent(x+b.x, y+b.y) {
			multiplicity++
		} else if !isDifferent(x+a.x, y+a.y) && !isDifferent(x+b.x, y+b.y) && isDifferent(x+a.x+b.x, y+a.y+b.y) {
			multiplicity++
		}
	}

	return multiplicity
}
