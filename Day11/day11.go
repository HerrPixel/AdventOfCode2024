package day11

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type state struct {
	rock  int
	round int
}

func parseInput() []int {
	rocks := make([]int, 0)
	for _, s := range strings.Split(Tools.Read("./Day11/input.txt"), " ") {
		rock, _ := strconv.Atoi(s)
		rocks = append(rocks, rock)
	}
	return rocks
}

// Independently simulating each rock, memoizing the result of each (rock, blink) combination to speed up calculation
func Blinking() string {
	rocks := parseInput()

	return strconv.Itoa(blink(25, rocks))
}

// Same as Part1
// Independently simulating each rock, memoizing the result of each (rock, blink) combination to speed up calculation
func MoreBlinking() string {
	rocks := parseInput()

	return strconv.Itoa(blink(75, rocks))
}

func blink(n int, rocks []int) int {
	dp := make(map[state]int)
	total := 0

	var recurse func(int, int) int
	recurse = func(rock int, round int) int {
		scene := state{rock, round}

		res, ok := dp[scene]

		// If we have seen this rock with the same number of blinks remaining before, return that memoized result
		if ok {
			return res
		}

		// otherwise, when we have blinked enough, stop the recursion
		if round == n {
			return 1
		}

		// blinking procedure
		if rock == 0 {
			res = recurse(1, round+1)
		} else if digits(rock)%2 == 0 {
			left, right := split(rock)
			res = recurse(left, round+1) + recurse(right, round+1)
		} else {
			res = recurse(rock*2024, round+1)
		}

		// memoize this result
		dp[scene] = res
		return res
	}

	for _, rock := range rocks {
		total += recurse(rock, 0)
	}

	return total
}

func digits(n int) int {
	return len(strconv.Itoa(n))
}

func split(n int) (int, int) {
	s := strconv.Itoa(n)
	left, _ := strconv.Atoi(s[:len(s)/2])
	right, _ := strconv.Atoi(s[len(s)/2:])
	return left, right
}
