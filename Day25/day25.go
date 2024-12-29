package day25

import (
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

// Encodes each key/lock as a list of 5 integers between 0 and 7
// A key and lock then fit if the sum of the two lists is <= 7 on each index
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

// Just press the button on the site
func MerryChristmas() string {
	return "Merry Christmas"
}
