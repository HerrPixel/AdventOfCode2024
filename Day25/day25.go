package day25

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func KeyLockCombinations() string {
	keys := make([][5]int, 0)
	locks := make([][5]int, 0)

	for _, input := range strings.Split(Tools.Read("./Day25/input.txt"), "\n\n") {
		list := [5]int{0, 0, 0, 0, 0}
		lines := strings.Split(input, "\n")
		for _, l := range lines {
			for i, c := range strings.Split(l, "") {
				if c == "#" {
					list[i]++
				}
			}
		}
		if lines[0] == "#####" {
			locks = append(locks, list)
		} else {
			keys = append(keys, list)
		}
	}

	total := 0

	for _, l := range locks {
		for _, k := range keys {

			isValid := true

			for i := range 5 {
				if l[i]+k[i] > 7 {
					isValid = false
					break
				}
			}

			if isValid {
				total++
			}
		}
	}

	return strconv.Itoa(total)
}

func MerryChristmas() string {
	return "Merry Christmas"
}
