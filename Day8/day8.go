package day8

import (
	"strconv"
	"strings"

	"git.jonasseiler.de/Jonas/AdventOfCode2024/Tools"
)

type tuple struct {
	x int
	y int
}

func parseInput() (map[string][]tuple, int, int) {
	lines := Tools.ReadByLines("./Day8/input.txt")
	height := len(lines)
	width := len(lines[0])
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

func Part1() string {
	antennas, height, width := parseInput()

	antinodes := make(map[tuple]bool, 0)
	for _, coords := range antennas {
		for i, a := range coords {
			for j, b := range coords {
				if j <= i {
					continue
				}

				//fmt.Println("Testing", a, "and", b)

				difX := a.x - b.x
				difY := a.y - b.y

				//fmt.Println("difX", difX, "difY", difY)

				if a.x+difX == b.x && a.y+difY == b.y {
					newX := a.x - difX
					newY := a.y - difY

					if newX >= 0 && newY >= 0 && newX < height && newY < width {
						newCoords := tuple{newX, newY}
						//fmt.Println("newCoords", newCoords)

						antinodes[newCoords] = true
					}

					newX = b.x + difX
					newY = b.y + difY

					if newX >= 0 && newY >= 0 && newX < height && newY < width {
						newCoords := tuple{newX, newY}
						//fmt.Println("newCoords", newCoords)

						antinodes[newCoords] = true
					}
				} else {
					newX := a.x + difX
					newY := a.y + difY

					if newX >= 0 && newY >= 0 && newX < height && newY < width {
						newCoords := tuple{newX, newY}
						//fmt.Println("newCoords", newCoords)

						antinodes[newCoords] = true
					}

					newX = b.x - difX
					newY = b.y - difY

					if newX >= 0 && newY >= 0 && newX < height && newY < width {
						newCoords := tuple{newX, newY}
						//fmt.Println("newCoords", newCoords)

						antinodes[newCoords] = true
					}
				}
			}
		}
	}

	total := len(antinodes)

	//fmt.Println(antinodes)

	return strconv.Itoa(total)
}

func Part2() string {
	antennas, height, width := parseInput()

	antinodes := make(map[tuple]bool, 0)
	for _, coords := range antennas {
		for i, a := range coords {
			for j, b := range coords {
				if j <= i {
					continue
				}

				//fmt.Println("Testing", a, "and", b)

				difX := a.x - b.x
				difY := a.y - b.y

				currX := a.x
				currY := a.y

				isOnBoard := true

				for isOnBoard {
					antinodes[tuple{currX, currY}] = true

					currX = currX - difX
					currY = currY - difY

					isOnBoard = currX >= 0 && currY >= 0 && currX < height && currY < width
				}

				isOnBoard = true
				currX = a.x
				currY = a.y

				for isOnBoard {
					antinodes[tuple{currX, currY}] = true

					currX = currX + difX
					currY = currY + difY

					isOnBoard = currX >= 0 && currY >= 0 && currX < height && currY < width
				}

				//fmt.Println("difX", difX, "difY", difY)

				/*
					if a.x+difX == b.x && a.y+difY == b.y {
						newX := a.x - difX
						newY := a.y - difY

						if newX >= 0 && newY >= 0 && newX < height && newY < width {
							newCoords := tuple{newX, newY}
							//fmt.Println("newCoords", newCoords)

							antinodes[newCoords] = true
						}

						newX = b.x + difX
						newY = b.y + difY

						if newX >= 0 && newY >= 0 && newX < height && newY < width {
							newCoords := tuple{newX, newY}
							//fmt.Println("newCoords", newCoords)

							antinodes[newCoords] = true
						}
					} else {
						newX := a.x + difX
						newY := a.y + difY

						if newX >= 0 && newY >= 0 && newX < height && newY < width {
							newCoords := tuple{newX, newY}
							//fmt.Println("newCoords", newCoords)

							antinodes[newCoords] = true
						}

						newX = b.x - difX
						newY = b.y - difY

						if newX >= 0 && newY >= 0 && newX < height && newY < width {
							newCoords := tuple{newX, newY}
							//fmt.Println("newCoords", newCoords)

							antinodes[newCoords] = true
						}
					}
				*/
			}
		}
	}

	total := len(antinodes)

	//fmt.Println(antinodes)

	return strconv.Itoa(total)
}
