package day13

import (
	"regexp"
	"strconv"
	"strings"

	"git.jonasseiler.de/Jonas/AdventOfCode2024/Tools"
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
