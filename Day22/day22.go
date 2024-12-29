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

// We implement the evolve instructions and just apply them 2000 times to each secret number
func SecretNumbers() string {
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

// We use a global map of 4-sequences -> bananas which we update every time we see this sequence the first time on a monkey
// We also use a local map of 4-sequences -> bool to keep track if we have seen this sequence on this monkey before and therefore ignore it
// Finally we take the maximum value of the global sequence map
func MarketOptimization() string {
	numbers := parseInputs()

	// to make it easier to use a map here, we constructed a hash function that hashes a 4-digit sequence with values between -9 and +9 to a unique 2-bit number and use this to index the map
	// see also the slide function below
	sequences := make(map[int32]int, 0)

	for _, i := range numbers {
		hash := int32(0)
		n := i
		prev := n % 10
		buyerSequences := make(map[int32]bool, 0)

		for j := range 2000 {
			n = evolve(n)
			value := n % 10

			// our unique hash
			hash = slide(hash, value-prev)
			prev = value

			// we have not seen enough values yet to have a 4-sequence
			if j < 3 {
				continue
			}

			_, seenBefore := buyerSequences[hash]

			if seenBefore {
				continue
			}

			// now we have seen this sequence on this monkey before and won't consider it again
			buyerSequences[hash] = true

			_, seenBefore = sequences[hash]

			if !seenBefore {
				sequences[hash] = 0
			}

			// add the value of this sequence to our global map
			sequences[hash] += value
		}
	}

	bananas := 0

	for _, v := range sequences {
		bananas = max(bananas, v)
	}

	return strconv.Itoa(bananas)
}

// maps a 4-sequence of digits between -9 and +9 to a unique 20-bit number
// Each 5th bit denotes the sign and the next 4 bits determine the number between 0 and 9
// When we add a new digit, we slide the last digit out by shifting the number 5 spaces to the left and fill in the next number
func slide(n int32, d int) int32 {
	n = n << 5

	// sign bit
	negative := 0
	if d < 0 {
		negative = 1 << 4
	}

	// we only care about the last digit of a secret number, so we cut off all but the last 4 bits
	d = Tools.Abs(d) % (1 << 4)

	// fill in the new number
	n = n ^ int32(d) ^ int32(negative)

	// cut off the old number
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
