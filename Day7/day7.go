package day7

import (
	"math"
	"slices"
	"strconv"
	"strings"

	"git.jonasseiler.de/Jonas/AdventOfCode2024/Tools"
)

type equation struct {
	result int
	inputs []int
}

func parseInput() []equation {
	equations := make([]equation, 0)

	for _, line := range Tools.ReadByLines("./Day7/input.txt") {
		splitLine := strings.Split(line, ": ")
		result, _ := strconv.Atoi(splitLine[0])
		numbers := make([]int, 0)

		for _, s := range strings.Split(splitLine[1], " ") {
			n, _ := strconv.Atoi(s)
			numbers = append(numbers, n)
		}
		equations = append(equations, equation{result, numbers})
	}

	return equations
}

func Calibrations() string {
	equations := parseInput()

	total := 0

	for _, equation := range equations {
		numbers := equation.inputs

		intermediateResults := make([]int, 0)
		intermediateResults = append(intermediateResults, equation.result)

		numbers = Tools.Reverse(numbers)
		isCorrect := false

		for index, y := range numbers {
			newIntermediateResults := make([]int, 0)

			if index == len(numbers)-1 {
				isCorrect = slices.Contains(intermediateResults, y)
				break
			}

			for _, x := range intermediateResults {
				if isDivisbleBy(x, y) {
					newIntermediateResults = append(newIntermediateResults, x/y)
				}

				if isSubtractableBy(x, y) {
					newIntermediateResults = append(newIntermediateResults, x-y)
				}
			}
			intermediateResults = newIntermediateResults
		}

		if isCorrect {
			total += equation.result
		}
	}

	return strconv.Itoa(total)
}

func CalibrationsWithConcatenation() string {
	equations := parseInput()

	total := 0

	for _, equation := range equations {
		numbers := equation.inputs

		intermediateResults := make([]int, 0)
		intermediateResults = append(intermediateResults, equation.result)

		numbers = Tools.Reverse(numbers)
		isCorrect := false

		for index, y := range numbers {
			newIntermediateResults := make([]int, 0)

			if index == len(numbers)-1 {
				isCorrect = slices.Contains(intermediateResults, y)
				break
			}

			for _, x := range intermediateResults {
				if isDivisbleBy(x, y) {
					newIntermediateResults = append(newIntermediateResults, x/y)
				}

				if isSubtractableBy(x, y) {
					newIntermediateResults = append(newIntermediateResults, x-y)
				}

				if endsWith(x, y) {
					newIntermediateResults = append(newIntermediateResults, removeLastNDigits(x, len(strconv.Itoa(y))))
				}
			}
			intermediateResults = newIntermediateResults
		}

		if isCorrect {
			total += equation.result
		}
	}

	return strconv.Itoa(total)
}

func endsWith(x int, y int) bool {
	len := len(strconv.Itoa(y))
	return x%int(math.Pow10(len)) == y
}

func isDivisbleBy(x int, y int) bool {
	return x%y == 0
}

func isSubtractableBy(x int, y int) bool {
	return x-y >= 0
}

func removeLastNDigits(x int, n int) int {
	return x / int(math.Pow10(n))
}
