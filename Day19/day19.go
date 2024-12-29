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

	towels = append(towels, strings.Split(input[0], ", ")...)
	designs = append(designs, strings.Split(input[1], "\n")...)

	return towels, designs
}

// We use Dynamic Programming to keep track of already seen partial designs and if they are possible to make
// Then just test each pattern if it is a prefix and recurse into the design without this prefix
func PossibleDesigns() string {
	towels, designs := parseInput()

	total := 0
	for _, d := range designs {
		// Our function outputs the number of combinations of each design
		// In Part 1 we are only interested if they are possible, not in how many ways, so we need to just check if #combinations > 0
		i := possibilities(d, towels, make(map[string]int, 0))
		if i > 0 {
			total += 1
		}
	}

	return strconv.Itoa(total)
}

// Same as Part 1 but we save the number of combinations per pattern, not the possibility itself
func DesignCombinations() string {
	towels, designs := parseInput()
	total := 0
	for _, d := range designs {
		i := possibilities(d, towels, make(map[string]int, 0))
		total += i
	}

	return strconv.Itoa(total)
}

// returns the number of possible combinations of design given the towel patterns
// Uses memoization of the number of combinations for a partial design
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

// Tests if b is a prefix of a
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
