package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() ([]int, []int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, s := range Tools.ReadByLines("./Day1/input.txt") {
		locations := strings.Split(s, "   ")

		left, _ := strconv.Atoi(locations[0])
		right, _ := strconv.Atoi(locations[1])

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList
}

// Sorting both lists and summing up their difference
func TotalDistance() string {
	leftList, rightList := parseInput()

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0

	for i := range leftList {
		totalDistance += Tools.Abs(leftList[i] - rightList[i])
	}

	return strconv.Itoa(totalDistance)
}

// First counting multiplicities of one list and iterating through the other list, summing up the respective scores
// This results in an expected linear time instead of naive quadratic time
func SimilarityScore() string {
	leftList, rightList := parseInput()

	hashedRightList := make(map[int]int)

	for _, x := range rightList {
		hashedRightList[x] += 1
	}

	similarityScore := 0

	for _, x := range leftList {
		similarityScore += x * hashedRightList[x]
	}

	return strconv.Itoa(similarityScore)
}
