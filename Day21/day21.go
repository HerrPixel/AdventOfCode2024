package day21

import (
	"math"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type state struct {
	x        int
	y        int
	sequence []int
}

func parseInput() ([][]int, []int) {
	codes := make([][]int, 0)
	numericalValues := make([]int, 0)

	for _, l := range Tools.ReadByLines("./Day21/input.txt") {
		code := make([]int, 0)

		for _, c := range strings.Split(l, "") {
			if c == "A" {
				code = append(code, 10)
			} else {
				n, _ := strconv.Atoi(c)
				code = append(code, n)
			}
		}

		numericalValue, _ := strconv.Atoi(l[:len(l)-1])

		codes = append(codes, code)
		numericalValues = append(numericalValues, numericalValue)
	}

	return codes, numericalValues
}

func Part1() string {
	codes, numericalValues := parseInput()

	total := 0
	for i, code := range codes {
		res := shortestNumpadSequence(code, 2)
		total += res * numericalValues[i]
	}

	return strconv.Itoa(total)
}

func Part2() string {
	codes, numericalValues := parseInput()

	total := 0
	for i, code := range codes {
		res := shortestNumpadSequence(code, 25)
		total += res * numericalValues[i]
	}

	return strconv.Itoa(total)
}

func shortestNumpadSequence(code []int, indirections int) int {
	total := 0

	position := 10

	for _, button := range code {
		moves := shortestNumpadMoves(button, position)

		shortestInputLength := math.MaxInt64
		for _, m := range moves {
			shortestInputLength = min(shortestInputLength, shortestDirectionalPadSequence(m, indirections))
		}

		total += shortestInputLength
		position = button
	}

	return total
}

func shortestNumpadMoves(destination int, start int) [][]int {
	codes := [11][2]int{
		{3, 1}, // 0
		{2, 0}, // 1
		{2, 1}, // 2
		{2, 2}, // 3
		{1, 0}, // 4
		{1, 1}, // 5
		{1, 2}, // 6
		{0, 0}, // 7
		{0, 1}, // 8
		{0, 2}, // 9
		{3, 2}, // A
	}

	startX := codes[start][0]
	startY := codes[start][1]

	endX := codes[destination][0]
	endY := codes[destination][1]

	sequences := make([][]int, 0)

	q := make([]state, 0)
	q = append(q, state{startX, startY, []int{}})

	for len(q) > 0 {
		newQ := make([]state, 0)

		for _, s := range q {
			x := s.x
			y := s.y

			if x == endX && y == endY {
				sequences = append(sequences, append(s.sequence, 4))
				continue
			}

			if x == 3 && y == 0 {
				continue
			}

			if x < endX {
				newQ = append(newQ, state{x + 1, y, append(s.sequence, 0)})
			} else if x > endX {
				newQ = append(newQ, state{x - 1, y, append(s.sequence, 1)})
			}

			if y < endY {
				newQ = append(newQ, state{x, y + 1, append(s.sequence, 2)})
			} else if y > endY {
				newQ = append(newQ, state{x, y - 1, append(s.sequence, 3)})
			}
		}

		q = newQ
	}
	return sequences
}

type hash struct {
	start       int
	destination int
	depth       int
}

var shortestSequence = make(map[hash]int, 0)

func shortestDirectionalPadSequence(code []int, depth int) int {

	if depth == 0 {
		return len(code)
	}

	total := 0

	position := 4

	for _, button := range code {

		total += shortestDirectionalPadMove(button, position, depth)
		position = button
	}

	return total
}

func shortestDirectionalPadMove(destination int, start int, depth int) int {
	res, ok := shortestSequence[hash{start, destination, depth}]

	if ok {
		return res
	}

	codes := [5][2]int{
		{1, 1}, // v
		{0, 1}, // ^
		{1, 2}, // >
		{1, 0}, // <
		{0, 2}, // A
	}

	startX := codes[start][0]
	startY := codes[start][1]

	endX := codes[destination][0]
	endY := codes[destination][1]

	bestSeen := math.MaxInt64

	q := make([]state, 0)
	q = append(q, state{startX, startY, []int{}})

	for len(q) > 0 {
		newQ := make([]state, 0)

		for _, s := range q {
			x := s.x
			y := s.y

			if x == endX && y == endY {
				bestSeen = min(bestSeen, shortestDirectionalPadSequence(append(s.sequence, 4), depth-1))
				continue
			}

			if x == 0 && y == 0 {
				continue
			}

			if x < endX {
				newQ = append(newQ, state{x + 1, y, append(s.sequence, 0)})
			} else if x > endX {
				newQ = append(newQ, state{x - 1, y, append(s.sequence, 1)})
			}

			if y < endY {
				newQ = append(newQ, state{x, y + 1, append(s.sequence, 2)})
			} else if y > endY {
				newQ = append(newQ, state{x, y - 1, append(s.sequence, 3)})
			}
		}

		q = newQ
	}

	shortestSequence[hash{start, destination, depth}] = bestSeen
	return bestSeen
}
