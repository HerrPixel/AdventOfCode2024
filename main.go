package main

import (
	"fmt"

	day1 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day1"
	day2 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day2"
	day3 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day3"
	day4 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day4"
)

type solution struct {
	name          string
	part1function func() string
	part2function func() string
}

func main() {

	solutions := []solution{
		{"Historian Hysteria", day1.TotalDistance, day1.SimilarityScore},
		{"Red-Nosed Reports", day2.SafeReports, day2.DampenedReports},
		{"Mull It Over", day3.Multiplications, day3.EnabledMultiplications},
		{"Ceres Search", day4.XMAS, day4.X_MAS},
	}

	for i, s := range solutions {
		fmt.Println("Day", i+1, ":", s.name)

		fmt.Println("    Part1:", s.part1function())

		fmt.Println("    Part2:", s.part2function())
		fmt.Println("")

	}
}
