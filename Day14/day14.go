package day14

import (
	"regexp"
	"strconv"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type robot struct {
	x  int
	y  int
	dx int
	dy int
}

func parseInput() []robot {
	r := regexp.MustCompile(`p=([0-9]+),([0-9]+) v=(-?[0-9]+),(-?[0-9]+)`)
	robots := make([]robot, 0)
	for _, s := range Tools.ReadByLines("./Day14/input.txt") {
		captures := r.FindStringSubmatch(s)

		x, _ := strconv.Atoi(captures[1])
		y, _ := strconv.Atoi(captures[2])
		dx, _ := strconv.Atoi(captures[3])
		dy, _ := strconv.Atoi(captures[4])

		robots = append(robots, robot{x, y, dx, dy})
	}

	return robots
}

func Bathroom() string {
	robots := parseInput()

	width := 101
	height := 103

	for i, r := range robots {
		robots[i].x = mod(r.x+100*r.dx, width)
		robots[i].y = mod(r.y+100*r.dy, height)
	}

	return strconv.Itoa(safetyFactor(robots, width, height))
}

func EasterEgg() string {
	robots := parseInput()

	width := 101
	height := 103

	round := 0

	for {
		for j, r := range robots {
			robots[j].x = mod(r.x+r.dx, width)
			robots[j].y = mod(r.y+r.dy, height)
		}
		round++

		if isTreeProbably(robots) {
			return strconv.Itoa(round)
		}
	}
}

func mod(n int, m int) int {
	return (n%m + m) % m
}

func safetyFactor(robots []robot, width int, height int) int {
	upperLeft := 0
	upperRight := 0
	lowerLeft := 0
	lowerRight := 0

	middleX := (width - 1) / 2
	middleY := (height - 1) / 2

	for _, r := range robots {
		if r.x < middleX && r.y < middleY {
			upperLeft++
		} else if r.x < middleX && r.y > middleY {
			lowerLeft++
		} else if r.x > middleX && r.y < middleY {
			upperRight++
		} else if r.x > middleX && r.y > middleY {
			lowerRight++
		}
	}

	return upperLeft * upperRight * lowerLeft * lowerRight
}

func isTreeProbably(robots []robot) bool {
	type tuple struct {
		x int
		y int
	}

	positions := make(map[tuple]bool)

	for _, r := range robots {
		_, exists := positions[tuple{r.x, r.y}]

		if exists {
			return false
		}

		positions[tuple{r.x, r.y}] = true
	}

	return true
}
