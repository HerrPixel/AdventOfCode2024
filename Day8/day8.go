package day8

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

func parseInput() (map[string][]tuple, int, int) {
	lines := Tools.ReadByLines("./Day8/input.txt")
	width := len(lines)
	height := len(lines[0])
	antennas := make(map[string][]tuple, 0)

	for i, s := range lines {
		for j, c := range strings.Split(s, "") {
			if c != "." {
				_, ok := antennas[c]

				if !ok {
					antennas[c] = make([]tuple, 0)
				}

				antennas[c] = append(antennas[c], tuple{i, j})
			}
		}
	}

	return antennas, height, width
}

func AntinodePairs() string {
	antennas, height, width := parseInput()

	antinodes := make(map[tuple]bool, 0)
	for _, coords := range antennas {
		for i, a := range coords {
			for j, b := range coords {
				if j <= i {
					continue
				}

				deltaX := a.x - b.x
				deltaY := a.y - b.y

				if isOnBoard(a.x+deltaX, a.y+deltaY, width, height) {
					antinodes[tuple{a.x + deltaX, a.y + deltaY}] = true
				}

				if isOnBoard(b.x-deltaX, b.y-deltaY, width, height) {
					antinodes[tuple{b.x - deltaX, b.y - deltaY}] = true
				}

			}
		}
	}

	total := len(antinodes)

	return strconv.Itoa(total)
}

func Antinodes() string {
	antennas, height, width := parseInput()

	antinodes := make(map[tuple]bool, 0)
	for _, coords := range antennas {
		for i, a := range coords {
			for j, b := range coords {
				if j <= i {
					continue
				}

				deltaX := a.x - b.x
				deltaY := a.y - b.y

				x := a.x
				y := a.y

				for isOnBoard(x, y, width, height) {
					antinodes[tuple{x, y}] = true

					x = x - deltaX
					y = y - deltaY
				}

				x = b.x
				y = b.y

				for isOnBoard(x, y, width, height) {
					antinodes[tuple{x, y}] = true

					x = x + deltaX
					y = y + deltaY
				}
			}
		}
	}

	total := len(antinodes)

	return strconv.Itoa(total)
}

func isOnBoard(x int, y int, width int, height int) bool {
	return x >= 0 && y >= 0 && x < width && y < height
}
