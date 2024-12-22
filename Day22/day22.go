package day22

import (
	"strconv"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInputs() []int {
	numbers := make([]int, 0)

	for _, l := range Tools.ReadByLines("./Day22/input.txt") {
		n, _ := strconv.Atoi(l)

		numbers = append(numbers, n)
	}

	return numbers
}

func Part1() string {
	numbers := parseInputs()

	total := 0

	for _, i := range numbers {
		n := i
		for range 2000 {
			n = evolve(n)
		}
		total += n
	}

	return strconv.Itoa(total)
}

func Part2() string {
	numbers := parseInputs()

	sequences := make(map[int32]int, 0)

	for _, i := range numbers {
		hash := int32(0)
		n := i
		prev := n % 10
		buyerSequences := make(map[int32]bool, 0)

		for j := range 2000 {
			n = evolve(n)
			value := n % 10
			hash = slide(hash, value-prev)
			prev = value

			if j < 3 {
				continue
			}

			_, seenBefore := buyerSequences[hash]

			if seenBefore {
				continue
			}

			buyerSequences[hash] = true

			_, seenBefore = sequences[hash]

			if !seenBefore {
				sequences[hash] = 0
			}

			sequences[hash] += value
		}
	}

	bananas := 0

	for _, v := range sequences {
		bananas = max(bananas, v)
	}

	return strconv.Itoa(bananas)
}

func slide(n int32, d int) int32 {
	n = n << 5
	negative := 0
	if d < 0 {
		negative = 1 << 4
	}
	d = Tools.Abs(d) % (1 << 4)
	n = n ^ int32(d) ^ int32(negative)
	n = n % (1 << 20)

	return n
}

func evolve(n int) int {
	// Step 1
	n = prune(mix(n, n*64))

	// Step 2
	n = prune(mix(n, n/32))

	// Step 3
	n = prune(mix(n, n*2048))

	return n
}

func mix(n int, m int) int {
	return n ^ m
}

func prune(n int) int {
	return n % 16777216
}
