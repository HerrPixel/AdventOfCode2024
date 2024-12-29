package day24

import (
	"slices"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

type gate struct {
	a string
	b string
	f func(bool, bool) bool
}

func parseInput() (map[string]bool, map[string]gate) {
	input := strings.Split(Tools.Read("./Day24/input.txt"), "\n\n")

	values := make(map[string]bool, 0)
	gates := make(map[string]gate, 0)

	for _, l := range strings.Split(input[0], "\n") {
		line := strings.Split(l, ": ")

		key := line[0]
		value := false

		if line[1] == "1" {
			value = true
		}

		values[key] = value
	}

	for _, l := range strings.Split(input[1], "\n") {
		line := strings.Split(l, " -> ")
		output := line[1]
		input := strings.Split(line[0], " ")
		a := input[0]
		b := input[2]
		op := XOR

		if input[1] == "OR" {
			op = OR
		} else if input[1] == "AND" {
			op = AND
		}

		gates[output] = gate{a, b, op}
	}

	return values, gates

}

// We simulate the circuit as instructed, starting from each zXX gate and following the gates backwards
// If a gate has a calculated output already, we use this precomputed value, otherwise we recursively calulate the output of that gate
func SimulateCircuit() string {
	values, gates := parseInput()

	var getValue func(string) bool

	getValue = func(s string) bool {
		res, ok := values[s]

		// if the output of a gate is already set, use it instead of calculating it again
		if ok {
			return res
		}

		// otherwise calculate all inputs and use the combinator of this gate
		gate := gates[s]
		a := getValue(gate.a)
		b := getValue(gate.b)

		res = gate.f(a, b)
		values[s] = res
		return res
	}

	total := 0

	for key := range gates {
		if key[0] != 'z' {
			continue
		}
		num, _ := strconv.Atoi(key[1:])
		value := getValue(key)

		if value {
			total += (1 << num)
		}
	}

	return strconv.Itoa(total)
}

// We work very heuristically here;
// We assume the circuitry is a Carry-Ripple-Adder, i.e.
// xN ───────╦───┬───┐
//
//	│   │XOR├─╦──────┬───┐
//
// yN ─────╦─────┴───┘ │      │XOR├────── zN
//
//	│ │     ┌──────────┴───┘
//
// carryN ─────────╣   │
//
//	│ │     │   └┬───┐
//	│ │     │    │AND├┐
//	│ │     └────┴───┘└┬───┐
//	│ │                │ OR├────── carryN+1
//	│ └───┬───┐    ┌───┴───┘
//	│     │AND├────┘
//	└─────┴───┘
//
// For our input, all faults where somewhere in the upper XOR-chain
// We therefore only test two things: if x XOR y is an input to the XOR-gate for z
// and also if XOR-gate with inputs (x XOR y) and carryN is actually the output z
// If there is a fault, we switch the ouput wire with the one it should actually be, i.e. the last XOR-gate before z or z itself.
// This is enough for our input, finally we sort the names of the faulty gates and output them.
func WrongWires() string {
	_, gates := parseInput()

	var x [45]string
	var y [45]string
	var z [45]string
	var gateAND [45]string
	var gateCarry [45]string
	var gateOR [45]string
	var gateXOR [45]string
	var gateZ [45]string

	for i := range 45 {
		s := strconv.Itoa(i)
		if i <= 9 {
			s = "0" + s
		}

		x[i] = "x" + s
		y[i] = "y" + s
		z[i] = "z" + s
	}

	i := 0
	gateXOR[0] = findGate(gates, "x00", "y00", XOR)
	gateAND[0] = findGate(gates, "x00", "y00", AND)
	gateZ[0] = findGate(gates, "x00", "y00", XOR)
	gateCarry[0] = ""
	gateOR[0] = findGate(gates, "x00", "y00", AND)
	i++

	faults := []string{}

	for i < 45 {

		gateXOR[i] = findGate(gates, x[i], y[i], XOR)

		// Tests if x XOR y is an input to z, otherwise swaps
		if gateXOR[i] != gates[z[i]].a && gateOR[i-1] == gates[z[i]].b {
			gates = swap(gates, gateXOR[i], gates[z[i]].a)
			faults = append(faults, gateXOR[i], gates[z[i]].a)

			gateXOR[i] = gates[z[i]].a
		} else if gateXOR[i] != gates[z[i]].b && gateOR[i-1] == gates[z[i]].a {
			gates = swap(gates, gateXOR[i], gates[z[i]].b)
			faults = append(faults, gateXOR[i], gates[z[i]].b)

			gateXOR[i] = gates[z[i]].b
		}

		gateZ[i] = findGate(gates, gateXOR[i], gateOR[i-1], XOR)

		//tests if (x XOR y) XOR carry is output z, otherwise swaps
		if gateZ[i] != z[i] {
			gates = swap(gates, gateZ[i], z[i])
			faults = append(faults, gateZ[i], z[i])

			gateZ[i] = z[i]
		}

		// apparently, these are the only errors
		// this is heuristically though

		// setting the rest of the gates
		gateAND[i] = findGate(gates, x[i], y[i], AND)
		gateCarry[i] = findGate(gates, gateXOR[i], gateOR[i-1], AND)
		gateOR[i] = findGate(gates, gateCarry[i], gateAND[i], OR)

		i++
	}

	slices.Sort(faults)

	return strings.Join(faults, ",")
}

// finds the output wire of a gate with input a and b and operation op
func findGate(gates map[string]gate, a string, b string, op func(bool, bool) bool) string {
	for key, value := range gates {
		if (value.a == a && value.b == b) || (value.a == b && value.b == a) {
			if sameOperation(value.f, op) {
				return key
			}
		}
	}
	return ""
}

// tests if f and g are the same function out of the set {AND,OR,XOR}
func sameOperation(f func(bool, bool) bool, g func(bool, bool) bool) bool {
	return f(true, true) == g(true, true) && f(true, false) == g(true, false)
}

func swap(gates map[string]gate, a string, b string) map[string]gate {
	gates[a], gates[b] = gates[b], gates[a]
	return gates
}

func AND(a bool, b bool) bool {
	return a && b
}

func OR(a bool, b bool) bool {
	return a || b
}

func XOR(a bool, b bool) bool {
	return (a || b) && !(a && b)
}
