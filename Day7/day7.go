package day7

import (
	"math"
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

	for _, s := range Tools.ReadByLines("./Day7/input.txt") {
		split := strings.Split(s, ": ")
		result, _ := strconv.Atoi(split[0])
		numbers := make([]int, 0)

		for _, n := range strings.Split(split[1], " ") {
			num, _ := strconv.Atoi(n)
			numbers = append(numbers, num)
		}
		equations = append(equations, equation{result, numbers})
	}

	return equations
}

func Part1() string {
	equations := parseInput()

	total := 0

	for _, e := range equations {
		nums := e.inputs

		intermediateResults := make([]int, 0)
		intermediateResults = append(intermediateResults, e.result)

		nums = Tools.Reverse(nums)

		isCorrect := false
		for index, i := range nums {
			newIntermediateResults := make([]int, 0)

			for _, j := range intermediateResults {
				if j%i == 0 {
					newIntermediateResults = append(newIntermediateResults, j/i)

					if index == len(nums)-1 && j/i == 1 {
						isCorrect = true
					}
				}

				if j-i >= 0 {
					newIntermediateResults = append(newIntermediateResults, j-i)
					if index == len(nums)-1 && j-i == 0 {
						isCorrect = true
					}
				}
			}
			intermediateResults = newIntermediateResults
		}

		if isCorrect {
			total += e.result
		}
	}

	return strconv.Itoa(total)
}

func Part2() string {
	equations := parseInput()

	total := 0

	for _, e := range equations {
		nums := e.inputs

		intermediateResults := make([]int, 0)
		intermediateResults = append(intermediateResults, e.result)

		nums = Tools.Reverse(nums)

		isCorrect := false
		for index, i := range nums {
			newIntermediateResults := make([]int, 0)

			for _, j := range intermediateResults {
				if j%i == 0 {
					newIntermediateResults = append(newIntermediateResults, j/i)

					if index == len(nums)-1 && j/i == 1 {
						isCorrect = true
					}
				}

				if j-i >= 0 {
					newIntermediateResults = append(newIntermediateResults, j-i)
					if index == len(nums)-1 && j-i == 0 {
						isCorrect = true
					}
				}

				if endsWith(j, i) {
					newIntermediateResults = append(newIntermediateResults, j/int(math.Pow10(len(strconv.Itoa(i)))))
					if index == len(nums)-1 && len(strconv.Itoa(j)) == len(strconv.Itoa(i)) {
						isCorrect = true
					}
				}
			}
			intermediateResults = newIntermediateResults
		}

		if isCorrect {
			total += e.result
		}
	}

	return strconv.Itoa(total)
}

func endsWith(x int, y int) bool {
	len := len(strconv.Itoa(y))
	return x%int(math.Pow10(len)) == y
}
