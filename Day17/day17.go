package day17

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() (int, int, int, []int) {
	instructions := make([]int, 0)

	input := strings.Split(Tools.Read("./Day17/input.txt"), "\n\n")
	r := regexp.MustCompile(`Register .: ([0-9]+)`)

	registers := strings.Split(input[0], "\n")

	a, _ := strconv.Atoi(r.FindStringSubmatch(registers[0])[1])
	b, _ := strconv.Atoi(r.FindStringSubmatch(registers[1])[1])
	c, _ := strconv.Atoi(r.FindStringSubmatch(registers[2])[1])

	r = regexp.MustCompile(`[, ]([0-9]+)`)

	for _, s := range r.FindAllStringSubmatch(input[1], -1) {
		n, _ := strconv.Atoi(s[1])
		instructions = append(instructions, n)
	}

	return a, b, c, instructions
}

// We just simulate the programm as instructed
func SimulateProgramm() string {
	a, b, c, instructions := parseInput()

	numbers := run(instructions, a, b, c)
	chars := make([]string, 0)

	for _, i := range numbers {
		chars = append(chars, strconv.Itoa(i))
	}

	return strings.Join(chars, ",")
}

// We heuristically solve this in reverse;
// In my input, every 3 bits of A lead to one output
// Registers B and C are reset between each output.
// Therefore we simulate each 3-bit-combination of A to test if the output is correct
// We then use this suffix and test the next 3 bits and so on until A outputs everything we need
func SelfReplicationValue() string {
	_, _, _, instructions := parseInput()

	candidates := make([]int, 0)
	candidates = append(candidates, 0)

	// For each digit in the instruction, we need to test all 3-bit combinations
	for matchingDigits := range len(instructions) {
		matchingDigits += 1
		newCandidates := make([]int, 0)

		// There are 8 = 2^3 combinations for 3 bits
		for i := range 8 {

			// We need to use the previous results, since the output is also influenced by some bits of A afterwards
			// This is why we try this in reverse, since the last 3 bits are not influenced by anything else
			for _, c := range candidates {
				candidate := 8*c + i

				outputs := run(instructions, candidate, 0, 0)

				if len(outputs) != matchingDigits {
					continue
				}

				// Test if the output is correct so far
				// But in reverse, see above
				valid := true
				for k := range matchingDigits {
					if instructions[len(instructions)-1-k] != outputs[len(outputs)-1-k] {
						valid = false
					}
				}

				if valid {
					newCandidates = append(newCandidates, candidate)
				}
			}
		}

		candidates = newCandidates
	}

	best := math.MaxInt64

	for _, n := range candidates {
		best = min(best, n)
	}

	return strconv.Itoa(best)
}

// implementation as described in the problem
// Some operations are simplified to bit magic
func run(instructions []int, a int, b int, c int) []int {
	output := make([]int, 0)

	instructionPointer := 0

	combo := func(x int) int {
		if x <= 3 {
			return x
		}

		if x == 4 {
			return a
		}

		if x == 5 {
			return b
		}

		if x == 6 {
			return c
		}

		return -1
	}

	for instructionPointer < len(instructions) {
		opCode := instructions[instructionPointer]
		operand := instructions[instructionPointer+1]

		if opCode == 0 {
			a = a >> combo(operand)
		}

		if opCode == 1 {
			b = b ^ operand
		}

		if opCode == 2 {
			b = combo(operand) & 7
		}

		if opCode == 3 {
			if a != 0 {
				instructionPointer = operand
				continue
			}
		}

		if opCode == 4 {
			b = b ^ c
		}

		if opCode == 5 {
			output = append(output, combo(operand)&7)
		}

		if opCode == 6 {
			b = a >> combo(operand)
		}

		if opCode == 7 {
			c = a >> combo(operand)
		}

		instructionPointer += 2
	}

	return output
}
