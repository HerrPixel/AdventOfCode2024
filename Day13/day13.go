package day13

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type machine struct {
	Ax     int
	Ay     int
	Bx     int
	By     int
	prizeX int
	prizeY int
}

func parseInput() []machine {
	machines := make([]machine, 0)

	buttonRegex := regexp.MustCompile(`Button [AB]: X\+([0-9]+), Y\+([0-9]+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)

	for _, input := range strings.Split(Tools.Read("./Day13/input.txt"), "\n\n") {
		lines := strings.Split(input, "\n")
		A := buttonRegex.FindStringSubmatch(lines[0])
		Ax, _ := strconv.Atoi(A[1])
		Ay, _ := strconv.Atoi(A[2])

		B := buttonRegex.FindStringSubmatch(lines[1])
		Bx, _ := strconv.Atoi(B[1])
		By, _ := strconv.Atoi(B[2])

		prize := prizeRegex.FindStringSubmatch(lines[2])
		prizeX, _ := strconv.Atoi(prize[1])
		prizeY, _ := strconv.Atoi(prize[2])

		machines = append(machines, machine{Ax, Ay, Bx, By, prizeX, prizeY})
	}

	return machines
}

// We calculate the solution to a 2-dimensional system of equations.
// If the solution is an integer, that is the number of presses.
// Essentially we want to solve c * Ax + d * Bx = prizeX and c * Ay + d * By = prizeY
// We can then write this as a matrix equation:
// |Ax Bx| |c| = |prizeX|
// |Ay By| |d| = |prizeY|
// Assuming that these are always invertible, we can easily write the inverse of this 2x2 matrix and solve for c and d
// |c| = _________1_________ |By -Bx| |prizeX|
// |d| = (Ax * By - Bx * Ay) |-Ay Ax| |prizeY|
// The solving can then be done purely with integers as only the last division can create non-integer solutions.
// If this happens, i.e the remainder of this division is not zero, the solution is not an integer and we can just throw it away
func PrizeTokens() string {
	machines := parseInput()

	total := 0

	for _, m := range machines {
		a, aRem := divRem(m.prizeX*m.By-m.prizeY*m.Bx, m.Ax*m.By-m.Ay*m.Bx)

		b, bRem := divRem(m.prizeY*m.Ax-m.prizeX*m.Ay, m.Ax*m.By-m.Ay*m.Bx)

		if aRem != 0 || bRem != 0 {
			continue
		}

		if 0 <= a && a <= 100 && 0 <= b && b <= 100 {
			total += 3*a + b
		}

	}
	return strconv.Itoa(total)
}

// Same as Part 1 but we increase each coordinate
func FarawayPrizeTokens() string {
	machines := parseInput()

	total := 0

	for _, m := range machines {
		m.prizeX += 10000000000000
		m.prizeY += 10000000000000

		a, aRem := divRem(m.prizeX*m.By-m.prizeY*m.Bx, m.Ax*m.By-m.Ay*m.Bx)

		b, bRem := divRem(m.prizeY*m.Ax-m.prizeX*m.Ay, m.Ax*m.By-m.Ay*m.Bx)

		if aRem != 0 || bRem != 0 {
			continue
		}

		if 0 <= a && 0 <= b {
			total += 3*a + b
		}

	}
	return strconv.Itoa(total)
}

func divRem(x int, y int) (int, int) {
	return x / y, x % y
}
