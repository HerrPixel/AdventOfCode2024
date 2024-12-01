package main

import (
	"fmt"

	day1 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day1"
)

type solution struct {
	name          string
	part1function func() string
	part2function func() string
}

func main() {

	solutions := []solution{
		{"Historian Hysteria", day1.TotalDistance, day1.SimilarityScore},
	}

	for i, s := range solutions {
		fmt.Println("Day", i+1, ":", s.name)

		fmt.Println("    Part1:", s.part1function())
		//fmt.Println("")

		fmt.Println("    Part2:", s.part2function())
		fmt.Println("")

	}
}
