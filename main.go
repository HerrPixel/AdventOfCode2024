package main

import (
	"fmt"
	"time"

	day1 "github.com/HerrPixel/AdventOfCode2024/Day1"
	day10 "github.com/HerrPixel/AdventOfCode2024/Day10"
	day11 "github.com/HerrPixel/AdventOfCode2024/Day11"
	day12 "github.com/HerrPixel/AdventOfCode2024/Day12"
	day13 "github.com/HerrPixel/AdventOfCode2024/Day13"
	day14 "github.com/HerrPixel/AdventOfCode2024/Day14"
	day15 "github.com/HerrPixel/AdventOfCode2024/Day15"
	day16 "github.com/HerrPixel/AdventOfCode2024/Day16"
	day17 "github.com/HerrPixel/AdventOfCode2024/Day17"
	day18 "github.com/HerrPixel/AdventOfCode2024/Day18"
	day19 "github.com/HerrPixel/AdventOfCode2024/Day19"
	day2 "github.com/HerrPixel/AdventOfCode2024/Day2"
	day20 "github.com/HerrPixel/AdventOfCode2024/Day20"
	day21 "github.com/HerrPixel/AdventOfCode2024/Day21"
	day22 "github.com/HerrPixel/AdventOfCode2024/Day22"
	day23 "github.com/HerrPixel/AdventOfCode2024/Day23"
	day24 "github.com/HerrPixel/AdventOfCode2024/Day24"
	day25 "github.com/HerrPixel/AdventOfCode2024/Day25"
	day3 "github.com/HerrPixel/AdventOfCode2024/Day3"
	day4 "github.com/HerrPixel/AdventOfCode2024/Day4"
	day5 "github.com/HerrPixel/AdventOfCode2024/Day5"
	day6 "github.com/HerrPixel/AdventOfCode2024/Day6"
	day7 "github.com/HerrPixel/AdventOfCode2024/Day7"
	day8 "github.com/HerrPixel/AdventOfCode2024/Day8"
	day9 "github.com/HerrPixel/AdventOfCode2024/Day9"
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
		{"Guard Gallivant", day6.GuardLeave, day6.Obstruction},
		{"Bridge Repair", day7.Calibrations, day7.CalibrationsWithConcatenation},
		{"Resonant Collinearity", day8.AntinodePairs, day8.Antinodes},
		{"Disk Fragmenter", day9.Part1, day9.Part2},
		{"Hoof It", day10.HikingTrails, day10.DistinctHikingTrails},
		{"Plutonian Pebbles", day11.Blinking, day11.MoreBlinking},
		{"Garden Groups", day12.PerimeterFences, day12.SideFences},
		{"Claw Contraption", day13.PrizeTokens, day13.FarawayPrizeTokens},
		{"Restroom Redoubt", day14.Bathroom, day14.EasterEgg},
		{"Warehouse Woes", day15.GPSCoordinates, day15.DoubledChests},
		{"Reindeer Maze", day16.ReindeerPath, day16.BenchSpots},
		{"Chronospatial Computer", day17.SimulateProgramm, day17.SelfReplicationValue},
		{"RAM Run", day18.ByteRainPath, day18.LatestByteRainPath},
		{"Linen Layout", day19.PossibleDesigns, day19.DesignCombinations},
		{"Race Condition", day20.ShortShortcuts, day20.LongShortcuts},
		{"Keypad Conundrum", day21.ShortComplexities, day21.LongComplexities},
		{"Monkey Market", day22.SecretNumbers, day22.MarketOptimization},
		{"LAN Party", day23.Triangles, day23.LargestClique},
		{"Crossed Wires", day24.SimulateCircuit, day24.WrongWires},
		{"Code Chronicle", day25.KeyLockCombinations, day25.MerryChristmas},
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
