package day3

import (
	"regexp"
	"strconv"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

// finding correct multiplication instructions via regex search and executing them
func Multiplications() string {
	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

	total := 0
	s := Tools.Read("./Day3/input.txt")

	for _, match := range r.FindAllStringSubmatch(s, -1) {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		total += x * y
	}
	return strconv.Itoa(total)
}

// finding correct instructions via regex search
// The new instructions just flip a boolean value that enables the multiplication to take place
func EnabledMultiplications() string {
	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`)

	total := 0
	s := Tools.Read("./Day3/input.txt")
	enabled := true

	for _, match := range r.FindAllStringSubmatch(s, -1) {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += x * y
		}

	}
	return strconv.Itoa(total)
}
