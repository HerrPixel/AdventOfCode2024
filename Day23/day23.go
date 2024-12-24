package day23

import (
	"slices"
	"strconv"
	"strings"

	"github.com/HerrPixel/AdventOfCode2024/Tools"
)

func parseInput() map[string][]string {
	nodes := make(map[string][]string)

	for _, l := range Tools.ReadByLines("./Day23/input.txt") {
		input := strings.Split(l, "-")

		_, exists := nodes[input[0]]
		if !exists {
			nodes[input[0]] = []string{}
		}

		nodes[input[0]] = append(nodes[input[0]], input[1])

		_, exists = nodes[input[1]]
		if !exists {
			nodes[input[1]] = []string{}
		}

		nodes[input[1]] = append(nodes[input[1]], input[0])
	}
	return nodes
}

func Part1() string {
	nodes := parseInput()

	found := make(map[string]bool, 0)

	for key, value := range nodes {
		if key[0] != 't' {
			continue
		}

		for _, n := range value {
			for _, m := range value {
				if slices.Contains(nodes[n], m) {
					tripple := []string{key, n, m}
					slices.Sort(tripple)
					hash := strings.Join(tripple[:], ",")

					_, exists := found[hash]
					if exists {
						continue
					}

					found[hash] = true
				}
			}
		}
	}

	return strconv.Itoa(len(found))
}

func Part2() string {
	nodes := parseInput()
	keys := make([]string, 0, len(nodes))
	for k, _ := range nodes {
		keys = append(keys, k)
	}

	maximalClique := BronKerbosch([]string{}, keys, []string{}, nodes)
	slices.Sort(maximalClique)

	return strings.Join(maximalClique, ",")
}

func BronKerbosch(R []string, P []string, X []string, nodes map[string][]string) []string {
	if len(P) == 0 && len(X) == 0 {
		return R
	}

	maximalClique := make([]string, 0)
	pivot := ""
	if len(P) == 0 {
		pivot = X[0]
	} else {
		pivot = P[0]
	}

	pivotNeighborhood := nodes[pivot]

	for _, v := range without(P, pivotNeighborhood) {
		vNeighborhood := nodes[v]

		res := BronKerbosch(
			union(R, v),
			intersection(P, vNeighborhood),
			intersection(X, vNeighborhood),
			nodes,
		)

		if len(res) > len(maximalClique) {
			maximalClique = res
		}

		P = without(P, []string{v})
		X = append(X, v)
	}

	return maximalClique
}

func union(l []string, s string) []string {
	return append(Tools.Clone(l), s)
}

func intersection(l []string, m []string) []string {
	res := make([]string, 0)

	for _, s := range l {
		if slices.Contains(m, s) {
			res = append(res, s)
		}
	}

	return res
}

func without(l []string, m []string) []string {
	res := make([]string, 0)

	for _, s := range l {
		if !slices.Contains(m, s) {
			res = append(res, s)
		}
	}

	return res
}
