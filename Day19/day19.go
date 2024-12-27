package day19

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() ([]string, []string) {
	input := strings.Split(Tools.Read("./Day19/input.txt"), "\n\n")
	towels := make([]string, 0)
	designs := make([]string, 0)

	for _, s := range strings.Split(input[0], ", ") {
		towels = append(towels, s)
	}

	for _, s := range strings.Split(input[1], "\n") {
		designs = append(designs, s)
	}

	return towels, designs
}

func Part1() string {
	towels, designs := parseInput()

	total := 0
	for _, d := range designs {
		i := possibilities(d, towels, make(map[string]int, 0))
		if i > 0 {
			total += 1
		}
	}

	return strconv.Itoa(total)
}

func Part2() string {
	towels, designs := parseInput()
	total := 0
	for _, d := range designs {
		i := possibilities(d, towels, make(map[string]int, 0))
		total += i
	}

	return strconv.Itoa(total)
}

func possibilities(design string, towels []string, memo map[string]int) int {
	if design == "" {
		return 1
	}

	res, ok := memo[design]

	if ok {
		return res
	}

	combinations := 0

	for _, t := range towels {

		if beginsWith(design, t) {
			res := possibilities(design[len(t):], towels, memo)
			combinations += res
		}
	}

	memo[design] = combinations

	return combinations
}

func beginsWith(a string, b string) bool {
	if len(b) > len(a) {
		return false
	}

	for i := range len(b) {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
