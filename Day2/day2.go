package day2

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() [][]int {
	reports := make([][]int, 0)

	for _, s := range Tools.ReadByLines("./Day2/input.txt") {
		levels := strings.Split(s, " ")

		reports = append(reports, make([]int, 0))
		for _, levelString := range levels {
			level, _ := strconv.Atoi(levelString)
			reports[len(reports)-1] = append(reports[len(reports)-1], level)
		}
	}

	return reports
}

// Iterating through every list, we check that the conditions are satisfied
func SafeReports() string {
	reports := parseInput()

	safeReports := 0

	for _, report := range reports {
		if hasMinAndMaxDifference(report, 1, 3) && (isIncreasing(report) || isDecreasing(report)) {
			safeReports++

		}
	}

	return strconv.Itoa(safeReports)
}

// We naively remove an element and check if that list is valid
// If this works for at least one element, we increment the count
// This is pure brute force but it is fast enough
func DampenedReports() string {
	reports := parseInput()

	safeReports := 0

	for _, report := range reports {
		if hasMinAndMaxDifference(report, 1, 3) && (isIncreasing(report) || isDecreasing(report)) {
			safeReports++

		} else {
			for i := range report {
				dampenedReport := Tools.Remove(report, i)

				if hasMinAndMaxDifference(dampenedReport, 1, 3) && (isIncreasing(dampenedReport) || isDecreasing(dampenedReport)) {
					safeReports++
					break
				}

			}
		}
	}

	return strconv.Itoa(safeReports)
}

func isIncreasing(l []int) bool {
	curr := l[0]

	for _, next := range l {
		if next < curr {
			return false
		}
		curr = next
	}
	return true
}

func isDecreasing(l []int) bool {
	curr := l[0]

	for _, next := range l {
		if next > curr {
			return false
		}
		curr = next
	}
	return true
}

func hasMinAndMaxDifference(l []int, min int, max int) bool {
	curr := l[0]

	for i, next := range l {
		if i == 0 {
			continue
		}
		if max < Tools.Abs(curr-next) || Tools.Abs(curr-next) < min {
			return false
		}
		curr = next
	}
	return true
}
