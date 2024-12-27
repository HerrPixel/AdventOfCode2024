package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() (map[int][]int, [][]int) {
	input := strings.Split(Tools.Read("./Day5/input.txt"), "\n\n")
	orderings := make(map[int][]int)
	for _, s := range Tools.Lines(input[0]) {
		order := strings.Split(s, "|")
		x, _ := strconv.Atoi(order[0])
		y, _ := strconv.Atoi(order[1])

		_, exists := orderings[x]

		if !exists {
			orderings[x] = make([]int, 0)
		}
		orderings[x] = append(orderings[x], y)
	}

	lists := make([][]int, 0)
	for _, s := range Tools.Lines(input[1]) {
		lists = append(lists, make([]int, 0))
		for _, i := range strings.Split(s, ",") {
			num, _ := strconv.Atoi(i)
			lists[len(lists)-1] = append(lists[len(lists)-1], num)
		}
	}

	return orderings, lists

}

// We sort the list with a custom ordering function based on the input
// If the given list does not change after sorting, we add the median of the list
func RightOrder() string {
	orderings, lists := parseInput()

	total := 0
	for _, l := range lists {
		if Tools.Equal(l, slices.SortedFunc(slices.Values(l), func(a int, b int) int {
			if slices.Contains(orderings[b], a) {
				return 1
			}

			if slices.Contains(orderings[a], b) {
				return -1
			}
			return 0
		})) {
			total += l[(len(l)-1)/2]
		}
	}

	return strconv.Itoa(total)
}

// We sort the list with a custom ordering function based on the input
// If this changes the list, we add the median
func WrongOrder() string {
	orderings, lists := parseInput()

	total := 0
	for _, l := range lists {
		sorted := slices.SortedFunc(slices.Values(l), func(a int, b int) int {
			if slices.Contains(orderings[b], a) {
				return 1
			}

			if slices.Contains(orderings[a], b) {
				return -1
			}
			return 0
		})

		if !Tools.Equal(l, sorted) {
			total += sorted[(len(l)-1)/2]
		}
	}

	return strconv.Itoa(total)
}
