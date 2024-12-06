package main

import (
	"fmt"
	"time"

	day1 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day1"
	day2 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day2"
	day3 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day3"
	day4 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day4"
	day5 "git.jonasseiler.de/Jonas/AdventOfCode2024/Day5"
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
		{"Print Queue", day5.RightOrder, day5.WrongOrder},
	}

	start_global := time.Now()

	for i, s := range solutions {
		fmt.Println("Day", i+1, ":", s.name)

		start := time.Now()
		fmt.Println("    Part1:", s.part1function())
		t := time.Now()
		fmt.Println("    Took", t.Sub(start).Microseconds(), "μs")

		start = time.Now()
		fmt.Println("    Part2:", s.part2function())
		t = time.Now()
		fmt.Println("    Took", t.Sub(start).Microseconds(), "μs")

		fmt.Println("")
	}

	t := time.Now()
	fmt.Println("Final time:", t.Sub(start_global).Milliseconds(), "ms")
}
